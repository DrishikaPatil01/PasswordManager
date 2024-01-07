package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (conn *DatabaseConnection) CreateSession(userId string) (sessionId string) {
	sessionId = uuid.New().String()
	conn.rdb.Set(context.Background(), sessionId, userId, 10*time.Minute)

	return
}

func (conn *DatabaseConnection) UpdateSession(sessionId string) {
	conn.rdb.Expire(context.Background(), sessionId, 10*time.Minute)
}

func (conn *DatabaseConnection) ValidateSession(sessionId string) (bool, string) {
	userId := conn.rdb.Get(context.Background(), sessionId).Val()

	if userId == "" {
		return false, ""
	}

	return true, userId
}

func (conn *DatabaseConnection) DeleteSession(sessionId string) {
	conn.rdb.Del(context.Background(), sessionId)
}
