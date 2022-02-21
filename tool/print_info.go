package tool

import "github.com/gin-gonic/gin"

func PrintInfo(context *gin.Context, info interface{}) {
	context.JSON(200, gin.H{
		"data": info,
	})
}

func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(200, gin.H{
		"status": false,
		"data":   v,
	})
}
