package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testTaskBitmediaLabs/entity"
	rep "testTaskBitmediaLabs/repository"
)

const DBUri = "mongodb://localhost:2717"

func GetUsersLimit(limit int64, pageNumber int64) (*[]entity.User, error) {
	return rep.ReadUsersPagination(limit, pageNumber)
}

func GetUserByID(id string) (*entity.User, error) {
	return rep.ReadUserByID(id)
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
