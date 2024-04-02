package handlers

import (
	"fmt"
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

		//create new sessionToken
		sessionToken := conn.CreateSession(user.UserId)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		utils.SendResetPasswordEmail(user.Email, sessionToken)

		c.JSON(http.StatusOK, "Successfully sent email too reset password")
	}

	return gin.HandlerFunc(fn)
}
