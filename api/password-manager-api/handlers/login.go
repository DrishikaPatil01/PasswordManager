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

		requestUser := validate(c)

		if (types.UserData{}) == requestUser {
			return
		}

		//Check if emailExists and encrypted password matches this password
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

		token, err := utils.CreateToken(user.UserId)

		if err != nil {
			log.Println("Error while generating token", err)
			c.JSON(http.StatusInternalServerError, "Error processing request")
			return
		} else {
			c.Writer.Header().Set("auth_token", token)
			c.JSON(http.StatusOK, user)
		}
	}
	return gin.HandlerFunc(fn)
}

func validate(c *gin.Context) (user types.UserData) {

	//validate
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "could not process request")
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, "Email is required to login user")
		return types.UserData{}
	}

	if user.Password == "" {
		c.JSON(http.StatusBadRequest, "Password is required to login user")
		return types.UserData{}
	}

	return
}
