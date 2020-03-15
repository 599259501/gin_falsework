package router

import (
	"github.com/gin-gonic/gin"
	"shop_mall/controller"
	"shop_mall/core/interface"
)

func Render(router *gin.Engine) {
	// 用户登录接口
	_interface.Route(router, "get", "/login", controller.DoLogin())
}
