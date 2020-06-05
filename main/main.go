package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"testTaskBitmediaLabs/controller"
	//"testTaskBitmediaLabs/data"
	//rep "testTaskBitmediaLabs/repository"
)

var router *gin.Engine

func main() {
	//err := rep.InsertUsers(data.ReadJSONData())
	//if err != nil {
	//	log.Fatal(err)
	//}
	CreateUrlMapping()
	err := router.Run() // listen and serve on localhost:8080
	if err != nil {
		log.Fatal(err)
	}
}

// map handlers with endpoints
func CreateUrlMapping() {
	router = gin.Default()
	routerGroup := router.Group(controller.RelativeUsersPath)
	{
		routerGroup.POST(controller.UsersPath, controller.CreateUser)
		routerGroup.GET(controller.UsersPath, controller.GetUsers)
		routerGroup.GET(controller.GetUserPath, controller.GetUser)
		routerGroup.PUT(controller.UsersPath, controller.ReplaceUser)
	}
}
