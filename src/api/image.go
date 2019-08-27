package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"my-blog-server/src/db"
	"my-blog-server/src/utils"
	)

func initImage(router *gin.Engine) {
	router.POST("/api/upload", func(context *gin.Context) {
		response,err := http.Get("http://www.baidu.com")
		if(err!=nil){
			fmt.Println(err)
		}
		defer response.Body.Close()
		body,err := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))

		tags, err := db.GetTags()

		fmt.Println("tags is: ", tags)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, nil)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, tags)
		}
	})
}
