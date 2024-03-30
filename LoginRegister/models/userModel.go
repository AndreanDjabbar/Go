package models

type User struct {
	ID uint `gorm:"primayKey"`
	Email string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password string 
}
