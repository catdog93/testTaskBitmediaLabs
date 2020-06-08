// data implements reading Users data from JSON file.
package data

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testTaskBitmediaLabs/entity"
)

const BasePath = "testTaskBitmediaLabs"
const TargPath = "data/users.json"

type UsersUnmarshalled struct {
	Users []entity.UserBody `json:"objects"`
}

// Function reads data that have JSON format from file
func ReadJSONData(basePath, targPath string) ([]interface{}, error) {
	usersUnmarshalled := UsersUnmarshalled{Users: []entity.UserBody{}}
	docsPath, err := filepath.Rel(basePath, targPath)
	if err != nil {
		return nil, err
	}
	byteValues, err := ioutil.ReadFile(docsPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValues, &usersUnmarshalled)
	if err != nil {
		return nil, err
	}
	docs := make([]interface{}, len(usersUnmarshalled.Users))
	// replicating original slice to docs
	for index, _ := range usersUnmarshalled.Users {
		docs[index] = usersUnmarshalled.Users[index]
	}
	return docs, nil
}
