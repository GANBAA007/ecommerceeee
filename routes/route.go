package routes

import (
	"ecommerceeee/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/Product", controllers.CreateProduct)
	api.POST("")
	api.POST("")
	api.POST("")
	api.POST("")
	api.POST("")
	api.POST("")
	api.POST("")
}
