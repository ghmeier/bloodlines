package helpers

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJobGetByIDSuccess(t *testing.T) {
	assert := assert.New(t)

	job := getDefaultJob()
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectQuery("SELECT id, sendTime, sendStatus, receipts FROM job").
		WithArgs(job.ID.String()).
		WillReturnRows(getJobRows().AddRow(job.ID.String(), job.SendTime, string(job.SendStatus), ""))

	res, err := c.GetJobByID(job.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(job.ID, res.ID)
	assert.Equal(0, len(res.Receipts))
	assert.EqualValues(job.SendStatus, res.SendStatus)
	assert.Equal(job.SendTime, res.SendTime)
}

func TestJobGetByIDQueryFail(t *testing.T) {
	assert := assert.New(t)

	job := getDefaultJob()
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectQuery("SELECT id, sendTime, sendStatus, receipts FROM job").
		WithArgs(job.ID.String()).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.GetJobByID(job.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestJobGetByIDMapFail(t *testing.T) {
	assert := assert.New(t)

	job := getDefaultJob()
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectQuery("SELECT id, sendTime, sendStatus, receipts FROM job").
		WithArgs(job.ID.String()).
		WillReturnRows(getJobRows().AddRow(job.ID.String(), job.SendTime, "INVALID", ""))

	_, err := c.GetJobByID(job.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestJobGetAllSuccess(t *testing.T) {
	assert := assert.New(t)

	job := getDefaultJob()
	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectQuery("SELECT id, sendTime, sendStatus, receipts FROM job").
		WithArgs(offset, limit).
		WillReturnRows(getJobRows().
			AddRow(job.ID.String(), job.SendTime, string(job.SendStatus), "").
			AddRow(job.ID.String(), job.SendTime, string(job.SendStatus), ""))

	res, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(2, len(res))
}

func TestJobGetAllMapFail(t *testing.T) {
	assert := assert.New(t)

	job := getDefaultJob()
	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectQuery("SELECT id, sendTime, sendStatus, receipts FROM job").
		WithArgs(offset, limit).
		WillReturnRows(getJobRows().
			AddRow(job.ID.String(), job.SendTime, "INVALID", ""))

	_, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestJobGetAllFail(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectQuery("SELECT id, sendTime, sendStatus, receipts FROM job").
		WithArgs(offset, limit).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestJobInsertSuccess(t *testing.T) {
	assert := assert.New(t)

	job := getDefaultJob()
	job.Receipts = []uuid.UUID{uuid.NewUUID()}
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectPrepare("INSERT INTO job").
		ExpectExec().
		WithArgs(job.ID, job.SendTime, string(job.SendStatus), job.Receipts[0].String()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.Insert(job)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestJobInsertFail(t *testing.T) {
	assert := assert.New(t)

	job := getDefaultJob()
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectPrepare("INSERT INTO job").
		ExpectExec().
		WithArgs(job.ID, job.SendTime, string(job.SendStatus), "").
		WillReturnError(fmt.Errorf("some error"))

	err := c.Insert(job)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestJobSetSendStatusSuccess(t *testing.T) {
	assert := assert.New(t)

	job := getDefaultJob()
	job.Receipts = []uuid.UUID{uuid.NewUUID()}
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectPrepare("UPDATE job").
		ExpectExec().
		WithArgs(string(job.SendStatus), job.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.SetSendStatus(job.ID, job.SendStatus)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestJobSetSendStatusFail(t *testing.T) {
	assert := assert.New(t)

	job := getDefaultJob()
	s, mock, _ := sqlmock.New()
	c := getMockJob(s)

	mock.ExpectPrepare("UPDATE job").
		ExpectExec().
		WithArgs(string(job.SendStatus), job.ID).
		WillReturnError(fmt.Errorf("some error"))

	err := c.SetSendStatus(job.ID, job.SendStatus)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func getDefaultJob() *models.Job {
	return models.NewJob(make([]uuid.UUID, 0), time.Now())
}

func getJobRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "sendTime", "sendStatus", "receipts"})
}

func getMockJob(s *sql.DB) *Job {
	return NewJob(&gateways.MySQL{DB: s})
}
