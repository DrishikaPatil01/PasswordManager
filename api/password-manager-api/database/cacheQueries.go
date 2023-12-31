package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (conn *DatabaseConnection) UpdateSession(userId string) (string, error) {
	authToken := uuid.New().String()
	expiryTime := (time.Now().Add(10 * time.Minute)).Format(time.RFC850)

	val, err := conn.rdb.HSet(context.Background(),
		userId,
		[]string{"userId", userId, "sessionToken", authToken, "expiry", expiryTime}).Result()

	if err != nil {
		return "", err
	}

	fmt.Println("Val after updating session: ", val)

	return authToken, nil
}

func (conn *DatabaseConnection) ValidateSession(userId string, sessionId string) (bool, error) {
	session := conn.rdb.HGetAll(context.Background(), userId).Val()

	// jsonbody, err := json.Marshal(session)
	// if err != nil {
	// 	fmt.Println("Marshal error", err)
	// }

	// var sessionData types.SessionData
	// if err := json.Unmarshal(jsonbody, &sessionData); err != nil {
	// 	fmt.Println("Unmarshal error", err)
	// } else {
	// 	fmt.Println("Session Data ", sessionData)
	// }

	expiryTime, err := time.Parse(time.RFC850, session["expiry"])

	if err != nil {
		fmt.Println("Error while parsing time")
		return false, err
	}

	if userId == session["userId"] &&
		sessionId == session["sessionToken"] &&
		time.Now().Before(expiryTime) {
		return true, err
	}

	return false, nil
}
