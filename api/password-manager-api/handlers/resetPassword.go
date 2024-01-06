package handlers

import (
	"fmt"
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
		isValid, userId, err := conn.ValidateSession(sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
			return
		}

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
		if err = conn.UpdateSession(sessionToken); err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", sessionToken)
		c.JSON(http.StatusOK, "Successfully reset password")
	}

	return gin.HandlerFunc(fn)
}
