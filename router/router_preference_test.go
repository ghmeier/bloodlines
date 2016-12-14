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

func TestPreferenceViewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	pref := getDefaultPreference()

	b, mock := mockPreference()
	mock.On("GetByUserID", pref.UserID.String()).Return(pref, nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/preference/"+pref.UserID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestPreferenceViewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	pref := getDefaultPreference()

	b, mock := mockPreference()
	mock.On("GetByUserID", pref.UserID.String()).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/preference/"+pref.UserID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestPreferenceNewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	pref := getDefaultPreference()
	s, _ := json.Marshal(pref)

	b, mock := mockPreference()
	mock.On("Insert", mocks.AnythingOfType("*models.Preference")).Return(nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/preference", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestPreferenceNewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	pref := getDefaultPreference()
	s, _ := json.Marshal(pref)

	b, mock := mockPreference()
	mock.On("Insert", mocks.AnythingOfType("*models.Preference")).Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/preference", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestPreferenceNewMapFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	s := "{\"userId\":\"invalid-id\"}"

	b, _ := mockPreference()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/preference", bytes.NewReader([]byte(s)))
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
}

func TestPreferenceUpdateFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	pref := getDefaultPreference()
	s, _ := json.Marshal(pref)

	b, mock := mockPreference()
	mock.On("Update", pref).Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PATCH", "/api/preference/"+pref.UserID.String(), bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestPreferenceUpdateMapFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	s := "{\"userId\":\"invalid-id\"}"

	b, _ := mockPreference()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PATCH", "/api/preference/invalid-id", bytes.NewReader([]byte(s)))
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
}

func TestPreferenceUpdateSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	pref := getDefaultPreference()
	s, _ := json.Marshal(pref)

	b, mock := mockPreference()
	mock.On("Update", pref).Return(nil)
	mock.On("GetByUserID", pref.UserID.String()).Return(pref, nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PATCH", "/api/preference/"+pref.UserID.String(), bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestPreferenceDeactivateSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	pref := getDefaultPreference()

	b, mock := mockPreference()
	mock.On("GetByUserID", pref.UserID.String()).Return(pref, nil)
	pref.Email = models.UNSUBSCRIBED
	mock.On("Update", pref).Return(nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/preference/"+pref.UserID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestPreferenceDeactivateGetFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	pref := getDefaultPreference()

	b, mock := mockPreference()
	mock.On("GetByUserID", pref.UserID.String()).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/preference/"+pref.UserID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestPreferenceDeactivateUpdateFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	pref := getDefaultPreference()

	b, mock := mockPreference()
	mock.On("GetByUserID", pref.UserID.String()).Return(pref, nil)
	pref.Email = models.UNSUBSCRIBED
	mock.On("Update", pref).Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/preference/"+pref.UserID.String(), nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func getDefaultPreference() *models.Preference {
	return &models.Preference{
		ID:     uuid.NewUUID(),
		UserID: uuid.NewUUID(),
		Email:  models.SUBSCRIBED,
	}
}
