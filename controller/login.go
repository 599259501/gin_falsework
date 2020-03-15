package controller

import (
	"github.com/gin-gonic/gin"
	"shop_mall/core"
	"shop_mall/core/support"
	"shop_mall/service"
)

type LoginForm struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form: "password" binding:"required"`
}

func DoLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login LoginForm
		// 获取用户名、登录密码校验身份
		support.Validate(core.GetApp(ctx), &login)

		loginService := service.NewLoginService(core.GetApp(ctx))
		if !loginService.Auth(service.RequestUser{}) {
			core.Json(ctx).Fail(500, nil, "登录失败")
			return
		}

		core.Json(ctx).Success(nil, "登录成功")
	}
}
