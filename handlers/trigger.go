package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type TriggerI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Remove(ctx *gin.Context)
	Activate(ctx *gin.Context)
}

type Trigger struct {}

func (t *Trigger) New(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (t *Trigger) ViewAll(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (t *Trigger) View(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (t *Trigger) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (t *Trigger) Remove(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (t *Trigger) Activate(ctx *gin.Context) {
	ctx.JSON(200, empty())
}
