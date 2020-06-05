package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"testTaskBitmediaLabs/controllers"
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
	routerGroup := router.Group(controllers.RelativeUsersPath)
	{
		routerGroup.POST(controllers.UsersPath, controllers.CreateUser)
		routerGroup.GET(controllers.UsersPath, controllers.GetUsers)
		routerGroup.GET(controllers.GetUserPath, controllers.GetUser)
		routerGroup.PUT(controllers.UsersPath, controllers.ReplaceUser)
	}
}
