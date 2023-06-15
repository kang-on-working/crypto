package main

import (
//	"bufio"
//	"os"
	"fmt"
	"pkg"
)

func main() {
	key := "My Example Key"

	git, err := pkg.JsonParse("./git.json")	
	if err != nil {
		fmt.Println("Json Parsing Error: ", err)
		return
	}
	fmt.Println(git)

	EncryptedGit, err := pkg.EncryptJsonValue(git, key)
	if err != nil {
		fmt.Println("Json Encrypt Error:", err)
		return
	}

	fmt.Println(EncryptedGit)

	DecryptedGit, err := pkg.DecryptJsonValue(EncryptedGit, key)
	if err != nil {
		fmt.Println("Json Decrypt Error:", err)
		return
	}

	fmt.Println(DecryptedGit)

	/*
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plain text: ")
	plainText, _ := reader.ReadString('\n')
	plainText = plainText[:len(plainText)-1] // 개행 문자 제거

	fmt.Print("Enter Encryption key: ")
	key, _ := reader.ReadString('\n')
	key = key[:len(key)-1] // 개행 문자 제거


	EncryptedText, err := pkg.EncryptStr(plainText, key)
	if err != nil {
		fmt.Println("Error Encrypting text:", err)
		return
	}

	fmt.Println("Encrypted text:", EncryptedText)

	DecryptedText, err := pkg.DecryptStr(EncryptedText, key)
	if err != nil {
		fmt.Println("Error Decrypting text:", err)
		return
	}

	fmt.Println("Decrypted text:", DecryptedText)

	*/
}
