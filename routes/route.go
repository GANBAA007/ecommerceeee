package routes

import (
	"ecommerceeee/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/ecomm")
	api.POST("/AddProduct", controllers.CreateProduct)
	api.POST("/CreateUser", controllers.CreateUser)
	api.POST("/Login", controllers.Login)
	api.GET("/products", controllers.GetProd)
}
