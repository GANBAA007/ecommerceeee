package models

import (
	"gorm.io/gorm"
)

type products struct {
	gorm.Model
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category"`
}
