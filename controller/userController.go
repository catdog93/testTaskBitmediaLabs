package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testTaskBitmediaLabs/entity"
	"testTaskBitmediaLabs/service"
)

const (
	signin        = "/signin"
	signupSucceed = "Signup succeed :)"

	incorrectEmailOrPassword = "Incorrect email or password"
	suchEmailAlreadyExists   = "Such email already exists"

	email    = "inputEmail"
	password = "inputPassword"
)

func TokenAuth(context *gin.Context) {
	token, err := context.Cookie(tokenString)
	if err == nil {
		_, ok := service.Tokens[token]
		if ok {
			return
		}
	}
	context.Redirect(http.StatusMovedPermanently, signin)
	context.Abort()
}

func SignupPost(c *gin.Context) {
	email := c.PostForm(email)
	pass := c.PostForm(password)

	user := entity.User{
		Email:    email,
		Password: pass,
	}

	err := service.CreateUser(user)
	if err != nil {
		c.HTML(http.StatusOK, "signup.html", suchEmailAlreadyExists)
		return
	}
	c.String(http.StatusOK, signupSucceed)
}

func SigninPost(c *gin.Context) {
	email := c.PostForm(email)
	password := c.PostForm(password)

	userBody := entity.UserBody{
		Email:    email,
		Password: password,
	}
	user, err := service.FindUser(userBody)
	if err != nil && user != nil {
		c.HTML(http.StatusOK, "signin.html", incorrectEmailOrPassword)
		return
	}
	token := service.CreateToken(user)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     tokenString,
		Value:    token,
		MaxAge:   600,
		Path:     "/",
		Domain:   "localhost",
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		HttpOnly: true,
	})
	c.String(http.StatusOK, token)
}

func Signin(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

//const (
//	RelativeUsersPath string = "/users"
//	UsersPath         string = "/"
//	GetUserByIDPath   string = "/:id"
//)
//
//const (
//	objectIDLength = 24
//	lowestLimit    = 1
//	highestLimit   = 10000
//
//	idLengthError   = "error: length of id must be = %v"
//	pageNumberError = "error: incorrect page number"
//	limitError      = "error: incorrect limit value"
//)
//
//// GetUsers() provides pagination data, requires limit number of users per page and number of page
//// Endpoint example: users/?limit=100&page=5
//func GetUsers(context *gin.Context) {
//	//limit, err := strconv.Atoi(context.Query("limit"))
//	//if err != nil {
//	//	context.String(http.StatusBadRequest, err.Error())
//	//	return
//	//}
//	//if limit < lowestLimit || limit > highestLimit { // set default limit value or return 400 BadRequest
//	//	context.String(http.StatusBadRequest, limitError)
//	//	return
//	//}
//	//
//	//pageNumber, err := strconv.Atoi(context.Query("page"))
//	//if err != nil {
//	//	context.String(http.StatusBadRequest, err.Error())
//	//	return
//	//}
//	//if pageNumber < lowestLimit {
//	//	context.String(http.StatusBadRequest, pageNumberError)
//	//	return
//	//}
//	//users, err := service.GetUsersPagination(int64(limit), int64(pageNumber))
//	//if err != nil {
//	//	context.String(http.StatusNotFound, err.Error())
//	//	return
//	//}
//	context.JSON(http.StatusOK, cache.Cache)
//}
//
//// Endpoint example: users/5eda0e63a84a6e050000d115
//func GetUserByID(context *gin.Context) {
//	stringURL := context.Request.URL.String()
//	id := path.Base(stringURL)
//	if len(id) != objectIDLength {
//		context.String(http.StatusBadRequest, idLengthError, objectIDLength)
//		return
//	}
//	user, err := service.GetUserByID(id)
//	if err != nil {
//		context.String(http.StatusNotFound, err.Error())
//		return
//	}
//	context.JSON(http.StatusOK, user)
//}
//
//// CreateUser() provides validation of request's UserBody. If creation is successful it returns ID in response's body
//// Endpoint example: users/
//func CreateUser(context *gin.Context) {
//	user := entity.UserBody{}
//	err := context.BindJSON(&user)
//	if err != nil {
//		context.String(http.StatusBadRequest, err.Error())
//		return
//	}
//	err = validator.UserValidation(user)
//	if err != nil {
//		context.String(http.StatusBadRequest, err.Error())
//		return
//	}
//	id, err := service.CreateUser(user)
//	if err != nil {
//		context.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//	context.JSON(http.StatusCreated, id)
//}
//
//// CreateUser() provides validation of request's UserBody. If updating is successful it returns ID in response's body
//// Endpoint example: users/
////func UpdateUser(context *gin.Context) {
////	user := entity.User{}
////	err := context.BindJSON(&user)
////	if err != nil {
////		context.String(http.StatusBadRequest, err.Error())
////		return
////	}
////	err = validator.UserValidation(user.ConvertUserToUserBody())
////	if err != nil {
////		context.String(http.StatusBadRequest, err.Error())
////		return
////	}
////	err = service.UpdateUser(user)
////	if err != nil {
////		context.String(http.StatusInternalServerError, err.Error())
////		return
////	}
////}
