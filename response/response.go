package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func response(ctx *gin.Context, httpStatus, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
func ApiSuccess(ctx *gin.Context, data gin.H, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"data": data,
		"msg":  msg,
	})
}
func ApiFail(ctx *gin.Context, data gin.H, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 4000,
		"data": data,
		"msg":  msg,
	})
}
