package main

import (
	"password-manager-service/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default()) // remove this in production.

	router.GET("/health", handlers.HealthCheck)
	router.Run(":8080")
}
