package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func EncryptUserPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func EncryptCredentialsPassword(plaintext string, secretKey string) (string, error) {

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

func DecryptCredentialsPassword(cipherstring string, secretKey string) (string, error) {
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

func GenerateSigningKey() string {
	signingKey := make([]rune, 32)
	for i := range signingKey {
		signingKey[i] = letters[rand.Intn(len(letters))]
	}
	return string(signingKey)
}
