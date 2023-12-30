package handlers

import (
	"log"
	"net/http"
	"password-manager-service/database"
	"password-manager-service/types"
	"password-manager-service/utils"

	"github.com/gin-gonic/gin"
)

func Login(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		var user types.UserData

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, "could not process the data")
			return
		}

		//Check if emailExists and encrypted password matches this password
		userPresent, err := conn.GetUserByEmail(user.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		user.Password = utils.EncryptPassword(user.Password)

		if user.Password != userPresent.Password {
			c.JSON(http.StatusBadRequest, "invalid email or password")
			return
		}

		//Add/Update Session details
		token, err := utils.CreateToken(userPresent.UserId)

		if err != nil {
			log.Println("Error while generating token", err)
		}

		c.Writer.Header().Set("auth_token", token)
		c.JSON(http.StatusOK, userPresent)
	}

	return gin.HandlerFunc(fn)
}
