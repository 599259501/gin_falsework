package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shop_mall/core"
	"shop_mall/core/support"
)

type LoginForm struct {
	UserName string `form:"email" binding:"required"`
	Password string `form: "password" binding:"required"`
}

func DoLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("开始处理login请求")
		var login LoginForm
		// 获取用户名、登录密码校验身份
		support.Validate(core.GetApp(ctx), &login)
	}
}
