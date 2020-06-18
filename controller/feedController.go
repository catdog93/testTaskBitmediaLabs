package controller

import "github.com/gin-gonic/gin"

func GetFeedPage(context *gin.Context) {
	context.HTML(200, "blog.html", nil)
}
