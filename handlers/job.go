package handlers

import (
	"time"

	"github.com/pborman/uuid"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

/*JobI is the interface for job endpoints*/
type JobI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Stop(ctx *gin.Context)
}

/*Job is the implimentation with helper of JobI*/
type Job struct {
	helper *helpers.Job
}

/*NewJob constructs a new Job handler*/
func NewJob(sql gateways.SQL) *Job {
	return &Job{helper: helpers.NewJob(sql)}
}

/*New creates and inserts a new job entity*/
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

/*ViewAll returns a list of jobs from offset and limit params (default 0,20)*/
func (j *Job) ViewAll(ctx *gin.Context) {
	offset, limit := getPaging(ctx)
	jobs, err := j.helper.GetAll(offset, limit)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}
	ctx.JSON(200, gin.H{"data": jobs})
}

/*View returns one job with the given id*/
func (j *Job) View(ctx *gin.Context) {
	id := ctx.Param("jobId")
	if id == "" {
		ctx.JSON(500, errResponse("JobId is a required parameter."))
		return
	}

	job, err := j.helper.GetJobByID(uuid.Parse(id))
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}
	ctx.JSON(200, gin.H{"data": job})
}

/*Update is not implimented*/
func (j *Job) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

/*Stop sets a job's status to Failure. Only used if the job hasn't started*/
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
