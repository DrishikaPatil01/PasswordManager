package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Extend expiry instead of creating new one
func (conn *DatabaseConnection) UpdateSession(userId string) (string, error) {
	authToken := uuid.New().String()
	expiryTime := (time.Now().Add(10 * time.Minute)).Format(time.RFC3339)

	_, err := conn.rdb.HSet(context.Background(),
		userId,
		[]string{"userId", userId, "sessionToken", authToken, "expiry", expiryTime}).Result()

	if err != nil {
		return "", err
	}
	return authToken, nil
}

func (conn *DatabaseConnection) ValidateSession(userId string, sessionId string) (bool, error) {
	session := conn.rdb.HGetAll(context.Background(), userId).Val()

	expiryTime, err := time.Parse(time.RFC3339, session["expiry"])
	if err != nil {
		fmt.Println("Error while parsing time")
		return false, err
	}

	//REmove userId mapping and set sessionToken as primary key, only check expiry
	if userId == session["userId"] &&
		sessionId == session["sessionToken"] &&
		time.Now().Before(expiryTime) {
		return true, err
	}

	return false, nil
}
