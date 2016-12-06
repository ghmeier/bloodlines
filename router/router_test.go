package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ghmeier/bloodlines/_mocks"
	"github.com/ghmeier/bloodlines/handlers"
	"github.com/ghmeier/bloodlines/models"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
	"gopkg.in/gin-gonic/gin.v1"
)

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
