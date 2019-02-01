package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initArticle(router *gin.Engine) {
	router.GET("/article", func(context *gin.Context) {
		fmt.Println("articles ....")
		context.JSON(http.StatusOK, gin.H{"msg": "welcome to Article"})
	})
}
