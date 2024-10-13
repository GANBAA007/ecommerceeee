package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/GANBAA007/ecommerceeee/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up the MySQL database connection
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the Product model
	DB.AutoMigrate(&models.Products{})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", createProduct).Methods("POST")

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Products

	// Decode the JSON body into the product struct
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the product to the database
	if result := DB.Create(&product); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created product as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
