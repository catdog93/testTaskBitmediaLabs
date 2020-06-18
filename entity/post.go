package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID           primitive.ObjectID
	Date         time.Time `json:"date"`
	TextContent  string    `json:"textContent"`
	MediaContent string    `json:"mediaContent"`
	Comments     *[]Comment
}
