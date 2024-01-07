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
	router.POST("/forgot-password", handlers.ForgotPassword(&conn))
	router.PUT("/login", handlers.Login(&conn))
	router.DELETE("/user/logout", handlers.Logout(&conn))
	router.POST("/user/reset-password", handlers.ResetPassword(&conn))

	//Remove userId in path
	router.GET("/user/credentials/:id", handlers.GetCredentials(&conn))
	router.GET("/user/credentials", handlers.GetAllCredentials(&conn))
	router.DELETE("/user/credentials/", handlers.DeleteCredentialsById(&conn))
	router.PUT("/user/credentials/:id", handlers.UpdateCredential(&conn))
	router.POST("/user/credentials", handlers.AddCredential(&conn))

	//TODO: Remove this later
	router.GET("/users/getall", handlers.GetAllUsers(&conn))
	router.GET("/testSessionToken", handlers.TestSessionToken(&conn))

	router.Run(":8080")
}
