package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initDraft(router *gin.Engine) {
	router.GET("/draft", func(context *gin.Context) {
		fmt.Println("articles ....")
		context.JSON(http.StatusOK, gin.H{"msg": "welcome to Draft"})
	})
}
