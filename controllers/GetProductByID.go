package controllers

import (
	"ecommerceeee/config"
	"ecommerceeee/models" // Update to the correct import path for your models
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductById(c *gin.Context) {
	// Retrieve the product ID from the URL parameters
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Create a variable to hold the product
	var product models.Product // Replace with your actual Product model

	// Query the database for the product
	result := config.DB.First(&product, productID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Return the product as a JSON response
	c.JSON(http.StatusOK, product)
}
