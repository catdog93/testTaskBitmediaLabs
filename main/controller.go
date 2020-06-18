package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TokenAuth(context *gin.Context) {
	token, err := context.Cookie("token")
	if err == nil {
		_, ok := Tokens[token]
		if ok {
			//context.JSON(http.StatusOK, user)
			return
		}
	}
	context.Redirect(http.StatusMovedPermanently, "/authorization")
	context.Abort()
}

func InsertSomeValues() {
	Tokens["1"] = UserAuth{Login: "lala23", Pass: "4444"}
	Tokens["2"] = UserAuth{Login: "oleg", Pass: "oleg"}
}

func Registration(c *gin.Context) {
	login := c.PostForm("login")
	pass := c.PostForm("pass")

	user := UserAuth{
		Login: login,
		Pass:  pass,
	}

	UsersLock.Lock()
	Users[user] = struct{}{}
	UsersLock.Unlock()

	c.String(http.StatusOK, "Registration OK")
}

func AuthorizationPost(c *gin.Context) {
	login := c.PostForm("login")
	pass := c.PostForm("pass")
	fmt.Println(login, pass)
	user := UserAuth{
		Login: login,
		Pass:  pass,
	}
	token := createToken(user)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   60,
		Path:     "/",
		Domain:   "localhost",
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		HttpOnly: true,
	})
	c.String(http.StatusOK, token)
}

func Authorization(c *gin.Context) {
	c.HTML(http.StatusOK, "authorization.html", nil)
}

func Default(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", nil)
}

func SecretHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "secretDate.html", time.Now())
}
