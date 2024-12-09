package controllers

import (
	"ecommerceeee/config"
	"ecommerceeee/models"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the product already exists
	var ExistingProduct models.Product
	err := config.DB.Where("ID=?", product.ID).First(&ExistingProduct).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Product doesn't exist, we can create it
		} else {
			// Log the actual error from the database query
			log.Printf("Error occurred during DB query: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error db"})
			return
		}
	} else {
		// If the product exists, return conflict
		c.JSON(http.StatusConflict, gin.H{"error": "Product already exists"})
		return
	}

	// Set the timestamps
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	// Create the product in the database
	if err := config.DB.Create(&product).Error; err != nil {
		log.Printf("Error creating product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating product"})
		return
	}
	// Return the created product
	c.JSON(http.StatusCreated, product)
}
