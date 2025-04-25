package main

import (
	"log"

	middleware "ecommerceeee/Middleware"
	"ecommerceeee/config"
	"ecommerceeee/models"
	"ecommerceeee/routes"

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

	r.Use(middleware.CORSMiddleware())

	routes.SetupRoutes(r)
	routes.RegisterProductRoutes(r)

	r.Run(":8080")

}

// CreateInvoice
// Pay
// checkPayment
