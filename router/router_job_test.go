package router

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"time"
	// "io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ghmeier/bloodlines/models"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
	"gopkg.in/gin-gonic/gin.v1"
)

func TestJobViewAllSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockJob()
	cMock.On("GetAll", 0, 20).Return(make([]*models.Job, 0), nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/job", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestJobViewAllParams(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockJob()
	cMock.On("GetAll", 20, 40).Return(make([]*models.Job, 0), nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/job?offset=20&limit=40", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestJobViewAllFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockJob()
	cMock.On("GetAll", 0, 20).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/job", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestJobViewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	job := getDefaultJob()
	b, cMock := mockJob()
	cMock.On("GetByID", job.ID.String()).Return(job, nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/job/"+job.ID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestJobViewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	job := getDefaultJob()
	b, cMock := mockJob()
	cMock.On("GetByID", job.ID.String()).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/job/"+job.ID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestJobStopSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	job := getDefaultJob()
	b, cMock := mockJob()
	cMock.On("SetStatus", job.ID, models.Status(models.FAILURE)).Return(nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/job/"+job.ID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestJobStopFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	job := getDefaultJob()
	b, cMock := mockJob()
	cMock.On("SetStatus", job.ID, models.Status(models.FAILURE)).Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/job/"+job.ID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func getDefaultJob() *models.Job {
	return &models.Job{
		ID:         uuid.NewUUID(),
		SendTime:   time.Now(),
		SendStatus: models.READY,
		Receipts:   []uuid.UUID{uuid.NewUUID()},
	}
}
