package controllers

import (
	"ecommerceeee/config"
	"ecommerceeee/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found"})
		return
	}
	userEmail, ok := email.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email"})
		return
	}

	var user models.User
	if err := config.DB.Where("email=?", userEmail).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	var cartItems []models.Cart
	if err := config.DB.Where("user_id=?", user.ID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	order := models.Order{
		UserID: user.ID,
		Paid:   false,
	}
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	var TotalPrice float64
	for _, cartItem := range cartItems {
		var product models.Product
		if err := config.DB.Where("id = ?", cartItem.ProductID).First(&product).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}

		// log.Printf("Product ID: %d, Price: %f", product.ID, product.Price)
		itemTotal := float64(cartItem.Quantity) * product.Price
		TotalPrice += itemTotal

		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: uint(cartItem.ProductID),
			Quantity:  cartItem.Quantity,
			Product:   product,
		}

		if err := config.DB.Create(&orderItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to order"})
			return
		}
	}
	order.TotalPrice = TotalPrice
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order with total price"})
		return
	}
	if err := config.DB.Where("user_id = ?", user.ID).Delete(&models.Cart{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order placed successfully",
		"order": gin.H{
			"id":      order.ID,
			"user_id": order.UserID,
			"total":   order.TotalPrice,
		},
	})
}
