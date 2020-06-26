package entity

import (
	"time"
)

type Post struct {
	ID       uint64    `json:"id" binding:"required" bson:"_id"`
	Date     time.Time `json:"date"`
	Text     string    `json:"textContent" binding:"required"`
	ImageURL string    `json:"imageURL" bson:"imageURL"`
	UserID   uint64    `json:"-" bson:"userID"`
}

type PostWithEmail struct {
	//*Post
	Date     time.Time `json:"date"`
	Text     string    `json:"textContent" binding:"required"`
	ImageURL string    `json:"imageURL"`

	Email string `json:"email"`
}
