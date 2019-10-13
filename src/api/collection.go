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


func initCollection(router *gin.Engine) {
	//收藏
	router.POST("/api/collection", func(context *gin.Context) {
		fmt.Println("create collection ....")

		var collection = db.CollectionSchema{}
		context.ShouldBind(&collection)
		pid := collection.Id
		uid := collection.Uid

		fmt.Println("params is: ",)

		ret, err := db.GetCollections(bson.M{"pid": pid, "uid": uid}, 10, 1)

		if err != nil || len(ret.Collections) != 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "该产品用户已收藏",
			})
			return

		}
		collectionInfo := db.CreateCollection(pid, uid)

		fmt.Println("collectionInfo: ", collectionInfo)
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, collectionInfo)
	})
	//查询产品信息

	router.GET("/api/collection/:id", func(context *gin.Context) {
		idStr := context.Param("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "id错误，请输入正确的id: " + idStr,
			})
			return

		}
		fmt.Println("id is: ", id)

		collection := db.CollectionSingleFindByKV(bson.M{"id": id})

		fmt.Println("collection: ", collection)
		if collection.Id == id && id != 0{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, collection)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
		}
	})
	//删除收藏
	router.POST("/api/deleteCollection",  middleware.JWTAuth(),func(context *gin.Context) {

		var collection = db.CollectionSchema{}
		context.ShouldBind(&collection)
		pid := collection.Id
		uid := collection.Uid

		fmt.Println("params is: ",)

		ret, err := db.GetCollections(bson.M{"pid": pid, "uid": uid}, 10, 1)

		if err != nil || len(ret.Collections) != 1 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "用户没有收藏该产品",
			})
			return

		}

		err = db.RemoveCollection(bson.M{"pid": pid, "uid": uid})

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"message": "删除失败"	,
			})
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
		}
	})


	//分页查询所有草稿

	router.GET("/api/collections", func(context *gin.Context) {
		pageSizeStr := context.DefaultQuery("pageSize", config.DEFAULTPAGESIZI)
		pageNoStr := context.DefaultQuery("pageNo", config.DEFAULTPAGENO)

		pid := context.Query("id")
		uid := context.Query("uid")


		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			pageSize = 10
			fmt.Println("err: ", err)
		}

		pageNo, err := strconv.Atoi(pageNoStr)
		if err != nil {
			pageNo = 1
		}

		fmt.Println("pagesize: ", pageSize, pageNo)
		var collections db.CollectionRetData

		if pid == "" && uid == "" {
			collections, err = db.GetCollections(bson.M{}, pageSize, pageNo)
		} else if pid != ""{
			_pid, err := strconv.Atoi(pid)
			if err != nil {
				fmt.Println("err: ", err)
			}

			collections, err = db.GetCollections(bson.M{"pid": _pid}, pageSize, pageNo)

		} else  {
			_uid, err := strconv.Atoi(uid)
			if err != nil {
				fmt.Println("err: ", err)
			}

			collections, err = db.GetCollections(bson.M{"uid": _uid}, pageSize, pageNo)
		}

		fmt.Println("collections:", collections)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pid: " + pid + " uid: " + uid,
			})
			return
		}
		if (len(collections.Collections) == 0) {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, gin.H{
				"products": []string{},
			})
		} else {
			products := []db.ProductSchema{}
			for i:=0; i < len(collections.Collections); i ++ {
				id := collections.Collections[i].Pid
				product := db.ProductSingleFindByKV(bson.M{"id": id})
				products = append(products, product)
			}
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, gin.H{
				"products": products,
			})
		}

	})
}

