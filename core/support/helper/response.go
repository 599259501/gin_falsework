package helper

import (
	"github.com/gin-gonic/gin"
	response2 "shop_mall/core/support/response"
)

func Json(ctx *gin.Context) *response2.JsonResponse {
	return response2.NewJsonResponse(ctx)
}