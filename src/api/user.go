package api

import (
	"album-server/src/config"
	"album-server/src/db"
	"album-server/src/middleware"
	"album-server/src/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
//TODO: 需要一个用户信息返回过滤函数。
//TODO: 缺少一个修改密码的函数。

func createUser(name, phone, password string, isKeeper int) db.UserSchema {
	salt := utils.MD5(utils.GetRandomString(8) + config.Salt)
	pw := utils.MD5(password)
	userInfo := db.CreateUser(name, phone, pw, salt, isKeeper)
	return userInfo
}
func init() {
	c, session := db.GetCollect("userdb", "user")
	defer session.Close()
	count, err := c.Count()
	utils.HandleError("查找错误：", err)
	if count == 0 {
		fmt.Println("用户数据库为空，初始化数据...")
		createUser("admin", "13726648009", "123", 1)
	}
}
func initUser(router *gin.Engine) {
	//创建新用户
	router.POST("/api/user", middleware.JWTAuth(), func(context *gin.Context) {
		fmt.Println("add user ....")

		user := db.UserSchema{}
		context.ShouldBind(&user)
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

		utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, gin.H{
			"id": userInfo.Id,
			"name": userInfo.Name,
			"phone": userInfo.Phone,
			"isKeeper": userInfo.IsKeeper,
			"token": userInfo.Token,
		})
	})

	//查询用户信息

	router.GET("/api/user", func(context *gin.Context) {
		phone := context.Query("phone")
		idStr := context.Query("id")
		fmt.Println("phone is: ", phone)
		var user db.UserSchema
		if phone != "" {
			user = db.FindByPhone(phone)
		} else if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
					"errorMsg": "id错误，请输入正确的id: " + idStr,
				})
				return

			}
			user = db.FindById(id)
		}

		fmt.Println("userinfo: ", user)
		if user.Phone != ""{
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
	router.PUT("/api/user",  middleware.JWTAuth(),func(context *gin.Context) {
		context.Request.ParseForm()

		user := db.UserSchema{}
		context.ShouldBind(&user)

		phone := user.Phone
		name := user.Name
		password := user.Password
		id := user.Id
		isKeeper := user.IsKeeper

		keeperPhone := context.GetString("phone")
		fmt.Println("keeperPhone is: ", keeperPhone)
		//if keeperPhone == "" {
		//	utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOAUTH, nil)
		//	return
		//}
		//keeperUser := db.FindByPhone(keeperPhone)

		//fmt.Println("keeperUser: ", keeperUser.IsKeeper)
		//if keeperUser.IsKeeper == 0 {
		//	utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOAUTH, nil)
		//	return
		//}
		if isKeeper == 1 {
			fmt.Println("id: ", id)
			keeperUser := db.FindById(id)

			fmt.Println(keeperUser.Id)
			if keeperUser.Id != id || id == 0{
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
				return
			}

			if (keeperUser.IsKeeper != 1) {
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOAUTH, nil)
				return
			}
		}




		if phone == "" {
			phone = user.Phone
		}
		if name == "" {
			name = user.Name
		}
		if password == "" {
			password = user.Password
		}
		pw := utils.MD5(password)
		salt := utils.MD5(utils.GetRandomString(8) + config.Salt)

		fmt.Println("hello world")

		err := db.ChangeUser(id, name, phone, salt, pw, isKeeper)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUPDATEERROR, nil)
		} else {
			result := db.FindById(id)
			j := middleware.NewJWT()
			claims := middleware.CustomClaims{result.Id, result.Name, result.Phone, jwt.StandardClaims{}}
			token, err := j.CreateToken(claims)

			fmt.Println("token is: ", result.IsKeeper)

			if err != nil {
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, nil)
			} else {
				result.Token = token
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, gin.H{
					"id":       result.Id,
					"name":     result.Name,
					"phone":    result.Phone,
					"isKeeper": result.IsKeeper,
					"token": 	result.Token,
				})
			}
		}
	})
	//删除用户
	router.DELETE("/api/user",  middleware.JWTAuth(),func(context *gin.Context) {

		user := db.UserSchema{}
		context.ShouldBind(&user)
		phone := user.Phone
		id := user.Id

		//查询用户

		var _user db.UserSchema

		if user.Phone != "" {
			_user = db.FindByPhone(phone)
		} else if user.Id != 0 {
			_user = db.FindById(id)
		}

		if _user.Phone == "" {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
			return
		}

		err := db.RemoveUser("phone", _user.Phone)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSESERVERERROR, gin.H{
				"message": "删除失败"	,
			})
		} else {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, nil)
		}
	})
}
