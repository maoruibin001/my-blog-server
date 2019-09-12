package api

import (
	"album-server/src/config"
	"album-server/src/db"
	"album-server/src/middleware"
	"album-server/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//TODO: 需要一个用户信息返回过滤函数。
//TODO: 缺少一个修改密码的函数。

func createUser(name, phone, password string, isKeeper int) interface{} {
	salt := utils.MD5(utils.GetRandomString(8) + config.Salt)
	pw := utils.MD5(password)
	userInfo := db.CreateUser(name, phone, pw, salt, isKeeper)
	return userInfo
}
func init() {
	fmt.Println("hello world")
	c, session := db.GetCollect(utils.GetDbName(), "user")
	defer session.Close()
	count, err := c.Count()
	utils.HandleError("查找错误：", err)
	if count == 0 {
		fmt.Println("用户数据库为空，初始化数据...")
		createUser("admin", "123", "123", 1)
	}
}
func initUser(router *gin.Engine) {
	//创建新用户
	router.POST("/api/user", middleware.JWTAuth(), func(context *gin.Context) {
		fmt.Println("add user ....")

		user := db.UserSchema{}
		context.ShouldBind(&user)
		fmt.Println("param: ", user)
		name := user.Name
		phone := user.Phone
		password := user.Password
		isKeeper := user.IsKeeper

		keeperPhone := context.GetString("phone")
		fmt.Println(keeperPhone)

		if keeperPhone == "" {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOAUTH, nil)
			return
		}
		keeperUser := db.FindByPhone(keeperPhone)

		if keeperUser.IsKeeper == 0 {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOAUTH, nil)
			return
		}

		checkUser := db.FindByPhone(phone)

		if checkUser.Phone == phone && user.Phone != ""{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUSEREXSIST, nil)
			return
		}



		userInfo := createUser(name, phone, password, isKeeper)

		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, userInfo)
	})

	//查询用户信息

	router.GET("/api/user", func(context *gin.Context) {
		phone := context.Query("phone")
		fmt.Println("phone is: ", phone)
		user := db.FindByPhone(phone)
		fmt.Println("userinfo: ", user)
		if user.Phone == phone && phone != ""{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, gin.H{
				"name": user.Name,
				"phone": user.Phone,
				"id": user.Id,
			})
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
		}
	})
	//修改用户
	router.PUT("/api/user", func(context *gin.Context) {
		context.Request.ParseForm()

		user := db.UserSchema{}
		context.BindJSON(&user)
		phone := user.Phone
		name := user.Name
		//password := context.Request.FormValue("password")
		id := user.Id
		isKeeper := user.IsKeeper
		checkUser := db.FindById(id)

		if checkUser.Id != id || id == 0{
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
			return
		}

		if phone == "" {
			phone = user.Phone
		}
		if name == "" {
			name = user.Name
		}
		//if password == "" {
		//	password = user.Password
		//}
		salt := checkUser.Salt


		err := db.ChangeUser(id, name, phone, salt, isKeeper)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			user = db.FindById(id)
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, gin.H{
				"id": user.Id,
				"name": user.Name,
				"phone": user.Phone,
				"isKeeper": user.IsKeeper,
			})
		}
	})
	//删除用户
	router.DELETE("/api/user/:phone", func(context *gin.Context) {

		phone := context.Param("phone")

		//查询用户
		fmt.Println("phone2 is: ", phone)
		user := db.FindByPhone(phone)

		if user.Phone != phone || phone == "" {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
			return
		}

		err := db.RemoveUser("phone", phone)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"message": "删除失败"	,
			})
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
		}
	})
}
