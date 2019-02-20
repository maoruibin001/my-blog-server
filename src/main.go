package main

import (
	"github.com/gin-gonic/gin"
	"my-blog-server/src/api"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("/Users/ruibin/go/src/my-blog-server/src/templates/*")
	router.Use(gin.Logger())
	//router.Use(middle())
	api.InitRouter(router)
	router.Run(":8082")
}