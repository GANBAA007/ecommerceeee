package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine) {
	router.GET("/product/:id", func(c *gin.Context) {
		c.File("./path/to/product.html")
	})
}
