package main

import (
	"final_project/belee/app/routes"
	_buyerUsecase "final_project/belee/business/buyers"
	_buyerController "final_project/belee/controllers/buyers"
	_buyerRepository "final_project/belee/drivers/databases/buyers"
	_buyersdb "final_project/belee/drivers/databases/buyers"
	_mysqlDriver "final_project/belee/drivers/databases/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//fetch config.json
func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_buyersdb.Buyers{})
}

func main() {
	// init koneksi database
	configDb := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDb.InitialDB()

	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	buyerRepository := _buyerRepository.NewMysqlBuyerRepository(Conn)
	buyerUsecase := _buyerUsecase.NewBuyerUsecase(buyerRepository, timeoutContext)
	buyerController := _buyerController.NewBuyerController(buyerUsecase)

	routesInit := routes.ControllerList{
		BuyerController: *buyerController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
