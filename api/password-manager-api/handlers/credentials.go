package handlers

import (
	"fmt"
	"net/http"
	"password-manager-service/database"
	"password-manager-service/types"
	"password-manager-service/utils"

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
		isValid, userId := conn.ValidateSession(sessionToken)

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		signingKey := conn.GetSigningKey(userId)

		//Add Credentials
		encryptedPassword, err := utils.EncryptCredentialsPassword(credentials.Password, signingKey)

		if err != nil {
			fmt.Println("Error while encrypting password", err)
			c.JSON(http.StatusInternalServerError, "Error while storing password")
		}

		credentials.Password = encryptedPassword

		if err = conn.AddCredential(userId, &credentials); err != nil {
			fmt.Println("Error while adding credentials", err)
			c.JSON(http.StatusInternalServerError, "Error while adding credentials")
			return
		}

		//Update Session
		conn.UpdateSession(sessionToken)

		c.JSON(http.StatusCreated, "Added Credentials")
	}
	return gin.HandlerFunc(fn)
}

func UpdateCredential(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		var credentials types.CredentialData

		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, "could not process request")
			return
		}

		credentials.CredentialId = c.Param("id")

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, userId := conn.ValidateSession(sessionToken)

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		signingKey := conn.GetSigningKey(userId)

		//Update Credentials
		encryptedPassword, err := utils.EncryptCredentialsPassword(credentials.Password, signingKey)

		if err != nil {
			fmt.Println("Error while encrypting password", err)
			c.JSON(http.StatusInternalServerError, "Error while storing password")
		}

		credentials.Password = encryptedPassword

		if err = conn.UpdateCredential(&credentials); err != nil {
			fmt.Println("Error while updating credentials", err)
			c.JSON(http.StatusInternalServerError, "Error while updating credentials")
			return
		}

		//Update Session
		conn.UpdateSession(sessionToken)

		c.JSON(http.StatusCreated, "Updated Credentials")
	}
	return gin.HandlerFunc(fn)
}

func GetAllCredentials(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, userId := conn.ValidateSession(sessionToken)

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		//Get All Credentials
		credentials, err := conn.GetAllCredentials(userId)

		if err != nil {
			fmt.Println("Error while fetching credentials: ", err)
			c.JSON(http.StatusInternalServerError, "Error while fetching credentials")
			return
		}

		//Update Session
		conn.UpdateSession(userId)

		c.JSON(http.StatusCreated, credentials)
	}
	return gin.HandlerFunc(fn)
}

func GetCredentials(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		credentialId := c.Param("id")

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, _ := conn.ValidateSession(sessionToken)

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		//Get Credentials By Id
		credential, err := conn.GetCredential(credentialId)

		if err != nil {
			fmt.Println("Error while fetching credentials: ", err)
			c.JSON(http.StatusInternalServerError, "Error while fetching credentials")
			return
		}

		//Update Session
		conn.UpdateSession(sessionToken)

		if (types.CredentialData{}) == credential {
			c.JSON(http.StatusNoContent, "No credentials found with this id")
			return
		}

		c.JSON(http.StatusOK, credential)
	}
	return gin.HandlerFunc(fn)
}

func DeleteCredentialsById(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, userId := conn.ValidateSession(sessionToken)

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		var credentialIds []string

		if err := c.BindJSON(&credentialIds); err != nil {
			fmt.Println("Error while parsing request: ", err)
			c.JSON(http.StatusInternalServerError, "Error while processing request")
			return
		}

		//Delete Credentials By Id
		conn.DeleteCredential(credentialIds)

		//Update Session
		conn.UpdateSession(userId)

		c.JSON(http.StatusOK, "Deleted Credentials")
	}
	return gin.HandlerFunc(fn)
}
