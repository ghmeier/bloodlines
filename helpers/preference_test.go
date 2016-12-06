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

func TestPreferenceGetByUserIDSuccess(t *testing.T) {
	assert := assert.New(t)

	p := getDefaultPreference()
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectQuery("SELECT id, userId, email FROM preference").
		WithArgs(p.UserID.String()).
		WillReturnRows(getPreferenceRows().AddRow(p.ID.String(), p.UserID.String(), string(p.Email)))

	res, err := c.GetByUserID(p.UserID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(p.ID, res.ID)
	assert.EqualValues(p.Email, res.Email)
	assert.Equal(p.UserID, res.UserID)
}

func TestPreferenceGetByUserIDMapFail(t *testing.T) {
	assert := assert.New(t)

	p := getDefaultPreference()
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectQuery("SELECT id, userId, email FROM preference").
		WithArgs(p.UserID.String()).
		WillReturnRows(getPreferenceRows().AddRow(p.ID.String(), p.UserID.String(), "INVALID"))

	_, err := c.GetByUserID(p.UserID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestPreferenceGetByIDQueryFail(t *testing.T) {
	assert := assert.New(t)

	p := getDefaultPreference()
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectQuery("SELECT id, userId, email FROM preference").
		WithArgs(p.UserID.String()).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.GetByUserID(p.UserID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestPreferenceGetAllSuccess(t *testing.T) {
	assert := assert.New(t)

	p := getDefaultPreference()
	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectQuery("SELECT id, userId, email FROM preference").
		WithArgs(offset, limit).
		WillReturnRows(getPreferenceRows().
			AddRow(p.ID.String(), p.UserID.String(), string(p.Email)).
			AddRow(p.ID.String(), p.UserID.String(), string(p.Email)))

	res, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(2, len(res))
}

func TestPreferenceGetAllMapFail(t *testing.T) {
	assert := assert.New(t)

	p := getDefaultPreference()
	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectQuery("SELECT id, userId, email FROM preference").
		WithArgs(offset, limit).
		WillReturnRows(getPreferenceRows().
			AddRow(p.ID.String(), p.UserID.String(), "INACTIVE").
			AddRow(p.ID.String(), p.UserID.String(), string(p.Email)))

	_, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestPreferenceGetAllFail(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectQuery("SELECT id, userId, email FROM preference").
		WithArgs(offset, limit).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestPreferenceInsertSuccess(t *testing.T) {
	assert := assert.New(t)

	p := getDefaultPreference()
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectPrepare("INSERT INTO preference").
		ExpectExec().
		WithArgs(p.ID, p.UserID, string(p.Email)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.Insert(p)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestPreferenceInsertFail(t *testing.T) {
	assert := assert.New(t)

	p := getDefaultPreference()
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectPrepare("INSERT INTO preference").
		ExpectExec().
		WithArgs(p.ID, p.UserID, string(p.Email)).
		WillReturnError(fmt.Errorf("some error"))

	err := c.Insert(p)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestPreferenceUpdateSuccess(t *testing.T) {
	assert := assert.New(t)

	p := getDefaultPreference()
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectPrepare("UPDATE preference").
		ExpectExec().
		WithArgs(string(p.Email), p.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.Update(p)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestPreferenceUpdateFail(t *testing.T) {
	assert := assert.New(t)

	p := getDefaultPreference()
	s, mock, _ := sqlmock.New()
	c := getMockPreference(s)

	mock.ExpectPrepare("UPDATE preference").
		ExpectExec().
		WithArgs(string(p.Email), p.ID).
		WillReturnError(fmt.Errorf("some error"))

	err := c.Update(p)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func getDefaultPreference() *models.Preference {
	return models.NewPreference(uuid.NewUUID())
}

func getPreferenceRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "userId", "email"})
}

func getMockPreference(s *sql.DB) PreferenceI {
	return NewPreference(&gateways.MySQL{DB: s})
}
