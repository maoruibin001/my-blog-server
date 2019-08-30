package main

import (
	"github.com/gin-gonic/gin"
	"album-server/src/api"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("/Users/ruibin/go/src/album-server/src/templates/*")
	router.Use(gin.Logger())
	//router.Use(middle())
	api.InitRouter(router)
	router.Run(":3443")
}
