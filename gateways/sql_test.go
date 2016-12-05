package gateways

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSQLModifySuccess(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	sql := &MySQL{DB: s}

	mock.ExpectPrepare("INSERT INTO test").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := sql.Modify("INSERT INTO test SET id=?", 1)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestSQLModifyFail(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	sql := &MySQL{DB: s}

	mock.ExpectPrepare("INSERT INTO test").
		ExpectExec().
		WithArgs(1).
		WillReturnError(fmt.Errorf("some error"))

	err := sql.Modify("INSERT INTO test SET id=?", 1)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)

}

func TestSQLModifyPrepareFail(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	sql := &MySQL{DB: s}

	mock.ExpectPrepare("INSERT INTO test").
		WillReturnError(fmt.Errorf("some error"))

	err := sql.Modify("INSERT INTO test SET id=?", 1)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestSQLSelectSuccess(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	sql := &MySQL{DB: s}

	mock.ExpectQuery("SELECT name FROM testing").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	_, err := sql.Select("SELECT name FROM testing")

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestSQLSelectMultiArgsSuccess(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	sql := &MySQL{DB: s}

	mock.ExpectQuery("SELECT name FROM testing").
		WithArgs(2, 2).
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	_, err := sql.Select("SELECT name FROM testing", 2, 2)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestSQLSelectFail(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	sql := &MySQL{DB: s}

	mock.ExpectQuery("SELECT id FROM testing").
		WithArgs(2, 2).
		WillReturnError(fmt.Errorf("some error"))

	_, err := sql.Select("SELECT id FROM testing", 2, 2)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestSQLDestroy(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	sql := &MySQL{DB: s}

	mock.ExpectClose()

	sql.Destroy()

	assert.Equal(mock.ExpectationsWereMet(), nil)
}
