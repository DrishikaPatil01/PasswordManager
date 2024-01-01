package handlers

import (
	"fmt"
	"net/http"
	"password-manager-service/database"
	"password-manager-service/types"

	"github.com/gin-gonic/gin"
)

func AddCredential(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		var credentials types.CredentialData

		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, "could not process request")
			return
		}

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, err := conn.ValidateSession(credentials.UserId, sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
		} else {
			c.JSON(http.StatusOK, "Validated session successfully")
		}

		//Add Credentials

		newSessionToken, err := conn.UpdateSession(credentials.UserId)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", newSessionToken)
	}
	return gin.HandlerFunc(fn)
}
