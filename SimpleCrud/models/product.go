package models

type Product struct {
	Id    int    `gorm:"primary_key"`
	Index int `gorm:"type:int"`
	Name  string `gorm:"type:varchar(100)"`
	Price   int    `gorm:"type:int"`
	Stock   int    `gorm:"type:int"`
	Description string `gorm:"type:text"`
}