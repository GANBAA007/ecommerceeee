package controllers

import (
	"ecommerceeee/config"
	"ecommerceeee/models"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ExistingUser models.User
	err := config.DB.Where("ID=?", user.ID).First(&ExistingUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		} else {
			log.Printf("error occurred during query:%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error DB"})
			return
		}
	} else {
		c.JSON(http.StatusConflict, gin.H{"error": "exsting user"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("encrypting error&v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := config.DB.Create(&user).Error; err != nil {
		log.Printf("error creating user:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating user"})
		return
	}
}
