package router

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ghmeier/bloodlines/models"

	//"github.com/pborman/uuid"
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
