package routes

import (
	"ecommerceeee/controllers"
	middleware "ecommerceeee/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes

	public := r.Group("/ecomm")
	{
		public.POST("/product", controllers.CreateProduct) // Endpoint for creating a product
		public.POST("/user", controllers.CreateUser)       // Endpoint for creating a user
		public.POST("/login", controllers.Login)
		public.GET("/products", controllers.GetProd) // Endpoint for user login
		public.GET("/getproducts/:id", controllers.GetProductById)

	}

	// Protected routes (require authentication)
	protected := r.Group("/ecomm")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.POST("/cart/add", controllers.AddToCart)           // Add product to cart
		protected.GET("/cart", controllers.GetCart)                  // Get user's cart
		protected.DELETE("/cart/remove", controllers.RemoveFromCart) // Remove product from cart
		protected.POST("/PlaceOrder", controllers.PlaceOrder)
	}
}
