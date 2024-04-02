package handlers

import (
	"net/http"
	"password-manager-service/database"
	"password-manager-service/types"
	"password-manager-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Signup(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		var user types.UserData

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, "could not process the data")
			return
		}

		ifExists, err := conn.CheckEmailExists(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "could not access the service")
			return
		}

		if ifExists {
			c.JSON(http.StatusNotAcceptable, "email id already exists")
			return
		}

		user.UserId = uuid.New().String()

		user.Password = utils.EncryptUserPassword(user.Password)
		if err := conn.AddUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, "error occured while adding user")
		}

		signingKey := utils.GenerateSigningKey()

		conn.CreateSigningKey(user.UserId, signingKey)

		c.JSON(http.StatusOK, "added user")

	}

	return gin.HandlerFunc(fn)
}
