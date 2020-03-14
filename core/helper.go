package core

import (
	"github.com/gin-gonic/gin"
	"shop_mall/core/interface"
	"shop_mall/core/service"
	"shop_mall/core/support/helper"
	"shop_mall/core/support/response"
)

func GetApp(ctx *gin.Context) *_interface.App {
	appInterface,_ := ctx.Get("App")
	app,_ := appInterface.(*_interface.App)
	return app
}

func Session(app *_interface.App) *service.Session{
	return app.GetSession()
}

func Json(ctx *gin.Context)*response.JsonResponse{
	return helper.Json(ctx)
}

func Env(key string) string{
	return helper.Env(key, "")
}

func Config(path string)(map[string]interface{},error){
	return helper.Config(path)
}