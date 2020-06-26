package main

import (
	"github.com/gin-gonic/gin"
	"testTaskBitmediaLabs/controller"
	//"github.com/go-redis/redis/v7"
	"log"
)

var Router = gin.Default()

const (
	signin = "/signin"
	signup = "/signup"

	mblog      = "/mblog"
	createPost = "/createPost"
	users      = "/users"
)

func main() {
	group := Router.Group(mblog)
	group.Use(controller.TokenAuth)
	group.GET("/index", controller.Default)

	group.GET(createPost, controller.CreatePostForm)
	group.POST(createPost, controller.CreatePost)

	Router.POST(signup, controller.SignupPost)
	Router.GET(signup, controller.Signup)
	Router.GET(signin, controller.Signin)
	Router.POST(signin, controller.SigninPost)

	Router.LoadHTMLGlob("../templates/*.html")
	Router.Static("files", "../templates")

	err := Router.Run() // listen and serve on localhost:8080
	if err != nil {
		log.Fatal(err)
	}
}
