package models

import (
	"time"
)

type Cart struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	ProductID int       `gorm:"not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
