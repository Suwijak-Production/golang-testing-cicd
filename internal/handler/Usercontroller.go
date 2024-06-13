package handler

import (
	"myapp/internal/models"
	"myapp/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	var user models.User
	if err := repository.DB.Preload("Profiles").First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GetAllUser(c *gin.Context) {
	var users []models.User
	if err := repository.DB.Preload("UserBanks").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"UserBank": users,
	})
}
