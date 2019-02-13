package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initTag(router *gin.Engine) {
	router.GET("/tag", func(context *gin.Context) {
		fmt.Println("articles ....")
		context.JSON(http.StatusOK, gin.H{"msg": "welcome to Tag"})
	})
}
