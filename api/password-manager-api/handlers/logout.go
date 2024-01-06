package handlers

import (
	"fmt"
	"net/http"
	"password-manager-service/database"

	"github.com/gin-gonic/gin"
)

func Logout(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, _, err := conn.ValidateSession(sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
			return
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		//Delete Session
		if err = conn.DeleteSession(sessionToken); err != nil {
			fmt.Println("Error while deleting session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while logging out user")
			return
		}

		c.JSON(http.StatusOK, "Successfully logged out user")
	}
	return gin.HandlerFunc(fn)
}
