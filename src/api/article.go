package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"my-blog-server/src/db"
	"my-blog-server/src/utils"
	"net/http"
	"strconv"
	"strings"
)

func createArticle(title, tags string, content string) interface{} {
	tagList := strings.Split(tags, ",")
	info := db.CreateArticle(title, tagList, 0, 0, content)
	return info
}
func initArticle(router *gin.Engine) {
	//创建文章
	router.POST("/api/article", func(context *gin.Context) {
		fmt.Println("create Article ....")

		title := context.PostForm("title")
		tags := context.PostForm("tags")
		content := context.PostForm("content")


		articleInfo := createArticle(title, tags, content)

		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, articleInfo)
	})

	//查询文章信息

	router.GET("/api/article/:aid", func(context *gin.Context) {
		aidStr := context.Param("aid")

		aid, err := strconv.Atoi(aidStr)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}
		fmt.Println("aid is: ", aid)

		article := db.ArticleSingleFindByKV("aid", aid)

		fmt.Println("article: ", article)
		if article.Aid == aid && aid != 0{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, article)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
		}
	})
	//修改文章
	router.PATCH("/api/article", func(context *gin.Context) {
		context.Request.ParseForm()
		title := context.Request.FormValue("title")
		content := context.Request.FormValue("content")
		_tags := context.Request.FormValue("tags")
		tags := []string{}
		aidStr := context.Request.FormValue("aid")
		aid, err := strconv.Atoi(aidStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}

		article := db.ArticleSingleFindByKV("aid", aid)

		if article.Aid != aid || aid == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		if title == "" {
			title = article.Title
		}
		if content == "" {
			content = article.Content
		}
		if _tags == "" {
			tags = article.Tags
		} else {
			tags = strings.Split(_tags, ",")
		}

		err = db.ChangeArticle(aid, title, tags, 0, article.CommentN, content)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			article := db.ArticleSingleFindByKV("aid", aid)
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, article)
		}
	})
	//删除文章
	router.DELETE("/api/article/:aid", func(context *gin.Context) {

		aidStr := context.Param("aid")

		aid, err := strconv.Atoi(aidStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}

		//查询用户
		article := db.ArticleSingleFindByKV("aid", aid)

		if article.Aid != aid || aid == 0 {
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


	//分页查询所有文章

	router.GET("/api/articles", func(context *gin.Context) {
		context.DefaultQuery("pageSize", "10")
		context.DefaultQuery("pageNo", "1")

		pageSizeStr := context.Query("pageSize")
		pageNoStr := context.Query("pageNo")
		tag := context.Query("tag")

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			pageSize = 10
		}

		pageNo, err := strconv.Atoi(pageNoStr)
		if err != nil {
			pageNo = 1
		}

		var articles []db.ArticleSchema

		fmt.Println("tags is: ", tag)
		if tag != "" && tag != "全部"{
			articles, err = db.GetArticles("tags", tag, pageSize, pageNo, 0)
		} else {
			articles, err = db.GetArticles("", "", pageSize, pageNo, 0)
		}
		//articles, err := db.GetArticles("", "", pageSize, pageNo, 0)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
			})
			return
		}
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, articles)
	})

	router.GET("/api/someArticles", func(context *gin.Context) {
		fmt.Println("hello world")
		context.DefaultQuery("pageSize", "10")
		context.DefaultQuery("pageNo", "1")


		pageSizeStr := context.Query("pageSize")
		pageNoStr := context.Query("pageNo")
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
		var articles []db.ArticleSchema

		if key == "" {
			articles, err = db.GetArticles("", key, pageSize, pageNo, 0)
		} else {
			conditions := []bson.M{
				bson.M{"title": bson.M{"$regex": key, "$options": "$i"}},
				bson.M{"content": bson.M{"$regex": key, "$options": "$i"}},
				bson.M{"datestr": bson.M{"$regex": key, "$options": "$i"}},
				bson.M{"tags": bson.M{"$regex": key, "$options": "$i"}},
			}
			articles, err = db.GetSomeArticles(bson.M{"$or": conditions}, pageSize, pageNo)
		}


		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
			})
			return
		}
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, articles)
	})
}

