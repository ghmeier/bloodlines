package handlers

import (
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func GetPaging(ctx *gin.Context) (int, int) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	return offset, limit
}

func UserError(ctx *gin.Context, msg string, obj interface{}) {
	ctx.JSON(400, &gin.H{"success": false, "message": msg, "error": obj})
}

func ServerError(ctx *gin.Context, err error, obj interface{}) {
	ctx.JSON(500, &gin.H{"success": false, "message": err.Error(), "error": obj})
}

func Success(ctx *gin.Context, obj interface{}) {
	ctx.JSON(200, gin.H{"success": true, "data": obj})
}
