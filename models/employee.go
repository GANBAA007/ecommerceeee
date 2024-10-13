package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name    string  `gorm:"not null"`
	Surname string  `gorm:"not null"`
	Email   string  `gorm:"unique;not null"`
	Phone   string  `gorm:"not null"`
	Orders  []Order `gorm:"foreignKey:EmployeeID"`
}
