package api

import (
	"album-server/src/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)


const Host  = "0.0.0.0:3433"

var simpleHostProxy = httputil.ReverseProxy{
	Director: func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = Host
		req.Host = Host
	},
}
func WithHeader(ctx *gin.Context) {
	ctx.Request.Header.Add("requester-uid", "id")
	simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func initImage(router *gin.Engine) {
	vi := router.Group("/api/upload", middleware.ImageRes())
	vi.Any("/image", WithHeader)
	vi.Any("/file", WithHeader)
	vu := router.Group("/api/download", middleware.ImageRes())
	vu.Any("/image", WithHeader)
	vu.Any("/file", WithHeader)
}
