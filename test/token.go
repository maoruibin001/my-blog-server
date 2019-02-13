package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"my-blog-server/src/utils"
	"time"
)

const SecretKey = "I have login"
func createToken() string  {

	token := jwt.New(jwt.SigningMethodHS256)
	//claims := jwt.MapClaims{}
	claims := make(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()

	claims["iat"] = time.Now().Unix()
	claims["id"]="12"
	claims["userName"]="324"

	token.Claims = claims

	fmt.Println("claims: ", claims)

	//tokenString, err := token.SignedString([]byte(a))
	tokenString, err := token.SignedString([]byte(SecretKey))

	utils.HandleError("token 转码错误： ", err)

	fmt.Println(tokenString)


	return tokenString
}

func main() {
	createToken()
}