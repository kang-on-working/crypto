package pkg

import (
	"encoding/json"
	"io/ioutil"
)

// GIT 구조체는 git.json 파일의 내용을 담는 구조체입니다.
type GIT struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Repo     string `json:"repo"`
}

// JsonParse 함수는 주어진 JSON 파일을 읽어들여 GIT 구조체로 반환합니다.
func JsonParse(jsonFile string) (git GIT, err error) {
	fileData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return GIT{}, err
	}

	err = json.Unmarshal(fileData, &git)
	if err != nil {
		return GIT{}, err
	}

	return git, nil
}
