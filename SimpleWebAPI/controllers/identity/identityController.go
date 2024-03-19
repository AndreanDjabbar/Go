package identity

import (
	"SimpleWebAPI/models"
	"SimpleWebAPI/models/identityModel"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RootHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H {
			"message":"Welcome to Simple Web API",
		},
	)
}

func IdentitiesHandler(c *gin.Context) {
	var identities []identityModel.Identity
	models.DB.Find(&identities)
	c.JSON(
		http.StatusOK,
		gin.H {"Identities":identities},
	)
}

func DetailIdentityHandler(c *gin.Context) {
	id := c.Param("id")

	var identities identityModel.Identity
	err := models.DB.First(&identities, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H {"Error":"Data Not Found"},
			)
			return
		default:
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H {"Error":err.Error()},
			)
			return
		}
	}
	c.JSON(
		http.StatusOK,
		gin.H {"Identity":identities},
	)
}

func CreateHandler(c *gin.Context) {
	var identities identityModel.Identity
	
	err := c.ShouldBindJSON(&identities)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"Error": err.Error()},
		)
		return
	}

	if err := models.DB.Create(&identities).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"Error": "Failed to create identity", "error": err.Error()},
		)
		return
	}

	models.DB.Create(&identities)
	c.JSON(
		http.StatusOK,
		gin.H{"message": "Create Successful"},
	)
}

func UpdateHandler(c *gin.Context) {
	id := c.Param("id")
    var identity identityModel.Identity

    if err := models.DB.First(&identity, id).Error; err != nil {
        c.AbortWithStatusJSON(
            http.StatusNotFound,
            gin.H{"error": "Record not found"},
        )
        return
    }

    if err := c.ShouldBindJSON(&identity); err != nil {
        c.AbortWithStatusJSON(
            http.StatusBadRequest,
            gin.H{"error": err.Error()},
        )
        return
    }

    if err := models.DB.Save(&identity).Error; err != nil {
        c.AbortWithStatusJSON(
            http.StatusInternalServerError,
            gin.H{"error": "Failed to update record"},
        )
        return
    }

    c.JSON(
        http.StatusOK,
        gin.H{"Identity": identity},
    )
}

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	var identity identityModel.Identity

	if err := models.DB.First(&identity, id).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H {"Error":"Data Not Found"},
		)
		return
	}

	if err := models.DB.Delete(&identity).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H {"Error":"Failed to Delete Data"},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H {"Message":"Delete Data Successfull"},
	)
}