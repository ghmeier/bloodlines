package helpers

import (
	"strings"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

type Job struct {
	*baseHelper
}

func NewJob(sql gateways.Sql) *Job {
	return &Job{baseHelper: &baseHelper{sql: sql}}
}

func (j *Job) Insert(job *models.Job) error {
	ids := make([]string, 0)
	for _, id := range job.Receipts {
		ids = append(ids, id.String())
	}
	err := j.sql.Modify("INSERT INTO job (id, sendTime, sendStatus, receipts) VALUES (?, ?, ?, ?)",
		job.Id,
		job.SendTime,
		job.SendStatus,
		strings.Join(ids, ","),
	)
	return err
}

func (j *Job) GetAll() ([]*models.Job, error) {
	rows, err := j.sql.Select("SELECT id, sendTime, sendStatus, receipts from job")
	if err != nil {
		return nil, err
	}

	jobs, err := models.JobFromSql(rows)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (j *Job) GetJobById(id uuid.UUID) (*models.Job, error) {
	rows, err := j.sql.Select("SELECT id, sendTime, sendStatus, receipts FROM job WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	jobs, err := models.JobFromSql(rows)
	if err != nil {
		return nil, err
	}

	return jobs[0], nil
}

func (j *Job) SetSendStatus(id uuid.UUID, state models.Status) error {
	err := j.sql.Modify("UPDATE receipt SET sendStatus=? where id=?", state, id)
	return err
}
