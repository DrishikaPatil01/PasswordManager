package main

import (
	"github.com/gin-gonic/gin"
	"password-manager-service/handlers"
)

var router = gin.Default()

func main() {
	initRouter()
}

func initRouter() {
	router.GET("/health", handlers.HealthCheck)
	router.Run(":8080")
}
