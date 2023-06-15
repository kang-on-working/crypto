package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"crypto/rand"
	"io"

)

func keyToHash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

// Decrypt 함수는 주어진 암호화된 문자열을 복호화하여 원래 문자열을 반환합니다.
func DecryptStr(encryptedStr string, key string) (string, error) {

	ciphertextBytes, err := base64.StdEncoding.DecodeString(encryptedStr)
	if err != nil {
		return "", err
	}
	HashKey := keyToHash(key)
	block, err := aes.NewCipher(HashKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertextBytes) < nonceSize {
		return "", errors.New("cipherText too short")
	}

	nonce, ciphertext := ciphertextBytes[:nonceSize], ciphertextBytes[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}


// Encrypt 함수는 주어진 원문 문자열을 암호화하여 암호화된 문자열을 반환합니다.
func EncryptStr(plainText string, key string) (string, error) {
	HashKey := keyToHash(key)
	block, err := aes.NewCipher(HashKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

