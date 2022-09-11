package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	// todo figure out how to connect to the wsl postgres
	connection, err := gorm.Open("postgres", "root:password/simplerest?")
	if err != nil {
		panic(err)
	}
	db = connection
}

func GetDB() *gorm.DB {
	return db
}