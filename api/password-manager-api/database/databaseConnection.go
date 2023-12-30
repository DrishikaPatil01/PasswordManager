package database

import (
	"database/sql"
	"fmt"
	"os"
)

type DatabaseConnection struct {
	db *sql.DB
}

func NewConnection() (DatabaseConnection, error) {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbLink := os.Getenv("DB_LINK")
	database := os.Getenv("DATABASE")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbLink, database))
	if err != nil {
		return DatabaseConnection{}, err
	}

	return DatabaseConnection{
		db: db,
	}, nil
}

func (conn *DatabaseConnection) Close() {
	conn.db.Close()
}
