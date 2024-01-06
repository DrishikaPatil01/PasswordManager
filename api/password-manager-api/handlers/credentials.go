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
		isValid, userId, err := conn.ValidateSession(sessionToken)

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

		if err != nil {
			fmt.Println("Error while encrypting password", err)
			c.JSON(http.StatusInternalServerError, "Error while storing password")
		}

		credentials.Password = encryptedPassword

		err = conn.AddCredential(userId, &credentials)

		if err != nil {
			fmt.Println("Error while adding credentials", err)
			c.JSON(http.StatusInternalServerError, "Error while adding credentials")
			return
		}

		//Update Session
		err = conn.UpdateSession(sessionToken)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", sessionToken)
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
		isValid, _, err := conn.ValidateSession(sessionToken)

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

		if err = conn.UpdateCredential(&credentials); err != nil {
			fmt.Println("Error while updating credentials", err)
			c.JSON(http.StatusInternalServerError, "Error while updating credentials")
			return
		}

		//Update Session
		if err = conn.UpdateSession(sessionToken); err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", sessionToken)
		c.JSON(http.StatusCreated, "Updated Credentials")
	}
	return gin.HandlerFunc(fn)
}

func GetAllCredentials(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, userId, err := conn.ValidateSession(sessionToken)

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
		if err := conn.UpdateSession(userId); err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", sessionToken)
		c.JSON(http.StatusCreated, credentials)
	}
	return gin.HandlerFunc(fn)
}

func GetCredentials(conn *database.DatabaseConnection) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		credentialId := c.Param("id")

		//Validate Session
		sessionToken := c.GetHeader("SessionToken")
		isValid, _, err := conn.ValidateSession(sessionToken)

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
		if err = conn.UpdateSession(sessionToken); err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", sessionToken)

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
		isValid, userId, err := conn.ValidateSession(sessionToken)

		if err != nil {
			fmt.Println("Error while validating session error:", err)
			c.JSON(http.StatusInternalServerError, "Error while validating token")
			return
		}

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
		err = conn.DeleteCredential(credentialIds)

		if err != nil {
			fmt.Println("Error while deleting credentials: ", err)
			c.JSON(http.StatusInternalServerError, "Error while fetching credentials")
			return
		}

		//Update Session
		err = conn.UpdateSession(userId)

		if err != nil {
			fmt.Println("Error while Creating session :", err)
			c.JSON(http.StatusInternalServerError, "Error creating session")
			return
		}

		c.Writer.Header().Set("SessionToken", sessionToken)
		c.JSON(http.StatusOK, "Deleted Credentials")
	}
	return gin.HandlerFunc(fn)
}
