package controllers

import (
	"ecommerceeee/config"
	"ecommerceeee/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProd retrieves all products from the database
func GetProd(c *gin.Context) {
	var items []models.Product
	if err := config.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get products error"})
		return
	}

	c.JSON(http.StatusOK, items)
}
