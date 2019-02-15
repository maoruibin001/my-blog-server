package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my-blog-server/src/db"
	"my-blog-server/src/middleware"
	"my-blog-server/src/utils"
	"net/http"
)

type JSONFormat struct {
	code int
	msg string
	data map[string]interface{}

}

func initLogin(router *gin.Engine) {
	router.POST("/api/login", func(context *gin.Context) {
		user := db.UserSchema{}
		context.BindJSON(&user)
		phone := user.Phone

		fmt.Println("phone is : ", phone, user)
		result := db.FindByPhone(phone)

		if result.Phone != phone || phone == "" {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
			return
		}

		j := middleware.NewJWT()
		claims := middleware.CustomClaims{result.Id, result.Name, result.Phone, jwt.StandardClaims{}}
		token, err := j.CreateToken(claims)

		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, nil)
		} else {
			result.Token = token
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, result)
		}
	})
}
