package middlewares

import (
	"github.com/gin-gonic/gin"
	"shop_mall/core/support/helper"
)

func FormatResponse()gin.HandlerFunc  {
	return func(context *gin.Context) {
		helper.Json(context).Fail(500, nil,"这是一个测试错误")
	}
}
