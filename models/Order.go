package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint        `gorm:"not null"`              // Foreign key referencing User
	EmployeeID uint        `gorm:"not null"`              // Foreign key referencing Employee
	Paid       bool        `gorm:"default:false"`         // Indicates if the order is paid
	Shipped    bool        `gorm:"default:false"`         // Indicates if the order is shipped
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`    // One-to-many relationship with OrderItem
	Employee   Employee    `gorm:"foreignKey:EmployeeID"` // Define the relationship with Employee
}
