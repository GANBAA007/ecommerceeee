package models

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	UserID   uint
	Products []Products `gorm:"many2many:order_products" json:"products"`
	Total    int64      `json:"total"`
}
