package response

import "github.com/gin-gonic/gin"

type ResponseStruct struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

type HttpResponse interface {
	Success(data interface{}, msg string)
	Fail(code int, data interface{}, msg string)
}

type JsonResponse struct {
	context *gin.Context
}

func NewJsonResponse(ctx *gin.Context) *JsonResponse {
	return &JsonResponse{
		context: ctx,
	}
}

func (json *JsonResponse)Success(data interface{}, msg string){
	json.context.JSON(200, ResponseStruct{
		Msg:msg,
		Code:0,
		Data:data,
	})

	json.context.Abort()
}

func (json *JsonResponse)Fail(code int,data interface{}, msg string){
	json.context.JSON(200, ResponseStruct{
		Msg:  msg,
		Code: code,
		Data: data,
	})

	json.context.Abort()
}
