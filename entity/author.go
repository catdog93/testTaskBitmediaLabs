package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id"`
	Nickname  string             `json:"nickname" binding:"required" bson:"nickname"`
	AvatarImg string             `json:"avatarImg" bson:"avatarImg"`
	*Feed     `json:"feed"`
	Posts     *[]Post
}

type AuthorBody struct {
	Nickname string `json:"nickname" binding:"required" bson:"nickname"`
}
