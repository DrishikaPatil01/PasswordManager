package main

import (
	"log"
	"password-manager-service/database"
	"password-manager-service/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	conn, err := database.NewConnection()
	if err != nil {
		log.Println("Could not establish a connection")
	}
	defer conn.Close()

	router := gin.Default()

	//TODO: remove this in production.
	router.Use(cors.Default())

	router.GET("/health", handlers.HealthCheck)
	router.POST("/signup", handlers.Signup(&conn))
	router.POST("/forgot-password", handlers.HealthCheck)
	router.PUT("/login", handlers.HealthCheck)
	router.DELETE("/user/logout", handlers.HealthCheck)
	router.POST("/user/reset-password", handlers.HealthCheck)

	router.GET("/user/credentials", handlers.HealthCheck)
	router.DELETE("/user/credentials", handlers.HealthCheck)
	router.PUT("/user/credentials", handlers.HealthCheck)
	router.POST("/user/credentials", handlers.HealthCheck)

	//TODO: Remove this later
	router.GET("/users/getall", handlers.GetAllUsers(&conn))

	router.Run(":8080")
}
