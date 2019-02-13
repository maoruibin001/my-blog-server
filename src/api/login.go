package api

import (
	"github.com/gin-gonic/gin"
	"my-blog-server/src/db"
	"my-blog-server/src/utils"
	"net/http"
)

type JSONFormat struct {
	code int
	msg string
	data map[string]interface{}

}

func initLogin(router *gin.Engine) {
	router.GET("/api/login", func(context *gin.Context) {
		phone := context.PostForm("phone")
		result := db.FindByPhone(phone)
		if result.Phone != phone || phone == "" {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
			return
		}

		context.JSON(http.StatusOK, "hello")


	})
}
