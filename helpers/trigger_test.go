package helpers

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTriggerGetByKeySuccess(t *testing.T) {
	assert := assert.New(t)

	trigger := getDefaultTrigger()
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectQuery("SELECT id, contentId, tkey, params FROM b_trigger").
		WithArgs(trigger.Key).
		WillReturnRows(getTriggerRows().AddRow(trigger.ID.String(), trigger.ContentID.String(), trigger.Key, ""))

	res, err := h.GetByKey(trigger.Key)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(trigger.ID, res.ID)
	assert.EqualValues(trigger.ContentID, res.ContentID)
	assert.EqualValues(trigger.Key, res.Key)
	assert.Equal(0, len(trigger.Params))
}

func TestTriggerGetByKeyFail(t *testing.T) {
	assert := assert.New(t)

	trigger := getDefaultTrigger()
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectQuery("SELECT id, contentId, tkey, params FROM b_trigger").
		WithArgs(trigger.Key).
		WillReturnError(fmt.Errorf("some error"))

	_, err := h.GetByKey(trigger.Key)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestTriggerGetAllSuccess(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	trigger := getDefaultTrigger()
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectQuery("SELECT id, contentId, tkey, params FROM b_trigger").
		WithArgs(offset, limit).
		WillReturnRows(getTriggerRows().
			AddRow(trigger.ID.String(), trigger.ContentID.String(), trigger.Key, "").
			AddRow(trigger.ID.String(), trigger.ContentID.String(), trigger.Key, ""))

	res, err := h.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(2, len(res))
}

func TestTriggerGetAllFail(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectQuery("SELECT id, contentId, tkey, params FROM b_trigger").
		WithArgs(offset, limit).
		WillReturnError(fmt.Errorf("some error"))

	_, err := h.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestTriggerInsertSuccess(t *testing.T) {
	assert := assert.New(t)

	trigger := getDefaultTrigger()
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectPrepare("INSERT INTO b_trigger").
		ExpectExec().
		WithArgs(trigger.ID, trigger.ContentID, trigger.Key, "").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := h.Insert(trigger)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestTriggerInsertFail(t *testing.T) {
	assert := assert.New(t)

	trigger := getDefaultTrigger()
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectPrepare("INSERT INTO b_trigger").
		ExpectExec().
		WithArgs(trigger.ID, trigger.ContentID, trigger.Key, "").
		WillReturnError(fmt.Errorf("some error"))

	err := h.Insert(trigger)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestTriggerUpdateSuccess(t *testing.T) {
	assert := assert.New(t)

	trigger := getDefaultTrigger()
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectPrepare("UPDATE b_trigger").
		ExpectExec().
		WithArgs(trigger.ContentID, "", trigger.Key).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := h.Update(trigger.Key, trigger.ContentID, trigger.Params)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestTriggerUpdateFail(t *testing.T) {
	assert := assert.New(t)

	trigger := getDefaultTrigger()
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectPrepare("UPDATE b_trigger").
		ExpectExec().
		WithArgs(trigger.ContentID, "", trigger.Key).
		WillReturnError(fmt.Errorf("some error"))

	err := h.Update(trigger.Key, trigger.ContentID, trigger.Params)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestTriggerDeleteSuccess(t *testing.T) {
	assert := assert.New(t)

	trigger := getDefaultTrigger()
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectPrepare("DELETE FROM b_trigger").
		ExpectExec().
		WithArgs(trigger.Key).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := h.Delete(trigger.Key)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestTriggerDeleteFail(t *testing.T) {
	assert := assert.New(t)

	trigger := getDefaultTrigger()
	s, mock, _ := sqlmock.New()
	h := getMockTrigger(s)

	mock.ExpectPrepare("DELETE FROM b_trigger").
		ExpectExec().
		WithArgs(trigger.Key).
		WillReturnError(fmt.Errorf("some error"))

	err := h.Delete(trigger.Key)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func getDefaultTrigger() *models.Trigger {
	return models.NewTrigger(uuid.NewUUID(), "DEFAULT", make([]string, 0))
}

func getTriggerRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "contentId", "tkey", "params"})
}

func getMockTrigger(s *sql.DB) TriggerI {
	return NewTrigger(&gateways.MySQL{DB: s})
}
