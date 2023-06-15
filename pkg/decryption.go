package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// Decrypt 함수는 주어진 암호화된 문자열을 복호화하여 원래 문자열을 반환합니다.
func DecryptStr(encryptedStr string, key string) (string, error) {
	HashKey := KeyToHash(key)

	ciphertextBytes, err := base64.StdEncoding.DecodeString(encryptedStr)
	if err != nil {
		return "", err
	}
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
		return "", fmt.Errorf("cipherText too short")
	}

	nonce, ciphertext := ciphertextBytes[:nonceSize], ciphertextBytes[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}



func DecryptJsonValue(git GIT, key string) (git_decrypted GIT, err error) {
	decryptedEmail, err := DecryptStr(git.Email, key)
	if err != nil {
		return GIT{}, err
	}

	decryptedUsername, err := DecryptStr(git.Username, key)
	if err != nil {
		return GIT{}, err
	}

	decryptedToken, err := DecryptStr(git.Token, key)
	if err != nil {
		return GIT{}, err
	}

	decryptedRepo, err := DecryptStr(git.Repo, key)
	if err != nil {
		return GIT{}, err
	}

	return GIT{
		Email:    decryptedEmail,
		Username: decryptedUsername,
		Token:    decryptedToken,
		Repo:     decryptedRepo,
	}, nil
}
