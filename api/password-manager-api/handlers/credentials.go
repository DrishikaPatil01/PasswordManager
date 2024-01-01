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

		userId := c.Param("userId")

		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, "could not process request")
			return
		}

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, err := conn.ValidateSession(userId, sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
			return
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		//Add Credentials
		encryptedPassword, err := utils.EncryptCredentialsPassword(credentials.Password)
		credentials.UserId = userId

		if err != nil {
			fmt.Println("Error while encrypting password", err)
			c.JSON(http.StatusInternalServerError, "Error while storing password")
		}

		credentials.Password = encryptedPassword

		err = conn.AddCredential(&credentials)

		if err != nil {
			fmt.Println("Error while adding credentials", err)
			c.JSON(http.StatusInternalServerError, "Error while adding credentials")
			return
		}

		//Update Session
		newSessionToken, err := conn.UpdateSession(credentials.UserId)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", newSessionToken)
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
		isValid, err := conn.ValidateSession(credentials.UserId, sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
			return
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		//Update Credentials
		encryptedPassword, err := utils.EncryptCredentialsPassword(credentials.Password)

		if err != nil {
			fmt.Println("Error while encrypting password", err)
			c.JSON(http.StatusInternalServerError, "Error while storing password")
		}

		credentials.Password = encryptedPassword

		err = conn.UpdateCredential(&credentials)

		if err != nil {
			fmt.Println("Error while updating credentials", err)
			c.JSON(http.StatusInternalServerError, "Error while updating credentials")
			return
		}

		//Update Session
		newSessionToken, err := conn.UpdateSession(credentials.UserId)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", newSessionToken)
		c.JSON(http.StatusCreated, "Updated Credentials")
	}
	return gin.HandlerFunc(fn)
}

func GetAllCredentials(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		userId := c.Param("userId")

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, err := conn.ValidateSession(userId, sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
			return
		}

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
		newSessionToken, err := conn.UpdateSession(userId)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", newSessionToken)
		c.JSON(http.StatusCreated, credentials)
	}
	return gin.HandlerFunc(fn)
}

func GetCredentials(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		userId := c.Param("userId")
		credentialId := c.Param("id")

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, err := conn.ValidateSession(userId, sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
			return
		}

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
		newSessionToken, err := conn.UpdateSession(userId)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", newSessionToken)

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
		userId := c.Param("userId")
		credentialId := c.Param("id")

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, err := conn.ValidateSession(userId, sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
			return
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, "Invalid Session token")
			return
		}

		//Delete Credentials By Id
		err = conn.DeleteCredential(credentialId)

		if err != nil {
			fmt.Println("Error while deleting credentials: ", err)
			c.JSON(http.StatusInternalServerError, "Error while fetching credentials")
			return
		}

		//Update Session
		newSessionToken, err := conn.UpdateSession(userId)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", newSessionToken)
		c.JSON(http.StatusOK, "Deleted Credentials")
	}
	return gin.HandlerFunc(fn)
}
