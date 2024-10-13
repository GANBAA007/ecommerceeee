package models

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category"`
}
