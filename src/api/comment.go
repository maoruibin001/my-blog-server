package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-blog-server/src/db"
	"my-blog-server/src/utils"
	"net/http"
	"strconv"
)


func initComment(router *gin.Engine) {
	// 获取某一篇文章的所有评论
	router.GET("/api/comments", func(context *gin.Context) {
		aidStr := context.Query("aid")

		fmt.Println("aidStr: ", aidStr)
		aid, err := strconv.Atoi(aidStr)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}
		fmt.Println("aid is: ", aid)
		comments, err := db.QueryCommentSet(aid)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"errorMsg": "服务端出错",
			})
			return

		}
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, comments)
	})

	// 发布评论并通知站长和评论者
	router.POST("/api/comment", func(context *gin.Context) {

		aidStr := context.PostForm("aid")

		aid, err := Str2Int(aidStr, context)

		if err != nil {
			return
		}

		imgName := context.PostForm("imgName")
		name := context.PostForm("name")
		address := context.PostForm("address")
		content := context.PostForm("content")


		comment, err := db.AddComment(aid, imgName, name, address, content)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"errorMsg": "服务端出错",
			})
			return
		}

		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, comment)
	})

	// 更新评论的点赞数
	router.PATCH("/api/comments", func(context *gin.Context) {

		context.Request.ParseForm()
		cidStr := context.Request.FormValue("cid")
		aidStr := context.Request.FormValue("aid")

		aid, err := Str2Int(aidStr, context)
		cid, cidErr := Str2Int(cidStr, context)

		fmt.Println(aid, cid)

		imgName := context.Request.FormValue("imgName")
		name := context.Request.FormValue("name")
		address := context.Request.FormValue("address")
		content := context.Request.FormValue("content")

		if err != nil || cidErr != nil {
			return
		}

		commentSet, index, err := db.QueryCommentIndex(aid, cid)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"errorMsg": err.Error(),
			})
			return
		}

		comment := commentSet.Comments[index]
		if comment.Aid != aid || comment.Cid != cid || aid == 0 || cid == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		if imgName == "" {
			imgName = comment.ImgName
		}
		if name == "" {
			name = comment.Name
		}
		if content == "" {
			content = comment.Content
		}
		if address == "" {
			address = comment.Address
		}

		comment, err = db.ChangeComment(aid, cid, imgName, name, address, content)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, comment)
		}
	})

	router.DELETE("/api/comments/:aid/:cid", func(context *gin.Context) {
		aidStr := context.Param("aid")
		cidStr := context.Param("cid")

		aid, err := Str2Int(aidStr, context)
		cid, cidErr := Str2Int(cidStr, context)

		if err != nil || cidErr != nil {
			return
		}

		err = db.DeleteComment(aid, cid)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"message": "删除失败"	,
			})
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
		}

	})
}
