package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plain text: ")
	plainText, _ := reader.ReadString('\n')
	plainText = plainText[:len(plainText)-1] // 개행 문자 제거

	fmt.Print("Enter encryption key: ")
	key, _ := reader.ReadString('\n')
	key = key[:len(key)-1] // 개행 문자 제거


	encryptedText, err := pkg.Encrypt(plainText, key)
	if err != nil {
		fmt.Println("Error encrypting text:", err)
		return
	}

	fmt.Println("Encrypted text:", encryptedText)

	decryptedText, err := pkg.Decrypt(encryptedText, key)
	if err != nil {
		fmt.Println("Error Decrypting text:", err)
		return
	}

	fmt.Println("Decrypted text:", decryptedText)
}
