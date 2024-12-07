package main

import (
	"log"

	"ecommerceeee/config"
	"ecommerceeee/models"
	"ecommerceeee/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Migrate() {
	// Migrate the database schema for all models
	err := config.DB.AutoMigrate(&models.Admin{}, &models.Employee{}, &models.Product{}, &models.User{}, &models.Order{}, &models.OrderItem{})
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

	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	routes.SetupRoutes(r)

	r.Run(":8080")
}
