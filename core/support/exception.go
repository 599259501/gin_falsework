package support

import "github.com/gin-gonic/gin"

const (
	SERVER_ERROR = 500
)

type Exception struct {
	context *gin.Context
	Code int
	Msg string
	Data interface{}
}

func ThrowException(ctx *gin.Context)  *Exception{
	return &Exception{
		context: ctx,
		Code:    SERVER_ERROR,
		Msg:     "系统繁忙",
		Data:    nil,
	}
}

func (exception *Exception)SetCode(code int) *Exception{
	exception.Code = code

	return exception
}

func (exception *Exception)SetData(data interface{})*Exception  {
	exception.Data = data
	return exception
}

func (exception *Exception)SetMsg(msg string) *Exception{
	exception.Msg = msg

	return exception
}

func (exception *Exception)Panic(){
	panic(exception)
}