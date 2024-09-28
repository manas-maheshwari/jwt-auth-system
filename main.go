package main

import (
	"jwt-auth-system/controllers"
	"jwt-auth-system/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/dashboard", func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(200, gin.H{"message": "Welcome to the dashboard, " + username})
	})

	r.Run(":8080") // Start the server on port 8080
}
