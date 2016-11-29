package handlers

import (
	"time"

	"github.com/pborman/uuid"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

type JobI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Stop(ctx *gin.Context)
}

type Job struct {
	helper *helpers.Job
}

func NewJob(sql gateways.Sql) *Job {
	return &Job{helper: helpers.NewJob(sql)}
}

func (j *Job) New(ctx *gin.Context) {
	var json models.Job
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(400, errResponse("Invalid Job input."))
		return
	}

	if json.SendTime == (time.Time{}) {
		json.SendTime = time.Now()
	}
	job := models.NewJob(json.Receipts, json.SendTime)
	err = j.helper.Insert(job)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}
	ctx.JSON(200, gin.H{"data": job})
}

func (j *Job) ViewAll(ctx *gin.Context) {
	jobs, err := j.helper.GetAll()
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}
	ctx.JSON(200, gin.H{"data": jobs})
}

func (j *Job) View(ctx *gin.Context) {
	id := ctx.Param("jobId")
	if id == "" {
		ctx.JSON(500, errResponse("JobId is a required parameter."))
		return
	}

	job, err := j.helper.GetJobById(uuid.Parse(id))
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}
	ctx.JSON(200, gin.H{"data": job})
}

func (j *Job) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (j *Job) Stop(ctx *gin.Context) {
	id := ctx.Param("jobId")
	if id == "" {
		ctx.JSON(400, errResponse("JobId is a required parameter"))
		return
	}

	err := j.helper.SetSendStatus(uuid.Parse(id), models.FAILURE)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, empty())
}
