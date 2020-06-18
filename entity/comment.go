package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID     primitive.ObjectID
	Author *Author `json:"author"`
	Text   string  `json:"text"`
}
