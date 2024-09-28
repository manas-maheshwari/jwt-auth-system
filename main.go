package main

import (
    "github.com/gin-gonic/gin"
    "jwt-auth-system/controllers"
)

func main() {
    router := gin.Default()

    router.POST("/login", controllers.Login)
    router.POST("/register", controllers.Register)

    // Protect routes using middleware here

    router.Run(":8080")
}
