package router

import (
	"github.com/gin-gonic/gin"
	"shop_mall/controller"
	"shop_mall/core"
)

func Render(router *gin.Engine){
	// 用户登录接口
	core.Route(router, "get", "/login", controller.DoLogin())
}
