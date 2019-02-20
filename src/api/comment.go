package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-blog-server/src/db"
	"my-blog-server/src/utils"
	"net/http"
	"strconv"
)

type CommentRequest struct {
	ImgName string `json:"imgName"`
	Name string `json:"name"`
	Address string `json:"address"`
	Content string `json:"content"`
	Aid int `json:"aid"`
	Cid int `json:"cid"`
	Date int64 `json:"date"`
	Like int `json:"like"`
	LikeEmails []string `json:"likeEmails"`
	AddLike int `json:"addLike"`
}


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

		var comment = db.CommentSchema{}
		context.BindJSON(&comment)

		comment, err := db.AddComment(comment.Aid, comment.ImgName, comment.Name, comment.Address, comment.Content)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"errorMsg": "服务端出错",
			})
			return
		}

		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, comment)
	})

	// 更新评论的点赞数
	router.PUT("/api/comments/:aid/:cid", func(context *gin.Context) {

		aidStr := context.Param("aid")
		cidStr := context.Param("cid")
		fmt.Println("aid is: ", aidStr, cidStr);
		context.Request.ParseForm()
		aid, err := Str2Int(aidStr, context)
		cid, cidErr := Str2Int(cidStr, context)

		params := CommentRequest{}
		context.BindJSON(&params)

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


		if params.ImgName== "" {
			params.ImgName = comment.ImgName
		}
		if params.Name == "" {
			params.Name = comment.Name
		}
		if params.Content == "" {
			params.Content = comment.Content
		}
		if params.Address == "" {
			params.Address = comment.Address
		}

		if params.AddLike == 1 {
			params.Like = comment.Like + 1
			params.LikeEmails = append(comment.LikeEmails, params.Address)
		} else {
			params.Like = comment.Like - 1
			likeEmails := utils.RemoveStringElement(comment.LikeEmails, func(e string, i int) bool {
				if e == comment.Address {
					return true
				}
				return false
			})
			params.LikeEmails = likeEmails
		}

		fmt.Println(aid, cid, params.ImgName, params.Name, params.Address, params.Content, params.Like)
		comment, err = db.ChangeComment(aid, cid, params.ImgName, params.Name, params.Address, params.Content, params.Like, params.LikeEmails)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, comment)
		}
	})

	router.DELETE("/api/comments/:aid/:cid", func(context *gin.Context) {
		aidStr := context.Param("aid")
		cidStr := context.Param("cid")

		//var params = CommentRequest{}
		//
		//context.BindJSON(&params)
		//
		//fmt.Println("params is: ", params)
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
