// Service is abstract layer between controller and repository. Preferably, it contains some logic here.
package service

import (
	ai "github.com/night-codes/mgo-ai"
	mgo "gopkg.in/mgo.v2"
	"testTaskBitmediaLabs/entity"
)

const (
	DBURL           = "mongodb://localhost:27017"
	DBName          = "Mblog"
	UsersCollection = "Users"
)

type Obj map[string]interface{}

type ID struct {
	ID string `json:"id" bson:"_id"`
}

// Function provides pagination of users data, requires limit number of users per page and number of page
//func GetUsersPagination(limit int64, pageNumber int64) (*[]entity.User, error) {
//	return rep.ReadUsersPagination(limit, pageNumber)
//}
//
//func GetUserByID(id string) (entity.User, error) {
//	return rep.ReadUserByID(id)
//}

func CreateUser(user entity.User) error {
	session, err := mgo.Dial(DBURL)
	if err != nil {
		return err
	}
	collection := session.DB(DBName).C(UsersCollection)

	ai.Connect(session.DB(DBName).C(UsersCollection))
	user.ID = ai.Next(UsersCollection)
	return collection.Insert(interface{}(user))
}

func FindUser(user entity.UserBody) (*entity.User, error) {
	session, err := mgo.Dial(DBURL)
	if err != nil {
		return nil, err
	}
	collection := session.DB(DBName).C(UsersCollection)
	//query := collection.Find(user)
	//number, err := query.Count()
	//if err != nil {
	//	return err
	//}
	//if number != 1 {
	//	return errors.New("Can't find user in DB")
	//}
	result := entity.User{}
	err = collection.Find(Obj{"email": user.Email, "password": user.Password}).One(&result)
	return &result, err
}

func FindUserByEmail(email string) (*entity.User, error) {
	session, err := mgo.Dial(DBURL)
	if err != nil {
		return nil, err
	}
	collection := session.DB(DBName).C(UsersCollection)
	result := entity.User{}
	err = collection.Find(Obj{"email": email}).One(&result)
	return &result, err
}
