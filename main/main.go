package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"sync"

	//"github.com/go-redis/redis/v7"
	"log"
	"time"
)

//var  client *redis.Client
//func init() {
//	//Initializing redis
//	dsn := os.Getenv("REDIS_DSN")
//	if len(dsn) == 0 {
//		dsn = "localhost:6379"
//	}
//	client = redis.NewClient(&redis.Options{
//		Addr: dsn, //redis port
//	})
//	_, err := client.Ping().Result()
//	if err != nil {
//		panic(err)
//	}
//}

var Router = gin.Default()

var Tokens = map[string]UserAuth{}
var Users = map[UserAuth]struct{}{}

var UsersLock sync.Mutex
var TokensLock sync.Mutex

const SECRET_WORD = "letsgo"

func main() {
	Router.GET("/", Default)

	group := Router.Group("/auth")
	group.Use(TokenAuth)
	group.GET("/secret", SecretHandler)

	Router.POST("/registration", Registration)
	Router.GET("/authorization", Authorization)
	Router.POST("/authorization", AuthorizationPost)

	Router.LoadHTMLGlob("../templates/*.html")

	InsertSomeValues()

	Router.POST("/token/", TokenAuth)
	err := Router.Run() // listen and serve on localhost:8080
	if err != nil {
		log.Fatal(err)
	}
}

func createToken(authData UserAuth) string {
	jwtToken, err := generateToken()
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	Tokens[jwtToken] = authData

	return jwtToken
}

func generateToken() (token string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err = at.SignedString([]byte(SECRET_WORD))
	if err != nil {
		return "", err
	}

	return token, nil
}

// handlers are mapped with endpoints
//func CreateUrlMapping() {
//	router = gin.Default()
//	router.LoadHTMLGlob("../templates/*.html")
//	router.Static("files", "../templates")
//	routerGroup := router.Group(controller.RelativeUsersPath)
//	{
//		routerGroup.POST(controller.UsersPath, controller.CreateUser)
//		routerGroup.GET(controller.UsersPath, controller.GetUsers)
//		routerGroup.GET(controller.GetUserByIDPath, controller.GetUserByID)
//		//routerGroup.PUT(controller.UsersPath, controller.UpdateUser)
//	}
//	//routerGroup.Static("files", "./templates")
//	router.GET("index", controller.GetMainPage)
//	router.GET("feed", controller.GetFeedPage)
//	router.GET("lul", controller.GetUsers)
//}
