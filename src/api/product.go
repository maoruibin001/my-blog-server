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

func createProduct(name,descImg,descImgThumb,gifImg,originFile string, prize, pId int, mainImgList []db.ImgInfoSchema) interface{} {
	info := db.CreateProduct(name,descImg,descImgThumb,gifImg,originFile, prize, pId, mainImgList)
	return info
}
func initProduct(router *gin.Engine) {
	//创建产品
	router.POST("/api/product",middleware.JWTAuth(), func(context *gin.Context) {
		fmt.Println("create product ....")

		var product = db.ProductSchema{}
		context.ShouldBind(&product)

		name := product.Name
		descImg := product.DescImg
		descImgThumb := product.DescImgThumb
		gifImg := product.GifImg
		originFile := product.OriginFile
		prize := product.Prize
		pId := product.PId
		mainImgList := product.MainImgList

		fmt.Println("params is: ",name,descImg,descImgThumb,gifImg,originFile, prize, pId, mainImgList)


		productInfo := createProduct(name,descImg,descImgThumb,gifImg,originFile, prize, pId, mainImgList)

		fmt.Println("productInfo: ", productInfo)
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, productInfo)
	})

	//查询产品信息

	router.GET("/api/product/:id", func(context *gin.Context) {
		idStr := context.Param("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "id错误，请输入正确的id: " + idStr,
			})
			return

		}
		fmt.Println("id is: ", id)

		product := db.ProductSingleFindByKV(bson.M{"id": id})

		fmt.Println("product: ", product)
		if product.Id == id && id != 0{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, product)
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
		}
	})
	//修改产品
	router.PUT("/api/product/:id",middleware.JWTAuth(), func(context *gin.Context) {

		var product = db.ProductSchema{}
		context.ShouldBind(&product)

		//context.Request.ParseForm()
		//title := article.Title
		//content := article.Content
		//tags := article.Tags
		name := product.Name
		descImg := product.DescImg
		descImgThumb := product.DescImgThumb
		gifImg := product.GifImg
		originFile := product.OriginFile
		prize := product.Prize
		pId := product.PId
		id := product.Id
		mainImgList := product.MainImgList

		fmt.Println("params is: ",name,descImg,descImgThumb,gifImg,originFile, prize, pId, mainImgList)
		idStr := context.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "aid错误，请输入正确的id: " + idStr,
			})
			return

		}

		oldProduct := db.ProductSingleFindByKV(bson.M{"id": id})

		if oldProduct.Id != id || id == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		//if title == "" {
		//	title = draft.Title
		//}
		//if content == "" {
		//	content = draft.Content
		//}
		//if len(tags) == 0 {
		//	tags = draft.Tags
		//}

		err = db.ChangeProduct(name,descImg,descImgThumb,gifImg,originFile, prize, pId,id, mainImgList)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			product := db.ProductSingleFindByKV(bson.M{"id": id})
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, product)
		}
	})
	//删除产品
	router.DELETE("/api/product/:id",middleware.JWTAuth(), func(context *gin.Context) {

		idStr := context.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "id错误，请输入正确的id: " + idStr,
			})
			return

		}

		//查询用户
		product := db.ProductSingleFindByKV(bson.M{"id": id})

		if product.Id != id || id == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOARTICLE, nil)
			return
		}

		err = db.RemoveProduct("id", id)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"message": "删除失败"	,
			})
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
		}
	})


	//分页查询所有草稿

	router.GET("/api/products", func(context *gin.Context) {
		pageSizeStr := context.DefaultQuery("pageSize", config.DEFAULTPAGESIZI)
		pageNoStr := context.DefaultQuery("pageNo", config.DEFAULTPAGENO)

		fmt.Println("pageSizeStr", pageNoStr)
		pId := context.Query("pId")

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
		var products db.RetData
		var childProducts db.RetData

		fmt.Println("pId is: ", pId)
		if pId != "" && pId != "全部"{
			products, err = db.GetProducts(bson.M{"pid": pId}, pageSize, pageNo)
		} else {
			products, err = db.GetProducts(bson.M{"pid": -1}, pageSize, pageNo)
			for i:=0; i < len(products.Products);i ++ {
				fmt.Println("helo: ", )
				childProducts, err = db.GetProducts(bson.M{"pid": products.Products[i].Id}, pageSize, pageNo)
				//, err := db.GetProducts(bson.M{"pid": products.Products[i].Id}, pageSize, pageNo)
				products.Products[i].Children = childProducts.Products
			}
		}

		fmt.Println("products:", products)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR,  gin.H{
				"message": "pageSize: " + pageSizeStr + "pageNo: " + pageNoStr,
			})
			return
		}
		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, products)
	})

	router.GET("/api/someProducts", func(context *gin.Context) {
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

