package Middleware

import (
	"ecommerceeee/utility"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token required"})
			c.Abort()
			return
		}

		// Split the token from the 'bearer' keyword in the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // Means no 'Bearer' keyword found
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token required"})
			c.Abort()
			return
		}

		// Parse the token
		claims := &jwt.MapClaims{}
		secretKey := []byte(utility.GetSecretKey()) // Fetch the secret key used to sign the JWT
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Return the secret key used to sign the JWT
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired or invalid"})
			c.Abort()
			return
		}

		// Extract the email from the claims
		userEmail, ok := (*claims)["email"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store the email in the context
		c.Set("email", userEmail)

		// Proceed to the next handler
		c.Next()
	}
}
