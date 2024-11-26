package utility

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string) (string, error) {
	// Check if the email is not empty
	if email == "" {
		return "", fmt.Errorf("email cannot be empty")
	}

	// Define claims for the JWT token
	claims := jwt.MapClaims{
		"email": email,                                 // Store the email in the claims
		"exp":   time.Now().Add(time.Hour * 12).Unix(), // Expiry time set to 12 hours from now
	}

	// Fetch the secret key from environment variables
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("secret key is not set in environment")
	}

	// Create the JWT token using the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	// Return the signed token
	return signedToken, nil
}
