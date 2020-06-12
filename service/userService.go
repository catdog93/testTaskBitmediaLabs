// Service is abstract layer between controller and repository. Preferably, it contains some logic here.
package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testTaskBitmediaLabs/entity"
	rep "testTaskBitmediaLabs/repository"
)

// Function provides pagination of users data, requires limit number of users per page and number of page
func GetUsersPagination(limit int64, pageNumber int64) (*[]entity.User, error) {
	return rep.ReadUsersPagination(limit, pageNumber)
}

func GetUserByID(id string) (entity.User, error) {
	return rep.ReadUserByID(id)
}

// If it's is successful CreateUser() returns ID of created user
func CreateUser(user entity.UserBody) (*primitive.ObjectID, error) {
	var id primitive.ObjectID

	objectID, err := rep.CreateUser(user)
	if err != nil {
		return nil, err
	}
	switch objectID.(type) {
	case primitive.ObjectID:
		id = objectID.(primitive.ObjectID)
	default:
		err = errors.New("user wasn't created")
	}
	return &id, err
}

// One of Update's variant: to replace User using ID
func UpdateUser(user entity.User) error {
	return rep.ReplaceUserByID(user.ID, interface{}(user))
}
