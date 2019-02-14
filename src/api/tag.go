package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-blog-server/src/db"
	"my-blog-server/src/utils"
	"net/http"
)

func initTag(router *gin.Engine) {
	router.GET("/api/tags", func(context *gin.Context) {
		fmt.Println("articles ....")
		tags, err := db.GetTags()

		fmt.Println("tags is: ", tags)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, nil)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, tags)
		}
	})
}
