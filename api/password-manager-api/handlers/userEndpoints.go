package handlers

import (
	"fmt"
	"net/http"
	"password-manager-service/database"
	"password-manager-service/types"

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

func TestSessionToken(conn *database.DatabaseConnection) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var requestUser types.UserData

		if err := c.BindJSON(&requestUser); err != nil {
			c.JSON(http.StatusBadRequest, "could not process request")
			return
		}

		sessionToken := c.GetHeader("SessionToken")
		fmt.Println(sessionToken)

		isValid, _, err := conn.ValidateSession(sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
		} else {
			c.JSON(http.StatusOK, "Validated session successfully")
		}
	}
	return gin.HandlerFunc(fn)
}
