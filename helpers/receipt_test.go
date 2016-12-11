package helpers

import (
	"database/sql"
	"fmt"
	"testing"

	mocks "github.com/ghmeier/bloodlines/_mocks/gateways"
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

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId, userId FROM receipt").
		WithArgs(receipt.ID.String()).
		WillReturnRows(getReceiptRows().
			AddRow(receipt.ID.String(), receipt.Created, "{}", string(receipt.SendState), receipt.ContentID.String(), receipt.UserID.String()))

	res, err := c.GetByID(receipt.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(receipt.ID, res.ID)
	assert.Equal(receipt.ContentID, res.ContentID)
	assert.Equal(receipt.UserID, res.UserID)
	assert.Equal(0, len(res.Values))
	assert.EqualValues(receipt.SendState, res.SendState)
	assert.Equal(receipt.Created, res.Created)
}

func TestReceiptGetByIDQueryFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId, userId FROM receipt").
		WithArgs(receipt.ID.String()).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.GetByID(receipt.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptGetByIDMapValueFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId, userId FROM receipt").
		WithArgs(receipt.ID.String()).
		WillReturnRows(getReceiptRows().AddRow(receipt.ID.String(), receipt.Created, "", "INVALID", receipt.ContentID.String(), receipt.UserID.String()))

	_, err := c.GetByID(receipt.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptGetByIDMapFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId, userId FROM receipt").
		WithArgs(receipt.ID.String()).
		WillReturnRows(getReceiptRows().AddRow(receipt.ID.String(), receipt.Created, "{}", "INVALID", receipt.ContentID.String(), receipt.UserID.String()))

	_, err := c.GetByID(receipt.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptGetReceiptsSuccess(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId, userId FROM receipt").
		WithArgs(offset, limit).
		WillReturnRows(getReceiptRows().
			AddRow(receipt.ID.String(), receipt.Created, "{}", string(receipt.SendState), receipt.ContentID.String(), receipt.UserID.String()).
			AddRow(receipt.ID.String(), receipt.Created, "{}", string(receipt.SendState), receipt.ContentID.String(), receipt.UserID.String()))

	res, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(2, len(res))
}

func TestReceiptGetReceiptsFail(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId, userId FROM receipt").
		WithArgs(offset, limit).
		WillReturnError(fmt.Errorf("some error"))

	_, err := c.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptGetReceiptsMapFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectQuery("SELECT id, ts, vals, sendState, contentId, userId FROM receipt").
		WithArgs(offset, limit).
		WillReturnRows(getReceiptRows().
			AddRow(receipt.ID.String(), receipt.Created, "{}", "INVALID", receipt.ContentID.String(), receipt.UserID.String()))

	_, err := c.GetAll(offset, limit)

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
		WithArgs(receipt.ID, receipt.Created, "{}", string(receipt.SendState), receipt.ContentID, receipt.UserID).
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
		WithArgs(receipt.ID, receipt.Created, "{}", string(receipt.SendState), receipt.ContentID, receipt.UserID).
		WillReturnError(fmt.Errorf("some error"))

	err := c.Insert(receipt)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestReceiptSetStatusSuccess(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectPrepare("UPDATE receipt").
		ExpectExec().
		WithArgs(string(receipt.SendState), receipt.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := c.SetStatus(receipt.ID, receipt.SendState)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestReceiptSetStatusFail(t *testing.T) {
	assert := assert.New(t)

	receipt := getDefaultReceipt()
	s, mock, _ := sqlmock.New()
	c := getMockReceipt(s)

	mock.ExpectPrepare("UPDATE receipt").
		ExpectExec().
		WithArgs(string(receipt.SendState), receipt.ID).
		WillReturnError(fmt.Errorf("some error"))

	err := c.SetStatus(receipt.ID, receipt.SendState)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func getDefaultReceipt() *models.Receipt {
	return models.NewReceipt(make(map[string]string), uuid.NewUUID(), uuid.NewUUID())
}

func getReceiptRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "ts", "vals", "sendState", "contentId", "userId"})
}

func getMockReceipt(s *sql.DB) ReceiptI {
	rabbitMock := &mocks.RabbitI{}
	r := NewReceipt(&gateways.MySQL{DB: s}, &mocks.SendgridI{}, &mocks.TownCenterI{}, rabbitMock)
	rabbitMock.On("Consume").Return(nil, nil)

	return r
}
