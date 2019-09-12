package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"album-server/src/db"
	"album-server/src/middleware"
	"album-server/src/utils"
	"net/http"
)

type JSONFormat struct {
	code int
	msg string
	data map[string]interface{}

}

func initLogin(router *gin.Engine) {
	router.POST("/api/login",func(context *gin.Context) {
		user := db.UserSchema{}
		context.ShouldBind(&user)
		phone := user.Phone
		password := user.Password

		fmt.Println("phone is : ", phone, user)
		result := db.FindByPhone(phone)

		fmt.Println("result: ", result)
		if result.Phone != phone || phone == "" {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
			return
		}

		if result.Password !=  utils.MD5(password) {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPASSWORDERROR, nil)
			return
		}
		j := middleware.NewJWT()
		claims := middleware.CustomClaims{result.Id, result.Name, result.Phone, jwt.StandardClaims{}}
		token, err := j.CreateToken(claims)

		fmt.Println("token is: ", token)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, nil)
		} else {
			result.Token = token
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, result)
		}
	})
}
