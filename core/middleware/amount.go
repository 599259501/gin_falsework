package middleware

import (
	"github.com/gin-gonic/gin"
	"shop_mall/core"
)

func AmountApp() gin.HandlerFunc {
	return func(context *gin.Context) {
		app := core.NewApp(context)
		context.Set("App", app)
	}
}
