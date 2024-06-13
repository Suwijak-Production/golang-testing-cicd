package handler

import (
	"myapp/internal/models"
	"myapp/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAdmin(c *gin.Context) {
	var Admins []models.Admin
	if err := repository.DB.Find(&Admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Admins": Admins,
	})
}

func CreateAdmin(c *gin.Context) {
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	if err := repository.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating admin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin created successfully",
		"admin":   admin,
	})
}

func UpdateAdmin(c *gin.Context) {
	var admin models.Admin
	id := c.Param("id")

	if err := repository.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Admin not found",
			"error":   err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	if err := repository.DB.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating admin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin updated successfully",
		"admin":   admin,
	})
}
