package main

import (
	"github.com/gin-gonic/gin"

	"password-manager-service/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/health", handlers.HealthCheck)
	router.Run("localhost:8080")
}
