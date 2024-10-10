package models

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model

	Number     string
	CustomerID uint
}
