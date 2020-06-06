package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id"`
	LastName  string             `json:"last_name" binding:"required" bson:"lastName"`
	Gender    Gender             `json:"gender,omitempty" bson:"gender"`
	Country   string             `json:"country,omitempty" bson:"country"`
	City      string             `json:"city,omitempty" bson:"city"`
	Email     string             `json:"email" binding:"required" bson:"email"`
	BirthDate string             `json:"birth_date,omitempty" bson:"birthDate"`
}

type UserBody struct {
	LastName  string `json:"last_name" binding:"required" bson:"lastName"`
	Gender    Gender `json:"gender,omitempty" bson:"gender"`
	Country   string `json:"country,omitempty" bson:"country"`
	City      string `json:"city,omitempty" bson:"city"`
	Email     string `json:"email" binding:"required" bson:"email"`
	BirthDate string `json:"birth_date,omitempty" bson:"birthDate"`
}

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"

	MaleLower   Gender = "male"
	FemaleLower Gender = "female"
)

func (user User) ConvertUserToUserBody() UserBody {
	userBody := UserBody{
		LastName:  user.LastName,
		Gender:    user.Gender,
		Country:   user.Country,
		City:      user.City,
		Email:     user.Email,
		BirthDate: user.BirthDate,
	}
	return userBody
}
