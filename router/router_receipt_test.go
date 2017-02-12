package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ghmeier/bloodlines/models"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	mocks "github.com/stretchr/testify/mock"
	"gopkg.in/gin-gonic/gin.v1"
)

func TestReceiptViewAllSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, mock, _ := mockReceipt()
	mock.On("GetAll", 0, 20).Return(make([]*models.Receipt, 0), nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/receipt", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestReceiptViewAllParamsSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, mock, _ := mockReceipt()
	mock.On("GetAll", 20, 40).Return(make([]*models.Receipt, 0), nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/receipt?offset=20&limit=40", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestReceiptViewAllParamsFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, mock, _ := mockReceipt()
	mock.On("GetAll", 0, 20).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/receipt", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestReceiptViewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	receipt := getDefaultReceipt()
	b, mock, _ := mockReceipt()
	mock.On("GetByID", receipt.ID.String()).Return(receipt, nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/receipt/"+receipt.ID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestReceiptViewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	receipt := getDefaultReceipt()
	b, mock, _ := mockReceipt()
	mock.On("GetByID", receipt.ID.String()).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/receipt/"+receipt.ID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestReceiptSendSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	receipt := getDefaultReceipt()
	receipt.Values["first_name"] = "test"
	content := &models.Content{
		ID:      receipt.ContentID,
		Type:    models.EMAIL,
		Text:    "Hello $first_name$",
		Params:  []string{"first_name"},
		Status:  models.ACTIVE,
		Subject: "Test",
	}
	s, _ := json.Marshal(receipt)

	b, mock, cmock := mockReceipt()
	mock.On("Insert", mocks.AnythingOfType("*models.Receipt")).Return(nil)
	mock.On("Send", mocks.AnythingOfType("*models.SendRequest")).Return(nil)
	cmock.On("Get", content.ID.String()).Return(content, nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/receipt/send", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestReceiptSendInsertFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	receipt := getDefaultReceipt()
	receipt.Values["first_name"] = "test"
	s, _ := json.Marshal(receipt)

	b, mock, _ := mockReceipt()
	mock.On("Insert", mocks.AnythingOfType("*models.Receipt")).Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/receipt/send", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestReceiptSendGetFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	receipt := getDefaultReceipt()
	receipt.Values["first_name"] = "test"
	s, _ := json.Marshal(receipt)

	b, mock, cmock := mockReceipt()
	mock.On("Insert", mocks.AnythingOfType("*models.Receipt")).Return(nil)
	cmock.On("Get", receipt.ContentID.String()).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/receipt/send", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestReceiptSendFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	receipt := getDefaultReceipt()
	receipt.Values["first_name"] = "test"
	content := &models.Content{
		ID:      receipt.ContentID,
		Type:    models.EMAIL,
		Text:    "Hello $first_name$",
		Params:  []string{"first_name"},
		Status:  models.ACTIVE,
		Subject: "Test",
	}
	s, _ := json.Marshal(receipt)

	b, mock, cmock := mockReceipt()
	mock.On("Insert", mocks.AnythingOfType("*models.Receipt")).Return(nil)
	mock.On("Send", mocks.AnythingOfType("*models.SendRequest")).Return(fmt.Errorf("some error"))
	cmock.On("Get", content.ID.String()).Return(content, nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/receipt/send", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func getDefaultReceipt() *models.Receipt {
	return &models.Receipt{
		ID:        uuid.NewUUID(),
		Values:    make(map[string]string),
		ContentID: uuid.NewUUID(),
		UserID:    uuid.NewUUID(),
		SendState: models.ACTIVE,
	}
}
