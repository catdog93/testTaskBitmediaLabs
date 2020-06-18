package controller

import "github.com/gin-gonic/gin"

func GetMainPage(context *gin.Context) {
	context.HTML(200, "index.html", nil)
}
