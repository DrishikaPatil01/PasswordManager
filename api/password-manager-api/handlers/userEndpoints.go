package handlers

import (
	"net/http"
	"password-manager-service/database"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(conn *database.DatabaseConnection) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usersData, err := conn.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "can't access the service")
		}

		c.JSON(http.StatusOK, usersData)
	}

	return gin.HandlerFunc(fn)
}
