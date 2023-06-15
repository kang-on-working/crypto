package pkg

import (
	"os"
	"bufio"
	"fmt"
	"crypto/sha256"
)


func KeyToHash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

func GetStrKey() string {
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
	return key
}
