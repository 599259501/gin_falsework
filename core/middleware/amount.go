package middleware

import
(
	"github.com/gin-gonic/gin"
	"shop_mall/core/interface"
)

func AmountApp()gin.HandlerFunc{
	return func(context *gin.Context) {
		_interface.BindApp(context)
	}
}