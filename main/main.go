package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testTaskBitmediaLabs/controller"
	"time"
	//"testTaskBitmediaLabs/data"
	//rep "testTaskBitmediaLabs/repository"
)

// gin http router
var router *gin.Engine

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	context, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// Connect to MongoDB
	client, err := mongo.Connect(context, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context, nil)

	if err != nil {
		log.Fatal(err)
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
	err := router.Run() // listen and serve on localhost:8080
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
