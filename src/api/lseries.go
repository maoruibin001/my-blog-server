package api


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"album-server/src/config"
	"album-server/src/db"
	// "album-server/src/middleware"
	"album-server/src/utils"
	"net/http"
	"strconv"
)

func createLseries(bId int, name, mainImg, mainImgThumb string) interface{} {
	info := db.CreateLseries(bId, name, mainImg, mainImgThumb)
	return info
}
func initLseries(router *gin.Engine) {
	//创建大系列
	router.POST("/api/lseries", func(context *gin.Context) {
		fmt.Println("create lseries ....")

		var lseries = db.LseriesSchema{}
		context.ShouldBind(&lseries)

		bId := lseries.BId
		name := lseries.Name
		mainImg := lseries.MainImg
		mainImgThumb := lseries.MainImgThumb
		fmt.Println("params is: ",name)
		lseriesInfo := createLseries(bId, name, mainImg, mainImgThumb)

		fmt.Println("lseries: ", lseriesInfo)
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, lseriesInfo)
	})

	//创建大系列
	router.POST("/api/lseries/move", func(context *gin.Context) {
		fmt.Println("move lseries ....")

		var lseries = db.LseriesMoveSchema{}
		context.ShouldBind(&lseries)

		startId := lseries.Start
		endId := lseries.End
		var err error
		start := db.LseriesSingleFindByKV(bson.M{"lid": startId})
		end := db.LseriesSingleFindByKV(bson.M{"lid": endId})
		fmt.Println(start, end)

		err = db.ChangeLseries(start.LId, start.Name, start.MainImg, start.MainImgThumb, end.Seq)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "移动错误，请重试: ",
			})
			return

		}
		err = db.ChangeLseries(end.LId, end.Name, end.MainImg, end.MainImgThumb, start.Seq)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "移动错误，请重试: ",
			})
			return

		}

		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
	})

	//查询产品信息

	router.GET("/api/lseries/:id", func(context *gin.Context) {
		idStr := context.Param("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "id错误，请输入正确的id: " + idStr,
			})
			return

		}
		fmt.Println("id is: ", id)

		result := db.LseriesSingleFindByKV(bson.M{"lid": id})

		fmt.Println("result: ", result)
		var childLseries db.LRetData
		childLseries, err = db.GetLserieses(bson.M{"bid": result.LId}, 1000, 1)

		result.Children = childLseries.Lseries
		if result.LId == id && id != 0{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, result)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
		}
	})
	//修改产品
	router.PUT("/api/lseries/:id", func(context *gin.Context) {

		var lseries = db.LseriesSchema{}
		context.ShouldBind(&lseries)
		name := lseries.Name
		mainImg := lseries.MainImg
		mainImgThumb := lseries.MainImgThumb


		idStr := context.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的id: " + idStr,
			})
			return

		}
		fmt.Println("id: ", id)
		old := db.LseriesSingleFindByKV(bson.M{"lid": id})

		if old.LId != id || id == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		fmt.Println("params: ", id, name, mainImg, mainImgThumb, old.Seq)
		err = db.ChangeLseries(id, name, mainImg, mainImgThumb, old.Seq)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			lseries := db.LseriesSingleFindByKV(bson.M{"lid": id})
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, lseries)
		}
	})
	//删除产品
	router.DELETE("/api/lseries/:id", func(context *gin.Context) {

		idStr := context.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "id错误，请输入正确的id: " + idStr,
			})
			return

		}

		//查询用户
		result := db.LseriesSingleFindByKV(bson.M{"lid": id})

		if result.LId != id || id == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		err = db.RemoveLseries("lid", id)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"message": "删除失败"	,
			})
		} else {
			db.RemoveProduct("lid", id)
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
		}
	})


	//分页查询所有草稿

	router.GET("/api/lserieses", func(context *gin.Context) {
		pageSizeStr := context.DefaultQuery("pageSize", config.DEFAULTPAGESIZI)
		pageNoStr := context.DefaultQuery("pageNo", config.DEFAULTPAGENO)

		fmt.Println("pageSizeStr", pageNoStr)
		lId := context.Query("id")

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
		var lserieses db.LRetData
		var childLserieses db.LRetData

		fmt.Println("lId is: ", lId)
		if lId != "all" {
			lserieses, err = db.GetLserieses(bson.M{}, pageSize, pageNo)
		} else if lId != "" {
			_lId, err := strconv.Atoi(lId)
			if err != nil {
				fmt.Println("err: ", err)
			}
			lserieses, err = db.GetLserieses(bson.M{"bid": _lId}, pageSize, pageNo)
		} else {
			lserieses, err = db.GetLserieses(bson.M{}, pageSize, pageNo)
			for i:=0; i < len(lserieses.Lseries);i ++ {
				fmt.Println("helo: ", )
				childLserieses, err = db.GetLserieses(bson.M{"bid": lserieses.Lseries[i].BId}, pageSize, pageNo)
				lserieses.Lseries[i].Children = childLserieses.Lseries
			}
		}


		fmt.Println("lserieses:", lserieses)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
			})
			return
		}
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, lserieses)
	})

	router.GET("/api/someLserieses", func(context *gin.Context) {
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
		// var lseries db.LRetData
	
		var lserieses db.LRetData
		var childLserieses db.LRetData
		if key == "" {
			lserieses, err = db.GetLserieses(bson.M{}, pageSize, pageNo)
		} else {
			conditions := []bson.M{
				bson.M{"name": bson.M{"$regex": key, "$options": "$i"}},
				// bson.M{"content": bson.M{"$regex": key, "$options": "$i"}},
				bson.M{"modifydatestr": bson.M{"$regex": key, "$options": "$i"}},
				// bson.M{"tags": bson.M{"$regex": key, "$options": "$i"}},
				// bson.M{"ispublish": false},
			}
			lserieses, err = db.GetSomeLserieses(bson.M{"$or": conditions}, pageSize, pageNo)
		}
	
	
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
			})
			return
		}

		for i:=0; i < len(lserieses.Lseries);i ++ {
			childLserieses, err = db.GetLserieses(bson.M{"bid": lserieses.Lseries[i].BId}, pageSize, pageNo)
			if err != nil {
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
					"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
				})
				return
			}
			lserieses.Lseries[i].Children = childLserieses.Lseries
		}
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, lserieses)
	})
}
