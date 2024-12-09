package controllers

import (
	"ecommerceeee/config"
	"ecommerceeee/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var req struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	// Bind incoming JSON request to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract email from context
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found in token"})
		return
	}

	// Ensure email is of the correct type
	userEmail, ok := email.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email format"})
		return
	}

	// Retrieve user from the database using email
	var user models.User
	if err := config.DB.Where("email = ?", userEmail).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	// Use the `user.ID` for cart operations
	var products models.Product
	if err := config.DB.Where("id = ?", req.ProductID).First(&products).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	var cartItem models.Cart
	// Check if the user already has this product in their cart
	if err := config.DB.Where("user_id = ? AND product_id = ?", user.ID, req.ProductID).First(&cartItem).Error; err == nil {
		// Product already in cart, update the quantity
		cartItem.Quantity += req.Quantity
		if err := config.DB.Save(&cartItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update cart"})
			return
		}
	} else {
		// New product, add to the cart
		newCartItem := models.Cart{
			UserID:    int(user.ID), // Directly use user ID
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
		}
		if err := config.DB.Create(&newCartItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add to cart"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "product added to cart"})
}

func GetCart(c *gin.Context) {
	// Extract email from context
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found in token"})
		return
	}

	// Ensure email is of the correct type
	userEmail, ok := email.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email format"})
		return
	}

	// Retrieve the user from the database using email
	var user models.User
	if err := config.DB.Where("email = ?", userEmail).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	// Get all cart items for this user
	var cartItems []models.Cart
	if err := config.DB.Where("user_id = ?", user.ID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

func RemoveFromCart(c *gin.Context) {
	var req struct {
		ProductID int `json:"product_id"`
	}

	// Bind incoming request data
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract email from context
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found in token"})
		return
	}

	// Ensure email is of the correct type
	userEmail, ok := email.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email format"})
		return
	}

	// Retrieve the user from the database using email
	var user models.User
	if err := config.DB.Where("email = ?", userEmail).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	// Delete the cart item for this user
	if err := config.DB.Where("user_id = ? AND product_id = ?", user.ID, req.ProductID).Delete(&models.Cart{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
