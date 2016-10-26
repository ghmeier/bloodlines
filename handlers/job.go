package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type JobI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Stop(ctx *gin.Context)
}

type Job struct {}

func (j *Job) New(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (j *Job) ViewAll(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (j *Job) View(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (j *Job) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (j *Job) Stop(ctx *gin.Context) {
	ctx.JSON(200, empty())
}
