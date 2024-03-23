package controllers

import (
	"LoginRegister/models"
	_"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ViewRegisterHandler(c *gin.Context) {
	context := gin.H {
		"title":"Register",
	}
	c.HTML(
		http.StatusOK,
		"register.html",
		context,
	)
}

func RegisterHandler(c *gin.Context) {
	var user models.User

	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	isAgree := c.PostForm("isAgree")
	if isAgree != "on" {
		c.JSON(
			http.StatusBadRequest,
			gin.H {"Error":"You must agree terms and condition"},
		)
		return
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	
	user.Username = username
	user.Email = email
	user.Password = string(hashedPassword)
	
	err = models.DB.Create(&user).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H {"Error":"Failed to create data"},
		)
	}
	c.Redirect(
		http.StatusMovedPermanently,
		"/login/",
	)
}

func ViewLoginHandler(c *gin.Context) {
	context := gin.H {
		"title":"Login",
	}
	c.HTML(
		http.StatusOK,
		"login.html",
		context,
	)
}

func LoginHandler(c *gin.Context) {
	var user models.User

	username := c.PostForm("username")
	password := c.PostForm("password")

	err := models.DB.Where("Username = ?", username).First(&user).Error
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H {"Error":"Invalid Username or Password"},
		)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H {"Error":"Invalid username os password"},
		)
		return
	}
	
	target := "/home/" + username
	c.Redirect(
		http.StatusMovedPermanently,
		target,
	)
}

func HomeHandler(c *gin.Context) {
	name := c.Param("name")

	context := gin.H {
		"title":"Home",
		"name":name,
	}
	c.HTML(
		http.StatusOK,
		"home.html",
		context,
	)
}