package utils

import (
	"crypto/sha256"
	"fmt"
)

func EncryptPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
