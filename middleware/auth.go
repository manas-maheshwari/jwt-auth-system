package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to verify JWT token

		// If valid, proceed
		c.Next()

		// Otherwise, block the request
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
