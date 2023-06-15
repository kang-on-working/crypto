package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"crypto/rand"
	"io"
)
// Encrypt 함수는 주어진 원문 문자열을 암호화하여 암호화된 문자열을 반환합니다.
func EncryptStr(plainText string, key string) (string, error) {
	HashKey := KeyToHash(key)
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

// EncryptJsonValue 함수는 주어진 git 구조체의 각각의 필드 값을 암호화하여 새로운 git_encrypted 구조체를 반환합니다.
func EncryptJsonValue(git GIT, key string) (git_encrypted GIT, err error) {
	encryptedEmail, err := EncryptStr(git.Email, key)
	if err != nil {
		return GIT{}, err
	}

	encryptedUsername, err := EncryptStr(git.Username, key)
	if err != nil {
		return GIT{}, err
	}

	encryptedToken, err := EncryptStr(git.Token, key)
	if err != nil {
		return GIT{}, err
	}

	encryptedRepo, err := EncryptStr(git.Repo, key)
	if err != nil {
		return GIT{}, err
	}

	return GIT{
		Email:    encryptedEmail,
		Username: encryptedUsername,
		Token:    encryptedToken,
		Repo:     encryptedRepo,
	}, nil
}

