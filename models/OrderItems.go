package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`             // Foreign key referencing Order
	ProductID uint    `gorm:"not null"`             // Foreign key referencing Product
	Quantity  int     `gorm:"not null"`             // Quantity of the product in the order
	Product   Product `gorm:"foreignKey:ProductID"` // Define the relationship with Product
}
