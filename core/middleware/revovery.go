package middleware

import (
	"github.com/gin-gonic/gin"
	"shop_mall/core/support"
	"shop_mall/core/support/helper"
)

func ExceptionHandle() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		defer func() {
			var (
				exception *support.Exception
				isException bool
				err interface{}
			)
			if err := recover();err!=nil{
				return;
			}

			if exception,isException = err.(*support.Exception);!isException{
				panic(err)
				return
			}

			helper.Json(ctx).Fail(exception.Code, exception.Data, exception.Msg)
			return
		}()

		ctx.Next()
	}
}
