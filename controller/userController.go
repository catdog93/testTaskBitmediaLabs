package controller

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"testTaskBitmediaLabs/entity"
	"testTaskBitmediaLabs/service"
)

const (
	RelativeUsersPath string = "/users"
	UsersPath         string = "/"
	GetUserPath       string = "/:id"
)

const idLength = 24

// users/5eda0e63a84a6e050000d115
func GetUser(context *gin.Context) {
	stringURL := context.Request.URL.String()
	id := path.Base(stringURL)
	if len(id) != idLength {
		context.String(http.StatusBadRequest, "length of id must be = %v", idLength)
		return
	}
	user, err := service.GetUserById(id)
	if err != nil {
		context.String(http.StatusNotFound, err.Error())
		return
	}
	context.JSON(http.StatusOK, user)
}

//users/
func GetUsers(context *gin.Context) {
	limit, err := strconv.Atoi(context.Query("limit"))
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	users, err := service.GetUsersLimit(uint64(limit))
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusNotFound, err.Error())
		return
	}
	context.JSON(http.StatusOK, users)
}

//users/
func CreateUser(context *gin.Context) {
	user := entity.UserBody{}
	err := context.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := service.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, id)
}

//users/
func ReplaceUser(context *gin.Context) {
	user := entity.User{}
	err := context.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	err = service.ReplaceUser(&user)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
}
