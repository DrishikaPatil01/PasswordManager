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
	router.POST("/signup", handlers.Signup(&conn)) //done
	router.POST("/forgot-password", handlers.HealthCheck)
	router.PUT("/login", handlers.Login(&conn)) //done
	router.DELETE("/user/logout", handlers.HealthCheck)
	router.POST("/user/reset-password", handlers.HealthCheck)

	//Remove userId in path
	router.GET("/user/:userId/credentials/:id", handlers.GetCredentials(&conn))           //done
	router.GET("/user/:userId/credentials", handlers.GetAllCredentials(&conn))            //done
	router.DELETE("/user/:userId/credentials/:id", handlers.DeleteCredentialsById(&conn)) //convert to list
	router.DELETE("/user/:userId/credentials/", handlers.HealthCheck)
	router.PUT("/user/:userId/credentials/:id", handlers.UpdateCredential(&conn)) //done
	router.POST("/user/:userId/credentials", handlers.AddCredential(&conn))       //done

	//TODO: Remove this later
	router.GET("/users/getall", handlers.GetAllUsers(&conn))
	router.GET("/testSessionToken", handlers.TestSessionToken(&conn))

	router.Run(":8080")
}
