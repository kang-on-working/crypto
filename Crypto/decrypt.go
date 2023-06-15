package main

import (
	"fmt"
	"io/ioutil"
	"pkg"
)

func main() {

	JsonFile := "./git.json"

	key := pkg.GetStrKey()

	git, err := pkg.JsonParse(JsonFile)	
	if err != nil {
		fmt.Println("Json Parsing Error: ", err)
		return
	}
	fmt.Println(git)

	DecryptedGit, err := pkg.DecryptJsonValue(git, key)
	if err != nil {
		fmt.Println("Json Decrypt Error:", err)
		return
	}
	JsonData, err := pkg.GitToJson(DecryptedGit)
	
	err = ioutil.WriteFile(JsonFile, JsonData, 0644)

	fmt.Println(DecryptedGit)

}
