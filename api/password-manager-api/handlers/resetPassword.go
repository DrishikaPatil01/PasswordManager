package handlers

import (
	"net/http"
	"password-manager-service/database"
	"password-manager-service/types"
	"password-manager-service/utils"

	"github.com/gin-gonic/gin"
)

func ResetPassword(conn *database.DatabaseConnection) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user types.UserData

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, "could not process the data")
			return
		}

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, userId := conn.ValidateSession(sessionToken)

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		user.Password = utils.EncryptUserPassword(user.Password)

		//reset password sent in request
		if err := conn.UpdateUserPassword(userId, user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, "error occured while resetting password")
			return
		}

		//Update Session
		conn.UpdateSession(sessionToken)

		c.Writer.Header().Set("SessionToken", sessionToken)
		c.JSON(http.StatusOK, "Successfully reset password")
	}

	return gin.HandlerFunc(fn)
}
