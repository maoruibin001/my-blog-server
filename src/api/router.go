package api

import (
	"github.com/gin-gonic/gin"
	"album-server/src/db"
	"album-server/src/utils"
	"net/http"
	"strconv"
)

func InitRouter(router *gin.Engine) {

	initBseries(router)
	initLseries(router)
	initImage(router)
	initProduct(router)
	initUser(router)
	initLogin(router)
	initAccount(router)
	initCollection(router)


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