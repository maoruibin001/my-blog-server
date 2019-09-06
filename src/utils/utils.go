package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"reflect"
	"time"
)

var ResArr = [1000000]string{}

func initResCode() {
	ResArr[RESPONSEOK] = "ok"
	ResArr[RESPONSEUNLOGIN] = "用户未登陆"
	ResArr[RESPONSENOUSER] = "用户不存在"
	ResArr[RESPONSEUSEREXSIST] = "此用户已存在"
	ResArr[RESPONSENOARTICLE] = "内容不存在"
	ResArr[RESPONSEPARAMERROR] = "请求参数错误"
	ResArr[RESPONSESERVERERROR] = "操作有误"
	ResArr[RESPONSEUPDATEERROR] = "数据更新失败"
	ResArr[RESPONSETOKENINVALID] = "token失效"
	ResArr[RESPONSENOTOKEN] = "没有token"

}

func init() {
	initResCode()
}

func HandleError(profileStr string, err error, args ...interface{}) {
	if err != nil {
		log.Fatal(profileStr, err, "\n")
	}
}

func CreateToken(id, name, password string) (string, error)  {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	claims["iat"] = time.Now().Unix()
	claims["id"] = id
	claims["name"] = name
	claims["password"] = password

	token.Claims = claims

	tokenString, err := token.SignedString(nil)

	fmt.Println(tokenString)

	return tokenString, err
}


func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)

	ret := []byte{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l ; i ++ {

		ret = append(ret, bytes[rand.Intn(len(str))])
	}
	return string(ret)

}

func MD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))

	fmt.Println("text: ", text)
	fmt.Println("md5: ", hex.EncodeToString(ctx.Sum(nil)))
	return hex.EncodeToString(ctx.Sum(nil))
}



func ResponseJson(context *gin.Context, code, retCode int, data interface{})  {
	var msg = ResArr[retCode]

	if msg == "" {
		msg = "unKnown"
	}
	var result = map[string]interface{} {
		"retCode": retCode,
		"msg": msg,
		"data": data,
	}

	log.Println(code, retCode)
	fmt.Println("data", data)
	context.JSON(code, result)
}


func GetType(param interface{}) string {
	return fmt.Sprint(reflect.TypeOf(param))
}