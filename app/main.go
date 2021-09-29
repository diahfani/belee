package main

import (
	_middleware "belee/app/middleware"
	// _middleware "final_project/belee/app/middleware"
	// _paymentRepository "belee/drivers/databases/paymentMethod"
	_paymentUsecase "belee/business/paymentMethod"
	_paymentController "belee/controllers/paymentMethod"
	_paymentRepository "belee/drivers/databases/paymentMethod"

	// "final_project/belee/app/middleware"
	"final_project/belee/app/routes"
	_buyerUsecase "final_project/belee/business/buyers"
	_ownerUsecase "final_project/belee/business/owners"
	_buyerController "final_project/belee/controllers/buyers"
	_ownerController "final_project/belee/controllers/owners"
	_buyerRepository "final_project/belee/drivers/databases/buyers"
	_buyersdb "final_project/belee/drivers/databases/buyers"
	_mysqlDriver "final_project/belee/drivers/databases/mysql"
	_ownerRepository "final_project/belee/drivers/databases/owners"
	_ownersdb "final_project/belee/drivers/databases/owners"
	_paymentdb "final_project/belee/drivers/databases/paymentMethod"

	"log"
	"time"

	// "final_project/belee/app/middleware"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	// "github.com/labstack/echo/v4/middleware"
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

// type Buyer struct {
// 	Role string
// }

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_buyersdb.Buyers{}, &_ownersdb.Owners{}, &_paymentdb.PaymentMethod{})

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

	configjwt := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	buyerRepository := _buyerRepository.NewMysqlBuyerRepository(Conn)
	buyerUsecase := _buyerUsecase.NewBuyerUsecase(buyerRepository, &configjwt, timeoutContext)
	buyerController := _buyerController.NewBuyerController(buyerUsecase)

	ownerRepository := _ownerRepository.NewMysqlOwnerRepository(Conn)
	ownerUsecase := _ownerUsecase.NewOwnerUsecase(ownerRepository, &configjwt, timeoutContext)
	ownerController := _ownerController.NewOwnerController(ownerUsecase)

	paymentRepository := _paymentRepository.NewMysqlPaymentRepo(Conn)
	paymentUsecase := _paymentUsecase.NewPaymentUsecase(paymentRepository, timeoutContext)
	paymentController := _paymentController.NewPaymentController(paymentUsecase)

	routesInit := routes.ControllerList{
		// Jwtconfig:     ,
		BuyerController:   *buyerController,
		OwnerController:   *ownerController,
		PaymentController: *paymentController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
