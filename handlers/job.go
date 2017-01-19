package handlers

import (
	"time"

	"github.com/pborman/uuid"
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

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

/*Job is the implementation with Helper of JobI*/
type Job struct {
	*BaseHandler
	Helper helpers.JobI
}

/*NewJob constructs a new Job handler*/
func NewJob(ctx *GatewayContext) *Job {
	stats := ctx.Stats.Clone(statsd.Prefix("api.job"))
	return &Job{
		Helper:      helpers.NewJob(ctx.Sql),
		BaseHandler: NewBaseHandler(stats),
	}
}

/*New creates and inserts a new job entity*/
func (j *Job) New(ctx *gin.Context) {
	var json models.Job
	err := ctx.BindJSON(&json)
	if err != nil {
		j.UserError(ctx, "Error: Unable to parse json", err)
		return
	}

	if json.SendTime == (time.Time{}) {
		json.SendTime = time.Now()
	}
	job := models.NewJob(json.Receipts, json.SendTime)
	err = j.Helper.Insert(job)
	if err != nil {
		j.ServerError(ctx, err, json)
		return
	}

	j.Success(ctx, job)
}

/*ViewAll returns a list of jobs from offset and limit params (default 0,20)*/
func (j *Job) ViewAll(ctx *gin.Context) {
	offset, limit := j.GetPaging(ctx)
	jobs, err := j.Helper.GetAll(offset, limit)
	if err != nil {
		j.ServerError(ctx, err, jobs)
		return
	}

	j.Success(ctx, jobs)
}

/*View returns one job with the given id*/
func (j *Job) View(ctx *gin.Context) {
	id := ctx.Param("jobId")

	job, err := j.Helper.GetByID(id)
	if err != nil {
		j.ServerError(ctx, err, id)
		return
	}

	j.Success(ctx, job)
}

/*Update is not implemented*/
func (j *Job) Update(ctx *gin.Context) {
	j.Success(ctx, nil)
}

/*Stop sets a job's status to Failure. Only used if the job hasn't started*/
func (j *Job) Stop(ctx *gin.Context) {
	id := ctx.Param("jobId")

	err := j.Helper.SetStatus(uuid.Parse(id), models.FAILURE)
	if err != nil {
		j.ServerError(ctx, err, id)
		return
	}

	j.Success(ctx, nil)
}
