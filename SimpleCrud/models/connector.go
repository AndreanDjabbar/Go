package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func ConnectToDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{});
	if err != nil {
		panic("Failed to connect to database")
	} else {
		database.AutoMigrate(&Product{});
		DB = database;
	}
}