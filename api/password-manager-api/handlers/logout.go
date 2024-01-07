package handlers

import (
	"net/http"
	"password-manager-service/database"

	"github.com/gin-gonic/gin"
)

func Logout(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, _ := conn.ValidateSession(sessionToken)

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		//Delete Session
		conn.DeleteSession(sessionToken)

		c.JSON(http.StatusOK, "Successfully logged out user")
	}
	return gin.HandlerFunc(fn)
}
