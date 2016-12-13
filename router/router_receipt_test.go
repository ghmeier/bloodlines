package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ghmeier/bloodlines/models"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
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

func getDefaultReceipt() *models.Receipt {
	return &models.Receipt{
		ID:        uuid.NewUUID(),
		Values:    make(map[string]string),
		ContentID: uuid.NewUUID(),
		UserID:    uuid.NewUUID(),
		SendState: models.ACTIVE,
	}
}
