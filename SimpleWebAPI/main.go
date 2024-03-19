package main

import (
	"SimpleWebAPI/controllers/identity"
	"SimpleWebAPI/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectToDatabase()

	router.GET("/", identity.RootHandler)
	router.GET("identities/", identity.IdentitiesHandler)
	router.GET("identities/:id", identity.DetailIdentityHandler)
	router.POST("identities/create/", identity.CreateHandler)
	router.PUT("identities/:id", identity.UpdateHandler)
	router.DELETE("identities/:id", identity.DeleteHandler)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}