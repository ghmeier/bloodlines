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

func TestReceiptGetByIDSuccess(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId FROM receipt").
		WithArgs(receipt.ID.String()).
		WillReturnRows(getReceiptRows().AddRow(receipt.ID.String(), receipt.Created, "{}", string(receipt.SendState), receipt.ContentID.String()))

	res, err := c.GetReceiptByID(receipt.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(receipt.ID, res.ID)
	assert.Equal(receipt.ContentID, res.ContentID)
	assert.Equal(0, len(res.Values))
	assert.EqualValues(receipt.SendState, res.SendState)
	assert.Equal(receipt.Created, res.Created)
}

func TestReceiptGetByIDQueryFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId FROM receipt").
		WithArgs(receipt.ID.String()).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.GetReceiptByID(receipt.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptGetByIDMapFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId FROM receipt").
		WithArgs(receipt.ID.String()).
		WillReturnRows(getReceiptRows().AddRow(receipt.ID.String(), receipt.Created, "{}", "INVALID", receipt.ContentID.String()))

	_, err := c.GetReceiptByID(receipt.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptGetReceiptsSuccess(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId FROM receipt").
		WithArgs(offset, limit).
		WillReturnRows(getReceiptRows().
			AddRow(receipt.ID.String(), receipt.Created, "{}", string(receipt.SendState), receipt.ContentID.String()).
			AddRow(receipt.ID.String(), receipt.Created, "{}", string(receipt.SendState), receipt.ContentID.String()))

	res, err := c.GetReceipts(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(2, len(res))
}

func TestReceiptGetReceiptsFail(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId FROM receipt").
		WithArgs(offset, limit).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.GetReceipts(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptGetReceiptsMapFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId FROM receipt").
		WithArgs(offset, limit).
		WillReturnRows(getReceiptRows().
			AddRow(receipt.ID.String(), receipt.Created, "{}", "INVALID", receipt.ContentID.String()))

	_, err := c.GetReceipts(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptInsertSuccess(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectPrepare("INSERT INTO receipt").
		ExpectExec().
		WithArgs(receipt.ID, receipt.Created, "{}", string(receipt.SendState), receipt.ContentID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.Insert(receipt)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestReceiptInsertFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectPrepare("INSERT INTO receipt").
		ExpectExec().
		WithArgs(receipt.ID, receipt.Created, "{}", string(receipt.SendState), receipt.ContentID).
		WillReturnError(fmt.Errorf("some error"))

	err := c.Insert(receipt)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptSetSendStateSuccess(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectPrepare("UPDATE receipt").
		ExpectExec().
		WithArgs(string(receipt.SendState), receipt.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.SetSendState(receipt.ID, receipt.SendState)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestReceiptSetSendStateFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectPrepare("UPDATE receipt").
		ExpectExec().
		WithArgs(string(receipt.SendState), receipt.ID).
		WillReturnError(fmt.Errorf("some error"))

	err := c.SetSendState(receipt.ID, receipt.SendState)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func getDefaultReceipt() *models.Receipt {
	return models.NewReceipt(make(map[string]string), uuid.NewUUID())
}

func getReceiptRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "ts", "vals", "sendState", "contentId"})
}

func getMockReceipt(s *sql.DB) *Receipt {
	return NewReceipt(&gateways.MySQL{DB: s})
}
