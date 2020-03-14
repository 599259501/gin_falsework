package _interface

import (
	"github.com/gin-gonic/gin"
)

func BindApp(context *gin.Context) *App {
	app := NewApp(context)
	context.Set("App", app)

	return app
}
