package main

import (
//	"bufio"
//	"os"
	"fmt"
	"pkg"
)

func main() {
	git, err := pkg.JsonParse("./git.json")	
	if err != nil {
		fmt.Println("Json Parsing Error: ", err)
		return
	}
	fmt.Println(git)

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
		fmt.Println("Error EncryptString text:", err)
		return
	}

	fmt.Println("EncryptStred text:", EncryptedText)

	DecryptedText, err := pkg.DecryptStr(EncryptedText, key)
	if err != nil {
		fmt.Println("Error DecryptString text:", err)
		return
	}

	fmt.Println("DecryptStred text:", DecryptedText)

	*/
}
