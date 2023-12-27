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
	router.PUT("/signup", handlers.HealthCheck)
	router.POST("/forgot-password", handlers.HealthCheck)
	router.PUT("/login", handlers.HealthCheck)
	router.DELETE("/user/logout", handlers.HealthCheck)
	router.POST("/user/reset-password", handlers.HealthCheck)

	router.GET("/user/credentials", handlers.HealthCheck)
	router.DELETE("/user/credentials", handlers.HealthCheck)
	router.PUT("/user/credentials", handlers.HealthCheck)
	router.POST("/user/credentials", handlers.HealthCheck)

	router.Run(":8080")
}
