package main

import (
	"log"

	"ecommerceeee/config"
	"ecommerceeee/models"
	"ecommerceeee/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Migrate() {
	// Migrate the database schema for all models
	err := config.DB.AutoMigrate(&models.Admin{}, &models.Product{}, &models.User{}, &models.Order{}, &models.OrderItem{}, &models.Cart{})
	if err != nil {
		log.Fatalf("Migrating failed: %v", err)
	} else {
		log.Println("Migration successful")
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	Migrate()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Replace with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Allow credentials like cookies or authorization headers
	}))

	routes.SetupRoutes(r)

	r.Run(":8080")
}

// CreateInvoice
// Pay
// checkPayment
