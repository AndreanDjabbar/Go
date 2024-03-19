package models

import (
	"SimpleWebAPI/models/identityModel"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(
		sqlite.Open("Identities.db"),
		&gorm.Config{},
	)

	if err != nil {
		panic("Failed Connect to Database")
	} else {
		database.AutoMigrate(&identityModel.Identity{})
		DB = database
	}
}