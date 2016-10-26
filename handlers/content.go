package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type ContentI interface {
	New(ctx*gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
}

type Content struct {}

func (c *Content) New(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (c *Content) ViewAll(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (c *Content) View(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (c *Content) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (c *Content) Deactivate(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func empty() *gin.H {
	return &gin.H{"success": true}
}
