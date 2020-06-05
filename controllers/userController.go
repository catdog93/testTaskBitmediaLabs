package controllers

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"testTaskBitmediaLabs/entities"
	"testTaskBitmediaLabs/services"
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
	user, err := services.GetUserById(id)
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
	users, err := services.GetUsersLimit(uint64(limit))
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusNotFound, err.Error())
		return
	}
	context.JSON(http.StatusOK, users)
}

//users/
func CreateUser(context *gin.Context) {
	user := entities.UserBody{}
	err := context.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := services.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, id)
}

//users/
func ReplaceUser(context *gin.Context) {
	user := entities.User{}
	err := context.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	err = services.ReplaceUser(&user)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
}
