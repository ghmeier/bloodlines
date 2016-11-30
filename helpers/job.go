package helpers

import (
	"strings"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

/*Job Helper has helper methods for job entities and the database*/
type Job struct {
	*baseHelper
}

/*NewJob reutrns a constructed job helper*/
func NewJob(sql gateways.SQL) *Job {
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
		job.SendStatus,
		strings.Join(ids, ","),
	)
	return err
}

/*GetAll returns <limit> job entities starting at <offset>*/
func (j *Job) GetAll(offset int, limit int) ([]*models.Job, error) {
	rows, err := j.sql.Select("SELECT id, sendTime, sendStatus, receipts from job ORDER BY id ASC LIMIT ?,?", offset, limit)
	if err != nil {
		return nil, err
	}

	jobs, err := models.JobFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

/*GetJobByID returns one job with the give id, nil otherwise*/
func (j *Job) GetJobByID(id uuid.UUID) (*models.Job, error) {
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

/*SetSendStatus updates the status of the job with the provided id*/
func (j *Job) SetSendStatus(id uuid.UUID, state models.Status) error {
	err := j.sql.Modify("UPDATE receipt SET sendStatus=? where id=?", state, id)
	return err
}
