package handlers

import (
	"fmt"
	"net/http"
	"password-manager-service/database"
	"password-manager-service/types"
	"password-manager-service/utils"

	"github.com/gin-gonic/gin"
)

func Login(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		var requestUser types.UserData

		if err := c.BindJSON(&requestUser); err != nil {
			c.JSON(http.StatusBadRequest, "could not process request")
			return
		}

		user, err := conn.GetUserByEmail(requestUser.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		requestUser.Password = utils.EncryptPassword(requestUser.Password)

		if user.Password != requestUser.Password {
			c.JSON(http.StatusBadRequest, "invalid email or password")
			return
		}

		sessionToken, err := conn.UpdateSession(user.UserId)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", sessionToken)
		c.JSON(http.StatusOK, user)
	}
	return gin.HandlerFunc(fn)
}
