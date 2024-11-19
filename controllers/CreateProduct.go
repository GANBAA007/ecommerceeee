package controllers

import (
	"ecommerceeee/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// Decode the JSON body into the product struct
	if err :=c.Should

	// Save the product to the database


	// Return the created product as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
