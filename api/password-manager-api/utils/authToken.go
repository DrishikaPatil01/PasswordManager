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
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("SIGNING_KEY")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}
