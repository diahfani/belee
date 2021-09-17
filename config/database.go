package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBconfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBconfig {
	dbConfig := DBconfig{
		Host:     "localhost",
		Port:     3308,
		User:     "root",
		Password: "diahaufa123",
		DBName:   "belee_test",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBconfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
