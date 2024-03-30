package main

import (
	"LoginRegister/controllers"
	"LoginRegister/models"

	"github.com/gin-gonic/gin"
)

func init() {
	models.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	router.GET("register/", controllers.ViewRegisterHandler)
	router.POST("register/", controllers.RegisterHandler)
	router.GET("login/", controllers.ViewLoginHandler)
	router.POST("login/", controllers.LoginHandler)
	router.GET("home/:name", controllers.HomeHandler)

	err := router.Run("localhost:4000")
	if err != nil {
		panic(err.Error())
	}
}