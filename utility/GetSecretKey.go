package utility

import (
	"fmt"
	"os"
)

// GetSecretKey retrieves the JWT secret key from the environment variable.
func GetSecretKey() string {
	// Fetch the secret key from the environment
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		// Handle error if the SECRET_KEY is not set
		fmt.Println("Error: SECRET_KEY is not set in the environment variables")
	}
	return secretKey
}
