package models

import (
	"gorm.io/gorm"
)

type product struct {
	gorm.Model
	Name        string      `gorm:"not null"`
	Description string      `gorm:"type:text"`
	Price       float64     `gorm:"not null"`
	Stock       int         `gorm:"not null"`
	OrderItems  []OrderItem `gorm:"foreignKey:productID"` // One-to-many relationship with OrderItem
}
