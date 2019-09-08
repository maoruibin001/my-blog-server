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

func createBseries(name string) interface{} {
	info := db.CreateBseries(name)
	return info
}
func initBseries(router *gin.Engine) {
	//创建大系列
	router.POST("/api/bseries",middleware.JWTAuth(), func(context *gin.Context) {
		fmt.Println("create bseries ....")

		var bseries = db.BseriesSchema{}
		context.ShouldBind(&bseries)

		name := bseries.Name

		fmt.Println("params is: ",name)


		bseriesInfo := createBseries(name)

		fmt.Println("bseriesInfo: ", bseriesInfo)
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, bseriesInfo)
	})

	//创建大系列
	router.POST("/api/bseries/move", func(context *gin.Context) {
		fmt.Println("move bseries ....")

		var bseries = db.BseriesMoveSchema{}
		context.ShouldBind(&bseries)

		startId := bseries.Start
		endId := bseries.End
		var err error
		start := db.BseriesSingleFindByKV(bson.M{"bid": startId})
		end := db.BseriesSingleFindByKV(bson.M{"bid": endId})
		fmt.Println(start, end)

		err = db.ChangeBseries(start.BId, start.Name, end.Seq)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "移动错误，请重试: ",
			})
			return

		}
		err = db.ChangeBseries(end.BId, end.Name, start.Seq)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "移动错误，请重试: ",
			})
			return

		}

		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
	})

	//查询产品信息

	router.GET("/api/bseries/:id", func(context *gin.Context) {
		idStr := context.Param("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "id错误，请输入正确的id: " + idStr,
			})
			return

		}
		fmt.Println("id is: ", id)

		result := db.BseriesSingleFindByKV(bson.M{"bid": id})

		fmt.Println("result: ", result)
		var childBseries db.LRetData
		childBseries, err = db.GetLserieses(bson.M{"bid": result.BId}, 1000, 1)

		result.Children = childBseries.Lseries
		if result.BId == id && id != 0{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, result)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
		}
	})
	//修改产品
	router.PUT("/api/bseries/:id", func(context *gin.Context) {

		var bseries = db.BseriesSchema{}
		context.ShouldBind(&bseries)

		name := bseries.Name

		idStr := context.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的id: " + idStr,
			})
			return

		}

		oldProduct := db.BseriesSingleFindByKV(bson.M{"bid": id})

		if oldProduct.BId != id || id == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		err = db.ChangeBseries(id, name, oldProduct.Seq)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			bseries := db.BseriesSingleFindByKV(bson.M{"bid": id})
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, bseries)
		}
	})
	//删除产品
	router.DELETE("/api/bseries/:id",middleware.JWTAuth(), func(context *gin.Context) {

		idStr := context.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "id错误，请输入正确的id: " + idStr,
			})
			return

		}

		//查询用户
		result := db.BseriesSingleFindByKV(bson.M{"bid": id})

		if result.BId != id || id == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		err = db.RemoveBseries("bid", id)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"message": "删除失败"	,
			})
		} else {
			db.RemoveLseries("bid", id)
			db.RemoveProduct("bid", id)
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
		}
	})


	//分页查询所有草稿

	router.GET("/api/bserieses", func(context *gin.Context) {
		pageSizeStr := context.DefaultQuery("pageSize", config.DEFAULTPAGESIZI)
		pageNoStr := context.DefaultQuery("pageNo", config.DEFAULTPAGENO)

		fmt.Println("pageSizeStr", pageNoStr)
		bId := context.Query("id")

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
		var bserieses db.BRetData
		var childBserieses db.LRetData

		fmt.Println("bId is: ", bId)
		if bId == "only" {
			bserieses, err = db.GetBserieses(bson.M{}, pageSize, pageNo)
		} else if bId != "" {
			_bId, err := strconv.Atoi(bId)
			if err != nil {
				fmt.Println("err: ", err)
			}
			bserieses, err = db.GetBserieses(bson.M{"bid": _bId}, pageSize, pageNo)
		} else {
			bserieses, err = db.GetBserieses(bson.M{}, pageSize, pageNo)
			for i:=0; i < len(bserieses.Bseries);i ++ {
				childBserieses, err = db.GetLserieses(bson.M{"bid": bserieses.Bseries[i].BId}, pageSize, pageNo)

				bserieses.Bseries[i].Children = childBserieses.Lseries
			}
		}


		fmt.Println("products:", bserieses)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
			})
			return
		}
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, bserieses)
	})

	router.GET("/api/someBserieses", func(context *gin.Context) {
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
		var bserieses db.BRetData
		var childBserieses db.LRetData
	
		if key == "" {
			bserieses, err = db.GetBserieses(bson.M{}, pageSize, pageNo)
		} else {
			conditions := []bson.M{
				bson.M{"name": bson.M{"$regex": key, "$options": "$i"}},
				// bson.M{"content": bson.M{"$regex": key, "$options": "$i"}},
				bson.M{"modifydatestr": bson.M{"$regex": key, "$options": "$i"}},
				// bson.M{"tags": bson.M{"$regex": key, "$options": "$i"}},
				// bson.M{"ispublish": false},
			}
			bserieses, err = db.GetSomeBserieses(bson.M{"$or": conditions}, pageSize, pageNo)
			
		}
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
			})
			return
		}
		for i:=0; i < len(bserieses.Bseries);i ++ {
			childBserieses, err = db.GetLserieses(bson.M{"bid": bserieses.Bseries[i].BId}, pageSize, pageNo)
			if err != nil {
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
					"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
				})
				return
			}
			bserieses.Bseries[i].Children = childBserieses.Lseries
		}
	
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, bserieses)
	})
}

