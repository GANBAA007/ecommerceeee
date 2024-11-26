package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint        `gorm:"not null"`
	EmployeeID uint        `gorm:"not null"`
	Paid       bool        `gorm:"default:false"`
	Shipped    bool        `gorm:"default:false"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	Employee   Employee    `gorm:"foreignKey:EmployeeID"`
}
