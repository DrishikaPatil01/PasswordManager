package database

import (
	"context"
	"log"
	"strings"

	"password-manager-service/types"
	"password-manager-service/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const (
	queryAllCredentials  string = "SELECT CREDENTIAL_ID, USERNAME, PASSWORD, OPTIONAL, TITLE FROM CREDENTIALS WHERE USER_ID=?"
	queryCredentialsById string = "SELECT CREDENTIAL_ID, USERNAME, PASSWORD, OPTIONAL, TITLE FROM CREDENTIALS WHERE CREDENTIAL_ID=?"
	insertCredential     string = "INSERT INTO CREDENTIALS (CREDENTIAL_ID, USER_ID, USERNAME, PASSWORD, OPTIONAL, TITLE) VALUES (?, ?, ?, ?, ?, ?)"
	updateCredential     string = "UPDATE CREDENTIALS SET USERNAME=?, PASSWORD=?, OPTIONAL=?, TITLE=? WHERE CREDENTIAL_ID=?"
	deleteCredential     string = "DELETE FROM CREDENTIALS WHERE CREDENTIAL_ID in(?)"
)

func (conn *DatabaseConnection) AddCredential(userId string, credentials *types.CredentialData) error {
	credentials.CredentialId = uuid.NewString()

	_, err := conn.db.ExecContext(context.Background(), insertCredential, credentials.CredentialId,
		userId, credentials.Username, credentials.Password, credentials.Optional, credentials.Title)

	return err
}

func (conn *DatabaseConnection) UpdateCredential(credentials *types.CredentialData) error {

	_, err := conn.db.ExecContext(context.Background(), updateCredential, credentials.Username,
		credentials.Password, credentials.Optional, credentials.Title, credentials.CredentialId)

	return err
}

func (conn *DatabaseConnection) GetAllCredentials(userId string) ([]types.CredentialData, error) {
	var credentials []types.CredentialData

	results, err := conn.db.Query(queryAllCredentials, userId)
	if err != nil {
		return credentials, err
	}

	signingKey := conn.GetSigningKey(userId)

	for results.Next() {
		var credential types.CredentialData

		err := results.Scan(&credential.CredentialId, &credential.Username, &credential.Password, &credential.Optional, &credential.Title)
		if err != nil {
			log.Printf("Could not process the row data: %s\n", err)
			break
		}

		credential.Password, err = utils.DecryptCredentialsPassword(credential.Password, signingKey)

		if err != nil {
			log.Printf("Could not decrypt the password: %s\n", err)
			break
		}

		credentials = append(credentials, credential)
	}

	return credentials, err
}

func (conn *DatabaseConnection) GetCredential(credentialId string) (types.CredentialData, error) {
	var credential types.CredentialData

	results, err := conn.db.Query(queryCredentialsById, credentialId)
	if err != nil {
		return credential, err
	}

	for results.Next() {
		err := results.Scan(&credential.CredentialId, &credential.Username, &credential.Password, &credential.Optional, &credential.Title)
		if err != nil {
			log.Printf("Could not process the row data: %s\n", err)
			break
		}

		signingKey := conn.GetSigningKey(credential.UserId)

		credential.Password, err = utils.DecryptCredentialsPassword(credential.Password, signingKey)

		if err != nil {
			log.Printf("Could not decrypt the password: %s\n", err)
			break
		}
	}

	return credential, err
}

func (conn *DatabaseConnection) DeleteCredential(credentialIds []string) error {

	_, err := conn.db.Query(deleteCredential, strings.Join(credentialIds, ","))

	return err
}
