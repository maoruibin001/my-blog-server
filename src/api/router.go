package api

import (
	"github.com/gin-gonic/gin"
	"my-blog-server/src/db"
	"my-blog-server/src/utils"
	"net/http"
	"strconv"
)

func InitRouter(router *gin.Engine) {

	initImage(router)
	initComment(router)
	initDraft(router)
	initTag(router)
	initUser(router)
	initLogin(router)
	initArticle(router)


	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	db.GetUserInfo(router)
}

func Str2Int(str string, context *gin.Context) (int, error)  {
	intRet, err := strconv.Atoi(str)
	if err != nil {
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
			"errorMsg": "参数错误，请输入正确的参数: " + str,
		})
		return utils.RESPONSEPARAMERROR, err

	}
	return intRet, nil

}