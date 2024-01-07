package handlers

import (
	"net/http"
	"password-manager-service/database"
	"password-manager-service/types"
	"password-manager-service/utils"

	"github.com/gin-gonic/gin"
)

func ForgotPassword(conn *database.DatabaseConnection) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		//checkIf user exists by email
		var requestUser types.UserData

		if err := c.BindJSON(&requestUser); err != nil {
			c.JSON(http.StatusBadRequest, "could not process the data")
			return
		}

		user, err := conn.GetUserByEmail(requestUser.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if user == (types.UserData{}) {
			c.JSON(http.StatusNotFound, "Email not found")
			return
		}

		user.Password = utils.EncryptUserPassword(requestUser.Password)

		//reset password sent in request
		if err := conn.UpdateUserPassword(user.UserId, user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, "error occured while resetting password")
			return
		}

		c.JSON(http.StatusOK, "Successfully reset password")
	}

	return gin.HandlerFunc(fn)
}
