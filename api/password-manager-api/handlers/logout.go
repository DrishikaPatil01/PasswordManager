package handlers

import (
	"net/http"
	"password-manager-service/database"
	"password-manager-service/types"

	"github.com/gin-gonic/gin"
)

func Logout(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		//Validate Session
		var headers types.Headers
		if err := c.ShouldBindHeader(&headers); err != nil {
			c.JSON(400, err.Error())
			return
		}
		isValid, _ := conn.ValidateSession(headers.SessionToken)

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		//Delete Session
		conn.DeleteSession(headers.SessionToken)

		c.JSON(http.StatusOK, "Successfully logged out user")
	}
	return gin.HandlerFunc(fn)
}
