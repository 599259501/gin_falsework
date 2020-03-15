package _interface

import (
	"github.com/gin-gonic/gin"
	"shop_mall/core/middleware"
	"strings"
)

var beforeMiddleware []gin.HandlerFunc = []gin.HandlerFunc{
	middleware.ExceptionHandle(),
	middleware.AmountApp(),
}

var afterMiddleware []gin.HandlerFunc = []gin.HandlerFunc{}

var GRouter *gin.Engine

func init() {
	GRouter = gin.New()
	renderBeforeMiddleware()
}

func renderBeforeMiddleware() *gin.Engine {
	GRouter.Use(gin.Logger()).Use(gin.Recovery()).Use(middleware.ExceptionHandle()).Use(beforeMiddleware...)
	return GRouter
}

func Route(router *gin.Engine, method string, relativePath string, handlerFunc ...gin.HandlerFunc) {
	switch strings.ToLower(method) {
	case "get":
		router.GET(relativePath, handlerFunc...).Use(afterMiddleware...)
	case "post":
		router.POST(relativePath, handlerFunc...).Use(afterMiddleware...)
	default:
		panic("不支持该方式的请求方式")
	}
}
