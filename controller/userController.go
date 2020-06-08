// Controller has required API mapped with endpoints below.
package controller

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"testTaskBitmediaLabs/entity"
	"testTaskBitmediaLabs/service"
	"testTaskBitmediaLabs/validator"
)

const (
	RelativeUsersPath string = "/users"
	UsersPath         string = "/"
	GetUserByIDPath   string = "/:id"
)

const (
	objectIDLength = 24
	lowestLimit    = 1
	highestLimit   = 10000

	idLengthError   = "error: length of id must be = %v"
	pageNumberError = "error: incorrect page number"
	limitError      = "error: incorrect limit value"
)

// GetUsers() provides pagination data, requires limit number of users per page and number of page
// Endpoint example: users/?limit=100&page=5
func GetUsers(context *gin.Context) {
	limit, err := strconv.Atoi(context.Query("limit"))
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	if limit < lowestLimit || limit > highestLimit { // set default limit value or return 400 BadRequest
		context.String(http.StatusBadRequest, limitError)
		return
	}

	pageNumber, err := strconv.Atoi(context.Query("page"))
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	if pageNumber < lowestLimit {
		context.String(http.StatusBadRequest, pageNumberError)
		return
	}
	users, err := service.GetUsersPagination(int64(limit), int64(pageNumber))
	if err != nil {
		context.String(http.StatusNotFound, err.Error())
		return
	}
	context.JSON(http.StatusOK, users)
}

// Endpoint example: users/5eda0e63a84a6e050000d115
func GetUserByID(context *gin.Context) {
	stringURL := context.Request.URL.String()
	id := path.Base(stringURL)
	if len(id) != objectIDLength {
		context.String(http.StatusBadRequest, idLengthError, objectIDLength)
		return
	}
	user, err := service.GetUserByID(id)
	if err != nil {
		context.String(http.StatusNotFound, err.Error())
		return
	}
	context.JSON(http.StatusOK, user)
}

// CreateUser() provides validation of request's UserBody. If creation is successful it returns ID in response's body
// Endpoint example: users/
func CreateUser(context *gin.Context) {
	user := entity.UserBody{}
	err := context.BindJSON(&user)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	err = validator.UserValidation(user)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := service.CreateUser(user)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, id)
}

// CreateUser() provides validation of request's UserBody. If updating is successful it returns ID in response's body
// Endpoint example: users/
func UpdateUser(context *gin.Context) {
	user := entity.User{}
	err := context.BindJSON(&user)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	err = validator.UserValidation(user.ConvertUserToUserBody())
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	err = service.UpdateUser(user)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
}
