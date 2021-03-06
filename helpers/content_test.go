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

func TestGetSuccess(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectQuery("SELECT id, contentType, text, parameters, status, subject FROM content").
		WithArgs(id.String()).
		WillReturnRows(getMockRows().AddRow(id.String(), "EMAIL", "HelloWorld", "", "ACTIVE", "test"))

	content, err := c.Get(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(content.ID, id)
	assert.EqualValues(content.Type, models.EMAIL)
	assert.Equal(content.Text, "HelloWorld")
	assert.Equal(content.Subject, "test")
	assert.Equal(len(content.Params), 0)
	assert.EqualValues(content.Status, models.ACTIVE)
}

func TestGetNoResults(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectQuery("SELECT id, contentType, text, parameters, status, subject FROM content").
		WithArgs(id.String()).
		WillReturnRows(getMockRows())

	content, err := c.Get(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Nil(content)
}

func TestGetQueryFail(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectQuery("SELECT id, contentType, text, parameters, status, subject FROM content").
		WithArgs(id.String()).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.Get(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestGetMapFail(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectQuery("SELECT id, contentType, text, parameters, status, subject FROM content").
		WithArgs(id.String()).
		WillReturnRows(getMockRows().AddRow(id.String(), "INVALID", "HelloWorld", "", "ACTIVE", "test"))

	_, err := c.Get(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestGetMapStatusFail(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectQuery("SELECT id, contentType, text, parameters, status, subject FROM content").
		WithArgs(id.String()).
		WillReturnRows(getMockRows().AddRow(id.String(), "ACTIVE", "HelloWorld", "", "INVALID", "test"))

	_, err := c.Get(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestGetAllSuccess(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectQuery("SELECT id, contentType, text, parameters, status, subject FROM content").
		WithArgs(models.ACTIVE, offset, limit).
		WillReturnRows(getMockRows().
			AddRow(uuid.New(), "EMAIL", "HelloWorld", "", "INACTIVE", "test").
			AddRow(uuid.New(), "EMAIL", "HelloWorld", "", "ACTIVE", "test"))

	contents, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(2, len(contents))
}

func TestGetAllQueryFail(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectQuery("SELECT id, contentType, text, parameters, status, subject FROM content").
		WithArgs(models.ACTIVE, offset, limit).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestGetAllMapFail(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectQuery("SELECT id, contentType, text, parameters, status, subject FROM content").
		WithArgs(models.ACTIVE, offset, limit).
		WillReturnRows(getMockRows().
			AddRow(uuid.New(), "INVALID", "HelloWorld", "", "ACTIVE", "test").
			AddRow(uuid.New(), "INVALID", "HelloWorld", "", "ACTIVE", "test"))

	_, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestInsertSuccess(t *testing.T) {
	assert := assert.New(t)

	content := getDefaultContent()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectPrepare("INSERT INTO content").
		ExpectExec().
		WithArgs(content.ID.String(), string(content.Type), content.Text, "", string(models.ACTIVE), "test").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.Insert(content)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestInsertFail(t *testing.T) {
	assert := assert.New(t)

	content := getDefaultContent()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectPrepare("INSERT INTO content").
		ExpectExec().
		WithArgs(content.ID, string(content.Type), content.Text, "", string(models.ACTIVE), "test").
		WillReturnError(fmt.Errorf("some error"))

	err := c.Insert(content)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestUpdateSuccess(t *testing.T) {
	assert := assert.New(t)

	content := getDefaultContent()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectPrepare("UPDATE content").
		ExpectExec().
		WithArgs(string(content.Type), content.Text, "", string(models.ACTIVE), "test", content.ID.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.Update(content)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestUpdateFail(t *testing.T) {
	assert := assert.New(t)

	content := getDefaultContent()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectPrepare("UPDATE content").
		ExpectExec().
		WithArgs(string(content.Type), content.Text, "", string(models.ACTIVE), "test", content.ID.String()).
		WillReturnError(fmt.Errorf("some error"))

	err := c.Update(content)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestSetStatusSuccess(t *testing.T) {
	assert := assert.New(t)

	content := getDefaultContent()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectPrepare("UPDATE content").
		ExpectExec().
		WithArgs(string(content.Status), content.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.SetStatus(content.ID.String(), content.Status)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestSetStatusFail(t *testing.T) {
	assert := assert.New(t)

	content := getDefaultContent()
	s, mock, _ := sqlmock.New()
	c := getMockContent(s)

	mock.ExpectPrepare("UPDATE content").
		ExpectExec().
		WithArgs(string(content.Status), content.ID).
		WillReturnError(fmt.Errorf("some error"))

	err := c.SetStatus(content.ID.String(), content.Status)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func getDefaultContent() *models.Content {
	return models.NewContent(models.EMAIL, "Hello", "test", make([]string, 0))
}

func getMockRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "contentType", "text", "parameters", "status", "subject"})
}

func getMockContent(s *sql.DB) *Content {
	return NewContent(&gateways.MySQL{DB: s})
}
