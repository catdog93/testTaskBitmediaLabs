package service

import (
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"testTaskBitmediaLabs/entity"
)

const (
	PostsCollection = "Posts"
)

func CreatePost(post entity.Post) error {
	session, err := mgo.Dial(DBURL)
	if err != nil {
		return err
	}
	collection := session.DB(DBName).C(PostsCollection)

	ai.Connect(session.DB(DBName).C(PostsCollection))
	post.ID = ai.Next(PostsCollection)
	return collection.Insert(interface{}(post))
}

func GetPostWithOwner(post *entity.Post) (*entity.PostWithEmail, error) {
	session, err := mgo.Dial(DBURL)
	if err != nil {
		return nil, err
	}
	collection := session.DB(DBName).C(PostsCollection)
	result := entity.User{}
	err = collection.FindId(post.UserID).One(&result)
	if err != nil {
		return nil, err
	}
	postWithEmail := &entity.PostWithEmail{
		Date:     post.Date,
		Text:     post.Text,
		ImageURL: post.ImageURL,
		Email:    result.Email,
	}
	return postWithEmail, nil
}

func GetPost() (*entity.Post, error) {
	session, err := mgo.Dial(DBURL)
	if err != nil {
		return nil, err
	}
	collection := session.DB(DBName).C(PostsCollection)
	post := entity.Post{}
	err = collection.Find(Obj{}).One(&post)
	return &post, err
}
