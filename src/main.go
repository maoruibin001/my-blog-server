package main

import (
	"github.com/gin-gonic/gin"
	"album-server/src/api"
	"album-server/src/middleware"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("/Users/ruibin/go/src/album-server/src/templates/*")
	router.Use(gin.Logger())
	router.Use(middleware.SelectDb())
	api.InitRouter(router)
	router.Run(":3443")
}
