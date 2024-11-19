package main

import (
	"encoding/json"
	"log"

	"net/http"
	"os"

	"ecommerceeee/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.Admin{}, &models.Employee{}, &models.Order{}, &models.OrderItem{}, &models.Product{}, &models.User{})
}

func main() {
	log.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}

}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result := DB.Create(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateInvoices(w http.ResponseWriter, r *http.Request) {
	var invoice models.Order

	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result := DB.Create(&invoice); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}
