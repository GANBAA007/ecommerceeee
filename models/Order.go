package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint    `gorm:"not null"`
	Paid       bool    `gorm:"default:false"`
	TotalPrice float64 `gorm:"default:0"`
}
