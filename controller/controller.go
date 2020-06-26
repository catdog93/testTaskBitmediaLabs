package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testTaskBitmediaLabs/service"
)

func Default(c *gin.Context) {
	post, err := service.GetPost()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	postWithEmail, err := service.GetPostWithOwner(post)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	obj := service.Obj{
		"date":  postWithEmail.Date,
		"text":  postWithEmail.Text,
		"email": postWithEmail.Email,
		"image": postWithEmail.ImageURL,
	}
	c.HTML(http.StatusOK, "index.html", obj)
}
