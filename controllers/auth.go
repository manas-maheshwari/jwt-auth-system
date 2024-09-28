package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Logic for user login and JWT generation
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Register(c *gin.Context) {
	// Logic for user registration
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
