package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ghmeier/bloodlines/_mocks"
	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/handlers"
	"github.com/ghmeier/bloodlines/models"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/gin-gonic/gin.v1"
)

func TestNewSuccess(t *testing.T) {
	assert := assert.New(t)

	r, err := New(&config.Root{SQL: config.MySQL{}})

	assert.NoError(err)
	assert.NotNil(r)
}

func TestTriggerViewAllSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, mock := mockTrigger()
	mock.On("GetAll", 0, 20).Return(make([]*models.Trigger, 0), nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/trigger", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestTriggerViewAllParams(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, mock := mockTrigger()
	mock.On("GetAll", 20, 100).Return(make([]*models.Trigger, 0), nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/trigger?offset=20&limit=100", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestTriggerViewAllFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, mock := mockTrigger()
	mock.On("GetAll", 0, 20).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/trigger", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestTriggerNewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	trigger := models.Trigger{
		ContentID: uuid.NewUUID(),
		Key:       "test_key",
		Params:    make([]string, 0),
	}

	b, tMock := mockTrigger()
	tMock.On("Insert", mock.AnythingOfType("*models.Trigger")).Return(nil)

	s, _ := json.Marshal(trigger)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestTriggerNewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	trigger := models.Trigger{
		ContentID: uuid.NewUUID(),
		Key:       "test_key",
		Params:    make([]string, 0),
	}

	b, tMock := mockTrigger()
	tMock.On("Insert", mock.AnythingOfType("*models.Trigger")).Return(fmt.Errorf("some error"))

	s, _ := json.Marshal(trigger)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestTriggerNewInvalid(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	b, _ := mockTrigger()

	s := []byte("{\"contentId\": \"invalid-id\"}")
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
}

func getMockBloodlines() *Bloodlines {
	sql := new(mocks.SQL)
	return &Bloodlines{
		content:    handlers.NewContent(sql),
		receipt:    handlers.NewReceipt(sql),
		job:        handlers.NewJob(sql),
		trigger:    handlers.NewTrigger(sql),
		preference: handlers.NewPreference(sql),
	}
}

func mockTrigger() (*Bloodlines, *mocks.TriggerI) {
	b := getMockBloodlines()
	mock := new(mocks.TriggerI)
	b.trigger = &handlers.Trigger{Helper: mock}
	InitRouter(b)

	return b, mock
}
