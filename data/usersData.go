// data implements reading Users data from JSON file.
package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func ReadJSONData(basePath, targPath string) ([]entity.User, error) {
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
	users := make([]entity.User, 0, len(usersUnmarshalled.Users))
	for index, value := range usersUnmarshalled.Users {
		id := primitive.NewObjectID()
		users[index].Author = new(entity.Author)
		users[index].Author.ID = id
		user := value.ConvertUserBodyToUser()
		user.ID = id
		users = append(users, user)
		users[index].ID = id
	}
	return users, err
}
