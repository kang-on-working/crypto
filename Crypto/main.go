package main

import (
	"bufio"
	"os"
	"fmt"
	"pkg"
)

func main() {

	var pnt *string
	if len(os.Args) < 2 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Encryption key: ")
		key, _ := reader.ReadString('\n')
		key = key[:len(key)-1] // 개행 문자 제거
		pnt = &key	

	} else { 
		var param string = os.Args[1]
		pnt = &param
	}

	key := *pnt

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

}
