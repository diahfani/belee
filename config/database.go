package config

import (
	// "github.com/jinzhu/gorm"
	"final_project/belee/models/buyers"
	"final_project/belee/models/owners"
	"final_project/belee/models/warung"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// **Contoh mas Rizky**

func DbConfig() {
	dsn := "root:diahaufa123@tcp(127.0.0.1:3308)/belee_test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed to connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&buyers.Buyers{}, &owners.Owners{}, &warung.Warungs{})
	// DB.AutoMigrate()
}

//**CONTOH DARI MEDIUM
// type DBconfig struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	DBName   string
// 	Password string
// }

// func BuildDBConfig() *DBconfig {
// 	dbConfig := DBconfig{
// 		Host:     "localhost",
// 		Port:     3308,
// 		User:     "root",
// 		Password: "diahaufa123",
// 		DBName:   "belee_test",
// 	}
// 	return &dbConfig
// }

// func DbURL(dbConfig *DBconfig) string {
// 	return fmt.Sprintf(
// 		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
// 		dbConfig.User,
// 		dbConfig.Password,
// 		dbConfig.Host,
// 		dbConfig.Port,
// 		dbConfig.DBName,
// 	)
// }
