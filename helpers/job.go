package helpers

import (
	"strings"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

/*JobI describes the methods of job helpers*/
type JobI interface {
	GetAll(int, int) ([]*models.Job, error)
	GetByID(string) (*models.Job, error)
	SetStatus(uuid.UUID, models.Status) error
	Insert(*models.Job) error
}

/*Job Helper has helper methods for job entities and the database*/
type Job struct {
	*baseHelper
}

/*NewJob reutrns a constructed job helper*/
func NewJob(sql gateways.SQL) JobI {
	return &Job{baseHelper: &baseHelper{sql: sql}}
}

/*Insert creates a new job from the arguments and errors if failure*/
func (j *Job) Insert(job *models.Job) error {
	ids := make([]string, 0)
	for _, id := range job.Receipts {
		ids = append(ids, id.String())
	}
	err := j.sql.Modify("INSERT INTO job (id, sendTime, sendStatus, receipts) VALUES (?, ?, ?, ?)",
		job.ID,
		job.SendTime,
		string(job.SendStatus),
		strings.Join(ids, ","),
	)
	return err
}

/*GetAll returns <limit> job entities starting at <offset>*/
func (j *Job) GetAll(offset int, limit int) ([]*models.Job, error) {
	rows, err := j.sql.Select("SELECT id, sendTime, sendStatus, receipts FROM job ORDER BY id ASC LIMIT ?,?", offset, limit)
	if err != nil {
		return nil, err
	}

	jobs, err := models.JobFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

/*GetByID returns one job with the give id, nil otherwise*/
func (j *Job) GetByID(id string) (*models.Job, error) {
	rows, err := j.sql.Select("SELECT id, sendTime, sendStatus, receipts FROM job WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	jobs, err := models.JobFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return jobs[0], nil
}

/*SetStatus updates the status of the job with the provided id*/
func (j *Job) SetStatus(id uuid.UUID, state models.Status) error {
	err := j.sql.Modify("UPDATE job SET sendStatus=? where id=?", string(state), id)
	return err
}

/*SendJob should queue messages based on a job, not sure whether to
  do this by jobID or receipt yet*/
func (j *Job) SendJob() error {
	return nil
}
