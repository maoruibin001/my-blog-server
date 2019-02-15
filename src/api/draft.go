package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"my-blog-server/src/config"
	"my-blog-server/src/db"
	"my-blog-server/src/middleware"
	"my-blog-server/src/utils"
	"net/http"
	"strconv"
	"strings"
)

func createDraft(title, tags string, content string) interface{} {
	tagList := strings.Split(tags, ",")
	info := db.CreateArticle(title, tagList, 0, 0, content)
	return info
}
func initDraft(router *gin.Engine) {
	//创建草稿
	router.POST("/api/draft", middleware.JWTAuth(),func(context *gin.Context) {
		fmt.Println("create Draft ....")

		title := context.PostForm("title")
		tags := context.PostForm("tags")
		content := context.PostForm("content")


		draftInfo := createDraft(title, tags, content)

		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, draftInfo)
	})

	//查询草稿信息

	router.GET("/api/drafts/:aid", func(context *gin.Context) {
		aidStr := context.Param("aid")

		aid, err := strconv.Atoi(aidStr)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}
		fmt.Println("aid is: ", aid)

		draft := db.ArticleSingleFindByKV(bson.M{"aid": aid})

		fmt.Println("draft: ", draft)
		if draft.Aid == aid && aid != 0{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, draft)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
		}
	})
	//修改草稿
	router.PATCH("/api/drafts/:aid",middleware.JWTAuth(), func(context *gin.Context) {

		var article = db.ArticleSchema{}
		context.BindJSON(&article)

		//context.Request.ParseForm()
		title := article.Title
		content := article.Content
		tags := article.Tags

		fmt.Println("draft data: ", article)
		aidStr := context.Param("aid")
		aid, err := strconv.Atoi(aidStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}

		draft := db.ArticleSingleFindByKV(bson.M{"aid": aid})

		if draft.Aid != aid || aid == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		if title == "" {
			title = draft.Title
		}
		if content == "" {
			content = draft.Content
		}
		if len(tags) == 0 {
			tags = draft.Tags
		}

		err = db.ChangeArticle(aid, title, tags, false, draft.CommentN, content)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			draft := db.ArticleSingleFindByKV(bson.M{"aid": aid})
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, draft)
		}
	})
	//删除草稿
	router.DELETE("/api/drafts/:aid",middleware.JWTAuth(), func(context *gin.Context) {

		aidStr := context.Param("aid")

		aid, err := strconv.Atoi(aidStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}

		//查询用户
		draft := db.ArticleSingleFindByKV(bson.M{"aid": aid, "ispublish": false})

		if draft.Aid != aid || aid == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		err = db.RemoveArticle("aid", aid)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"message": "删除失败"	,
			})
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
		}
	})


	//分页查询所有草稿

	router.GET("/api/drafts", func(context *gin.Context) {
		pageSizeStr := context.DefaultQuery("pageSize", config.DEFAULTPAGESIZI)
		pageNoStr := context.DefaultQuery("pageNo", config.DEFAULTPAGENO)

		fmt.Println("pageSizeStr", pageNoStr)
		tag := context.Query("tag")

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			//pageSize = 10
			fmt.Println("err: ", err)
		}

		pageNo, err := strconv.Atoi(pageNoStr)
		if err != nil {
			//pageNo = 1
		}

		fmt.Println("pagesize: ", pageSize, pageNo)
		var drafts db.ResData

		fmt.Println("tags is: ", tag)
		if tag != "" && tag != "全部"{
			drafts, err = db.GetArticles(bson.M{"tags": tag, "ispublish": false}, pageSize, pageNo)
		} else {
			drafts, err = db.GetArticles(bson.M{"ispublish": false}, pageSize, pageNo)
		}
		//drafts, err := db.GetDrafts("", "", pageSize, pageNo, 0)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
			})
			return
		}
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, drafts)
	})

	router.GET("/api/someDrafts", func(context *gin.Context) {
		pageSizeStr := context.DefaultQuery("pageSize", config.DEFAULTPAGESIZI)
		pageNoStr := context.DefaultQuery("pageNo", config.DEFAULTPAGENO)

		//pageSizeStr := context.Query("pageSize")
		//pageNoStr := context.Query("pageNo")
		key := context.Query("key")

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			pageSize = 10
		}

		pageNo, err := strconv.Atoi(pageNoStr)
		if err != nil {
			pageNo = 1
		}

		fmt.Println("key: ", key)
		var drafts db.ResData

		if key == "" {
			drafts, err = db.GetArticles(nil, pageSize, pageNo)
		} else {
			conditions := []bson.M{
				bson.M{"title": bson.M{"$regex": key, "$options": "$i"}},
				bson.M{"content": bson.M{"$regex": key, "$options": "$i"}},
				bson.M{"datestr": bson.M{"$regex": key, "$options": "$i"}},
				bson.M{"tags": bson.M{"$regex": key, "$options": "$i"}},
				bson.M{"ispublish": false},
			}
			drafts, err = db.GetSomeArticles(bson.M{"$or": conditions}, pageSize, pageNo)
		}


		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
			})
			return
		}
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, drafts)
	})
}

