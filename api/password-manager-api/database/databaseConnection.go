package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

type DatabaseConnection struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewConnection() (DatabaseConnection, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbLink := os.Getenv("DB_LINK")
	database := os.Getenv("DATABASE")
	redisLink := os.Getenv("REDIS_LINK")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	db, dbErr := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbLink, database))
	if dbErr != nil {
		return DatabaseConnection{}, dbErr
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisLink,
		Password: redisPassword,
		DB:       0,
	})

	return DatabaseConnection{
		db:  db,
		rdb: rdb,
	}, nil
}

func (conn *DatabaseConnection) Close() {
	conn.db.Close()
}
