package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testTaskBitmediaLabs/entity"
	rep "testTaskBitmediaLabs/repository"
)

const DBUri = "mongodb://localhost:2717"

func GetUsersLimit(limit uint64) ([]entity.User, error) {
	return rep.ReadUsers(limit)
}

func GetUserById(id string) (*entity.User, error) {
	return rep.ReadUser(id)
}

func CreateUser(user *entity.UserBody) (*primitive.ObjectID, error) {
	var id primitive.ObjectID
	objectID, err := rep.CreateUser(interface{}(user))
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

func ReplaceUser(user *entity.User) error {
	err := rep.ReplaceUser(user.ID, interface{}(user))
	if err != nil {
		return err
	}
	return err
}
