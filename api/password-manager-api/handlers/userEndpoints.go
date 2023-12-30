package handlers

import (
	"net/http"
	"password-manager-service/database"
	"password-manager-service/utils"

	"log"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(conn *database.DatabaseConnection) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usersData, err := conn.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "can't access the service")
		}

		c.JSON(http.StatusOK, usersData)
	}

	return gin.HandlerFunc(fn)
}

func TestAuthToken(conn *database.DatabaseConnection) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.GetHeader("Auth_Token")

		validatedToken, err := utils.ValidateJWT(token)

		if err != nil {
			log.Print("Error while validating token :", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
			return
		}
		if !validatedToken.Valid {
			log.Print("Invalid token :", err)
			c.JSON(http.StatusUnauthorized, "Invalid Token")
			return
		}

		c.JSON(http.StatusOK, "Tested authToken successfully!")
	}

	return gin.HandlerFunc(fn)
}
