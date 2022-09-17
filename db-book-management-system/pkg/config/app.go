package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	//Data Source Name
	dsn := "host=localhost user=postgres password=postgres dbname=goBookManagement port=5432 sslmode=disable TimeZone=Europe/Zurich"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// todo figure out how to connect to the wsl postgres
	// connection, err := gorm.Open("postgres", "postgres:postgres@/simplerest?charset=utf8")
	if err != nil {
		panic(err)
	}
	db = connection
}

func GetDB() *gorm.DB {
	return db
}