package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

var secretKey string = os.Getenv("SECRET_KEY")

func EncryptUserPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func EncryptCredentialsPassword(plaintext string) (string, error) {

	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	//GCM : AES works only for fixed size strings, GCM Mode handles this issue
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	cipherstring := base64.StdEncoding.EncodeToString(ciphertext)

	fmt.Println("EncryptedPassword :", cipherstring)

	return cipherstring, nil
}

func DecryptCredentialsPassword(cipherstring string) (string, error) {
	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		fmt.Println("Error while forming aes cipher")
		return "", err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		fmt.Println("Error while forming GCM")
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(cipherstring)

	if err != nil {
		fmt.Println("Error while decoding")
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		fmt.Println("Error while decrypting")
		return "", err
	}

	return string(plaintext), nil
}
