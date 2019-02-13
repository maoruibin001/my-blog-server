package main

import (
	"github.com/gin-gonic/gin"
	"my-blog-server/src/api"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("src/templates/*")
	router.Use(gin.Logger())
	//router.Use(middle())
	api.InitRouter(router)
	router.Run(":3433")
}
