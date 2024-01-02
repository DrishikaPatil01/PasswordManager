package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (conn *DatabaseConnection) CreateSession(userId string) (string, error) {
	sessionToken := uuid.New().String()
	expiryTime := (time.Now().Add(10 * time.Minute)).Format(time.RFC3339)

	val, err := conn.rdb.HSet(context.Background(),
		sessionToken,
		[]string{"userId", userId, "expiry", expiryTime}).Result()

	if err != nil {
		return "", err
	}

	fmt.Println("Val after updating session: ", val)

	return sessionToken, nil
}

func (conn *DatabaseConnection) UpdateSession(sessionId string) (string, error) {
	expiryTime := (time.Now().Add(10 * time.Minute)).Format(time.RFC3339)

	_, err := conn.rdb.HSet(context.Background(),
		sessionId,
		[]string{"expiry", expiryTime}).Result()

	return "", err
}

func (conn *DatabaseConnection) ValidateSession(sessionId string) (bool, string, error) {
	session := conn.rdb.HGetAll(context.Background(), sessionId).Val()

	expiryTime, err := time.Parse(time.RFC3339, session["expiry"])
	if err != nil {
		fmt.Println("Error while parsing time")
		return false, "", err
	}

	if time.Now().Before(expiryTime) {
		return true, session["userId"], err
	}

	return false, "", nil
}
