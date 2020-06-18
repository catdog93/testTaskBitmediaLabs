package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id"`
	*Author   `json:"author" binding:"required" bson:"author"`
	Gender    Gender `json:"gender,omitempty" bson:"gender"`
	Country   string `json:"country,omitempty" bson:"country"`
	Email     string `json:"email" binding:"required" bson:"email"`
	BirthDate string `json:"birth_date,omitempty" bson:"birthDate"`
}

type UserBody struct {
	*AuthorBody `json:"author" binding:"required" bson:"author"`
	Gender      Gender `json:"gender,omitempty" bson:"gender"`
	Country     string `json:"country,omitempty" bson:"country"`
	Email       string `json:"email" binding:"required" bson:"email"`
	BirthDate   string `json:"birth_date,omitempty" bson:"birthDate"`
}

type Gender string

const (
	// Preferable values for field Gender
	Male   Gender = "Male"
	Female Gender = "Female"

	MaleLower   Gender = "male"
	FemaleLower Gender = "female"
)

func (userBody UserBody) ConvertUserBodyToUser() User {
	user := User{
		Gender:    userBody.Gender,
		Country:   userBody.Country,
		Email:     userBody.Email,
		BirthDate: userBody.BirthDate,
	}
	user.Author.Nickname = userBody.AuthorBody.Nickname
	return user
}
