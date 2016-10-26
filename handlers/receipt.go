package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type ReceiptI interface {
	Send(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
}

type Receipt struct {}

func (r *Receipt) Send(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (r *Receipt) ViewAll(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (r *Receipt) View(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

