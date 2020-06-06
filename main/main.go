package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testTaskBitmediaLabs/controller"
	"testTaskBitmediaLabs/repository"
	//"testTaskBitmediaLabs/data"
	//rep "testTaskBitmediaLabs/repository"
)

const DBUri = "mongodb://localhost:27017"

// gin http router
var router *gin.Engine

func main() {
	ctx := repository.GetClient()
	err := ctx.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	// import users data to MongoDB
	//docs, err := data.ReadJSONData(data.BasePath, data.TargPath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = rep.InsertUsers(docs)
	//if err != nil {
	//	log.Fatal(err)
	//}

	CreateUrlMapping()
	err = router.Run() // listen and serve on localhost:8080
	if err != nil {
		log.Fatal(err)
	}
}

// handlers are mapped with endpoints
func CreateUrlMapping() {
	router = gin.Default()
	routerGroup := router.Group(controller.RelativeUsersPath)
	{
		routerGroup.POST(controller.UsersPath, controller.CreateUser)
		routerGroup.GET(controller.UsersPath, controller.GetUsers)
		routerGroup.GET(controller.GetUserPath, controller.GetUserByID)
		routerGroup.PUT(controller.UsersPath, controller.ReplaceUser)
	}
}
