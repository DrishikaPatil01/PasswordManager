package utils

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId string) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(10 * time.Minute),
		"UserId":    userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return "", err
	}
	fmt.Println(ss, err)

	return ss, nil
}
