package database

import (
	"context"

	"password-manager-service/types"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const (
	insertCredential string = "INSERT INTO CREDENTIALS (CREDENTIAL_ID, USER_ID, USERNAME, PASSWORD, OPTIONAL, TITLE) VALUES (?, ?, ?, ?, ?, ?)"
)

func (conn *DatabaseConnection) AddCredential(credentials *types.CredentialData) error {
	credentials.CredentialId = uuid.NewString()

	_, err := conn.db.ExecContext(context.Background(), insertCredential, credentials.CredentialId,
		credentials.UserId, credentials.Username, credentials.Password, credentials.Optional, credentials.Title)

	return err
}
