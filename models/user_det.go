package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Surname  string `json:"sur"`
	Name     string `json:"name"`
	Phone_no string `json:"phoneno"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}
