package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(
		sqlite.Open("Data.db"),
		&gorm.Config{},
	)
	if err != nil {
		panic(err.Error())
	} else {
		database.AutoMigrate(&User{})
		DB = database
	}
}