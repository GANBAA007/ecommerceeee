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
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Set up the MySQL database connection
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the Product model
	DB.AutoMigrate(&models.Product{})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/invoices-Create", CreateInvoices).Methods("POST")
	r.HandleFunc("/users", createUser).Methods("POST")

	log.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

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
func createUser(w http.ResponseWriter, r *http.Request) { // Corrected function name
	var user models.User

	// Decode the JSON body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the user to the database
	if result := DB.Create(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created user as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateInvoices(w http.ResponseWriter, r *http.Request) {
	var invoice models.Order

	// Decode the JSON body into the invoice struct
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the invoice to the database
	if result := DB.Create(&invoice); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created invoice as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}
