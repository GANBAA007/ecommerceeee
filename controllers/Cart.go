package controllers

// import (
// 	"ecommerceeee/config"
// 	"ecommerceeee/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func AddToCart(c *gin.Context) {
// 	var req struct {
// 		ProductID int `json:"product_id"`
// 		Quantity  int `json:"quantity"`
// 	}

// 	// Bind request data
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Get user ID (you can get this from JWT token, for example)
// 	userID := 1 // Assume userID is extracted from a valid JWT token

// 	// Check if product exists
// 	var product models.Product
// 	if err := config.DB.Where("id = ?", req.ProductID).First(&product).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
// 		return
// 	}

// 	// Check if cart item already exists for this user and product
// 	var cartItem models.Cart
// 	if err := config.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).First(&cartItem).Error; err == nil {
// 		// Product already in cart, update quantity
// 		cartItem.Quantity += req.Quantity
// 		if err := config.DB.Save(&cartItem).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart"})
// 			return
// 		}
// 	} else {
// 		// Add new product to cart
// 		newCartItem := models.Cart{
// 			UserID:    userID,
// 			ProductID: req.ProductID,
// 			Quantity:  req.Quantity,
// 		}
// 		if err := config.DB.Create(&newCartItem).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
// }

// func GetCart(c *gin.Context) {
// 	// Get user ID (again, probably from the JWT token)
// 	userID := 1

// 	// Get all cart items for this user
// 	var cartItems []models.Cart
// 	if err := config.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, cartItems)
// }

// func RemoveFromCart(c *gin.Context) {
// 	var req struct {
// 		ProductID int `json:"product_id"`
// 	}

// 	// Bind request data
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Get user ID (from JWT token, for example)
// 	userID := 1

// 	// Delete cart item
// 	if err := config.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).Delete(&models.Cart{}).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
// }
