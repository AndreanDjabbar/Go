package main

import (
	"SimpleCrud/controllers"
	"SimpleCrud/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default();
	router.LoadHTMLGlob("views/*.html");
	models.ConnectToDatabase();
	
	router.GET("/", controllers.MainRedirect);
	router.GET("/main/home/", controllers.HomePage);
	router.GET("/main/create/", controllers.ViewCreate);
	router.POST("/main/create/", controllers.Create);
	router.GET("/main/home/:id", controllers.Details);
	router.GET("/main/update/:id", controllers.ViewUpdate);
	router.POST("/main/update/:id", controllers.Update);
	router.GET("/main/delete/:id", controllers.DeleteValidation);
	router.POST("/main/delete/:id", controllers.Delete);
	
	err := router.Run("localhost:8000");
	if err != nil {
		log.Fatal(err)
	}
}
