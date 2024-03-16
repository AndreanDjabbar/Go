package controllers

import (
	"SimpleCrud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MainRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/main/home/")
}

func HomePage(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	for index := range products {
		products[index].Index = index + 1
	}

	context := gin.H{
		"title":   "Home",
		"allData": products,
	}
	c.HTML(
		http.StatusOK,
		"home.html",
		context,
	)
}

func ViewCreate(c *gin.Context) {
	context := gin.H{
		"title": "Create",
	}
	c.HTML(
		http.StatusOK,
		"create.html",
		context,
	)
}

func Create(c *gin.Context) {
	name := c.PostForm("name")
	priceString := c.PostForm("price")
	stockString := c.PostForm("stock")
	description := c.PostForm("description")

	price, err := strconv.Atoi(priceString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Price"})
		return
	}
	stock, err := strconv.Atoi(stockString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Price"})
		return
	}

	newProduct := models.Product{
		Name:        name,
		Price:       price,
		Stock:       stock,
		Description: description,
	}

	models.DB.Create(&newProduct)
	c.Redirect(
		http.StatusMovedPermanently,
		"/main/home/",
	)
}

func Details(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			nil,
		)
	}

	var products models.Product
	if err := models.DB.First(&products, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data Not Found",
		})
		return
	}

	context := gin.H{
		"title":   "Detail",
		"product": products,
	}
	c.HTML(
		http.StatusOK,
		"details.html",
		context,
	)
}

func ViewUpdate(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	var products models.Product

	if err := models.DB.First(&products, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data Not Found",
		})
		return
	}

	context := gin.H{
		"title":   "Update Data",
		"product": products,
	}
	c.HTML(http.StatusOK, "update.html", context)
}

func Update(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var products models.Product

	if err := models.DB.First(&products, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Id not Found",
		})
		return
	}

	name := c.PostForm("name")
	stockString := c.PostForm("stock")
	priceString := c.PostForm("price")
	description := c.PostForm("description")

	stock, err := strconv.Atoi(stockString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Stock",
		})
		return
	}

	price, err := strconv.Atoi(priceString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Price",
		})
		return
	}

	products.Name = name
	products.Price = price
	products.Stock = stock
	products.Description = description

	models.DB.Save(&products)
	c.Redirect(http.StatusMovedPermanently, "/main/home/")
}

func DeleteValidation(c *gin.Context) {
	idString := c.Param("id");
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error":"Invalid ID"})
		return
	}

	var products models.Product;
	if err := models.DB.First(&products, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data Not Found",
		})
		return
	}

	context := gin.H{
		"title": "Delete",
		"product":products,
	}
	c.HTML(
		http.StatusOK,
		"delete.html",
		context,
	)
}

func Delete(c *gin.Context) {
	idString := c.Param("id")
    id, err := strconv.Atoi(idString)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var products models.Product
    if err := models.DB.First(&products, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    if err := models.DB.Delete(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data"})
        return
    }

    c.Redirect(http.StatusMovedPermanently, "/main/home/")
}