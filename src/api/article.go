package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"album-server/src/config"
	"album-server/src/db"
	"album-server/src/middleware"
	"album-server/src/utils"
	"net/http"
	"strconv"
)

func createArticle(title string, tags []string, content string) interface{} {
	info := db.CreateArticle(title, tags, 1, 0, content)
	return info
}
func initArticle(router *gin.Engine) {
	//创建文章
	router.POST("/api/article", middleware.JWTAuth(),func(context *gin.Context) {
		fmt.Println("create Article ....")

		var article = db.ArticleSchema{}
		context.BindJSON(&article)

		title := article.Title
		tags := article.Tags
		content := article.Content

		fmt.Println("article is: ", article)

		articleInfo := createArticle(title, tags, content)

		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, articleInfo)
	})

	//查询文章信息

	router.GET("/api/articles/:aid", func(context *gin.Context) {
		aidStr := context.Param("aid")

		aid, err := strconv.Atoi(aidStr)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}
		fmt.Println("aid is: ", aid)

		article := db.ArticleSingleFindByKV(bson.M{"aid": aid})

		fmt.Println("article: ", article)
		if article.Aid == aid && aid != 0{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, article)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
		}
	})
	//修改文章
	router.PUT("/api/articles/:aid",middleware.JWTAuth(), func(context *gin.Context) {

		var article = db.ArticleSchema{}
		context.BindJSON(&article)

		title := article.Title
		content := article.Content
		tags := article.Tags

		fmt.Println("article: is : ", article)
		aidStr := context.Param("aid")
		aid, err := strconv.Atoi(aidStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}

		article = db.ArticleSingleFindByKV(bson.M{"aid": aid})

		if article.Aid != aid || aid == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}


		fmt.Println("content is : ", content)
		fmt.Println("title is: ", title)
		fmt.Println("tags is: ", tags)
		if title == "" {
			title = article.Title
		}
		if content == "" {
			content = article.Content
		}
		if len(tags) == 0 {
			tags = article.Tags
		}

		err = db.ChangeArticle(aid, title, tags, true, article.CommentN, content)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			article := db.ArticleSingleFindByKV(bson.M{"aid": aid})
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, article)
		}
	})
	//删除文章
	router.DELETE("/api/articles/:aid",middleware.JWTAuth(), func(context *gin.Context) {

		aidStr := context.Param("aid")

		aid, err := strconv.Atoi(aidStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的aid: " + aidStr,
			})
			return

		}

		//查询用户
		article := db.ArticleSingleFindByKV(bson.M{"aid": aid})

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
		pageSizeStr := context.DefaultQuery("pageSize", config.DEFAULTPAGESIZI)
		pageNoStr := context.DefaultQuery("pageNo", config.DEFAULTPAGENO)
		isAll := context.DefaultQuery("isAll", "0")

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
		var articles db.ResData

		fmt.Println("tags is: ", tag)
		if isAll == "1" {
			if tag != "" && tag != "全部"{
				articles, err = db.GetArticles(bson.M{"tags": tag}, pageSize, pageNo)
			} else {
				articles, err = db.GetArticles(nil, pageSize, pageNo)
			}
		} else {
			if tag != "" && tag != "全部"{
				articles, err = db.GetArticles(bson.M{"tags": tag, "ispublish": true}, pageSize, pageNo)
			} else {
				fmt.Println()
				articles, err = db.GetArticles(bson.M{"ispublish": true}, pageSize, pageNo)
			}
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
		pageSizeStr := context.DefaultQuery("pageSize", config.DEFAULTPAGESIZI)
		pageNoStr := context.DefaultQuery("pageNo", config.DEFAULTPAGENO)
		isAll := context.DefaultQuery("isAll", "0")


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
		var articles db.ResData

		if isAll == "1" {
			if key == "" {
				articles, err = db.GetArticles(nil, pageSize, pageNo)
			} else {
				conditions := []bson.M{
					bson.M{"title": bson.M{"$regex": key, "$options": "$i"}},
					bson.M{"content": bson.M{"$regex": key, "$options": "$i"}},
					bson.M{"datestr": bson.M{"$regex": key, "$options": "$i"}},
					bson.M{"tags": bson.M{"$regex": key, "$options": "$i"}},
				}
				articles, err = db.GetSomeArticles(bson.M{"$or": conditions}, pageSize, pageNo)
			}
		} else {
			if key == "" {
				articles, err = db.GetArticles(bson.M{"ispublish": true}, pageSize, pageNo)
			} else {
				conditions := []bson.M{
					bson.M{"title": bson.M{"$regex": key, "$options": "$i"}},
					bson.M{"content": bson.M{"$regex": key, "$options": "$i"}},
					bson.M{"datestr": bson.M{"$regex": key, "$options": "$i"}},
					bson.M{"tags": bson.M{"$regex": key, "$options": "$i"}},
					bson.M{"ispublish": true},
				}
				articles, err = db.GetSomeArticles(bson.M{"$or": conditions}, pageSize, pageNo)
			}
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

