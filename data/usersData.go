package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"testTaskBitmediaLabs/entity"
)

const basePath = "testTaskBitmediaLabs"
const targPath = "data/users.json"

type UsersUnmarshalled struct {
	Users []entity.UserBody `json:"objects"`
}

func ReadJSONData() []interface{} {
	usersUnmarshalled := UsersUnmarshalled{Users: []entity.UserBody{}}
	docsPath, err := filepath.Rel(basePath, targPath)
	if err != nil {
		log.Fatal(err)
	}
	byteValues, err := ioutil.ReadFile(docsPath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValues, &usersUnmarshalled)
	if err != nil {
		log.Fatal(err)
	}
	docs := make([]interface{}, len(usersUnmarshalled.Users))
	// replicating original slice to docs
	for index, _ := range usersUnmarshalled.Users {
		docs[index] = usersUnmarshalled.Users[index]
	}
	return docs
}
