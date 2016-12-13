package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ghmeier/bloodlines/models"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/gin-gonic/gin.v1"
)

func TestContentViewAllSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockContent()
	cMock.On("GetAll", 0, 20).Return(make([]*models.Content, 0), nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/content", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestContentViewAllParams(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockContent()
	cMock.On("GetAll", 20, 40).Return(make([]*models.Content, 0), nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/content?offset=20&limit=40", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestContentViewAllFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockContent()
	cMock.On("GetAll", 0, 20).Return(make([]*models.Content, 0), fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/content", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestContentNewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockContent()
	cMock.On("Insert", mock.AnythingOfType("*models.Content")).Return(nil)

	c := getContentString(models.NewContent(models.EMAIL, "test", "test", make([]string, 0)))
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/content", c)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestContentNewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockContent()
	cMock.On("Insert", mock.AnythingOfType("*models.Content")).Return(fmt.Errorf("some error"))

	c := getContentString(models.NewContent(models.EMAIL, "test", "test", make([]string, 0)))
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/content", c)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestContentNewInvalid(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockContent()
	cMock.On("Insert", mock.AnythingOfType("*models.Content")).Return(nil)

	c := bytes.NewReader([]byte("{\"id\": \"INVALID\"}"))
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/content", c)
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
}

func TestContentViewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	b, cMock := mockContent()
	cMock.On("GetByID", id.String()).Return(&models.Content{}, nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/content/"+id.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestContentViewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	b, cMock := mockContent()
	cMock.On("GetByID", id.String()).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/content/"+id.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestContentUpdateSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	content := models.NewContent(models.EMAIL, "test", "test", make([]string, 0))

	b, cMock := mockContent()
	cMock.On("Update", content).Return(nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(
		"PUT",
		"/api/content/"+content.ID.String(),
		getContentString(content),
	)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestContentUpdateFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	content := models.NewContent(models.EMAIL, "test", "test", make([]string, 0))

	b, cMock := mockContent()
	cMock.On("Update", content).Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(
		"PUT",
		"/api/content/"+content.ID.String(),
		getContentString(content),
	)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestContentUpdateInvalid(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, cMock := mockContent()
	cMock.On("Update", mock.AnythingOfType("*models.Content")).
		Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(
		"PUT",
		"/api/content/INVALID",
		bytes.NewReader([]byte("{\"id\": \"INVALID\"}")),
	)
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
}

func TestContentDeactivateSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	b, cMock := mockContent()
	cMock.
		On("SetStatus", id.String(), models.ContentStatus(models.INACTIVE)).
		Return(nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/content/"+id.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestContentDeactivateFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	b, cMock := mockContent()
	cMock.
		On("SetStatus", id.String(), models.ContentStatus(models.INACTIVE)).
		Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/content/"+id.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func getContentString(m *models.Content) io.Reader {
	s, _ := json.Marshal(m)
	return bytes.NewReader(s)
}
