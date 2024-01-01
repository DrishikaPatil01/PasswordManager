package database

import (
	"context"
	"log"

	"password-manager-service/types"
	"password-manager-service/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const (
	queryAllCredentials  string = "SELECT * FROM CREDENTIALS WHERE USER_ID=?"
	queryCredentialsById string = "SELECT * FROM CREDENTIALS WHERE CREDENTIAL_ID=?"
	insertCredential     string = "INSERT INTO CREDENTIALS (CREDENTIAL_ID, USER_ID, USERNAME, PASSWORD, OPTIONAL, TITLE) VALUES (?, ?, ?, ?, ?, ?)"
	updateCredential     string = "UPDATE CREDENTIALS SET USERNAME=?, PASSWORD=?, OPTIONAL=?, TITLE=? WHERE CREDENTIAL_ID=?"
	deleteCredential     string = "DELETE FROM CREDENTIALS WHERE CREDENTIAL_ID=?"
)

func (conn *DatabaseConnection) AddCredential(credentials *types.CredentialData) error {
	credentials.CredentialId = uuid.NewString()

	_, err := conn.db.ExecContext(context.Background(), insertCredential, credentials.CredentialId,
		credentials.UserId, credentials.Username, credentials.Password, credentials.Optional, credentials.Title)

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

	for results.Next() {
		var credential types.CredentialData

		err := results.Scan(&credential.CredentialId, &credential.UserId, &credential.Username, &credential.Password, &credential.Optional, &credential.Title)
		if err != nil {
			log.Printf("Could not process the row data: %s\n", err)
		}

		credential.Password, err = utils.DecryptCredentialsPassword(credential.Password)

		if err != nil {
			log.Printf("Could not decrypt the password: %s\n", err)
		}

		credentials = append(credentials, credential)
	}

	return credentials, nil
}

func (conn *DatabaseConnection) GetCredential(credentialId string) (types.CredentialData, error) {
	var credential types.CredentialData

	results, err := conn.db.Query(queryCredentialsById, credentialId)
	if err != nil {
		return credential, err
	}

	for results.Next() {
		err := results.Scan(&credential.CredentialId, &credential.UserId, &credential.Username, &credential.Password, &credential.Optional, &credential.Title)
		if err != nil {
			log.Printf("Could not process the row data: %s\n", err)
		}

		credential.Password, err = utils.DecryptCredentialsPassword(credential.Password)

		if err != nil {
			log.Printf("Could not decrypt the password: %s\n", err)
			return credential, err
		}
	}

	return credential, nil
}

func (conn *DatabaseConnection) DeleteCredential(credentialId string) error {

	_, err := conn.db.Query(deleteCredential, credentialId)

	return err
}
