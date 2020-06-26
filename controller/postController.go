package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testTaskBitmediaLabs/entity"
	"testTaskBitmediaLabs/service"
	"time"
)

const (
	text        = "inputText"
	image       = "inputImage"
	tokenString = "token"
)

func CreatePost(context *gin.Context) {
	text := context.PostForm(text)
	imageURL := context.PostForm(image)

	token, err := context.Cookie(tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	userAuth, _ := service.Tokens[token]
	user, err := service.FindUserByEmail(userAuth.Email)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	post := entity.Post{
		Text:     text,
		ImageURL: imageURL,
		Date:     time.Now(),
		UserID:   user.ID,
	}
	err = service.CreatePost(post)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
	}
}

func CreatePostForm(context *gin.Context) {
	token, err := context.Cookie(tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, _ := service.Tokens[token]
	context.HTML(http.StatusOK, "createPostForm.html", user.Email)
}
