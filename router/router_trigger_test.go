package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "github.com/ghmeier/bloodlines/_mocks"
	"github.com/ghmeier/bloodlines/handlers"
	"github.com/ghmeier/bloodlines/models"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		Values:    make(map[string]string),
	}

	b, tMock := mockTrigger()
	tMock.On("Get", "test_key").Return(nil, nil)
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
		Values:    make(map[string]string),
	}

	b, tMock := mockTrigger()
	tMock.On("Get", "test_key").Return(nil, nil)
	tMock.On("Insert", mock.AnythingOfType("*models.Trigger")).Return(fmt.Errorf("some error"))

	s, _ := json.Marshal(trigger)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestTriggerNewDuplicate(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	trigger := models.Trigger{
		ContentID: uuid.NewUUID(),
		Key:       "test_key",
		Values:    make(map[string]string),
	}

	b, tMock := mockTrigger()
	tMock.On("Get", "test_key").Return(&models.Trigger{}, nil)

	s, _ := json.Marshal(trigger)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
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

func TestTriggerViewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, mock := mockTrigger()
	mock.On("Get", "test_key").Return(&models.Trigger{}, nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/trigger/test_key", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestTriggerViewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, mock := mockTrigger()
	mock.On("Get", "test_key").Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/trigger/test_key", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestTriggerUpdateSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	trigger := models.Trigger{
		ID:        uuid.NewUUID(),
		ContentID: uuid.NewUUID(),
		Values:    make(map[string]string),
		Key:       "test_key",
	}

	b, tMock := mockTrigger()
	tMock.On("Update", trigger.Key, trigger.ContentID, trigger.Values).Return(nil)

	s, _ := json.Marshal(trigger)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/trigger/test_key", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestTriggerUpdateInvalid(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b, _ := mockTrigger()

	s := []byte("{\"contentId\": \"invalid-id\"}")
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/trigger/test_key", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
}

func TestTriggerUpdateFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	trigger := models.Trigger{
		ID:        uuid.NewUUID(),
		ContentID: uuid.NewUUID(),
		Key:       "test_key",
		Values:    make(map[string]string),
	}

	b, tMock := mockTrigger()
	tMock.On("Update", trigger.Key, trigger.ContentID, trigger.Values).Return(fmt.Errorf("some error"))

	s, _ := json.Marshal(trigger)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/trigger/test_key", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestTriggerDeleteSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	key := "test_key"

	b, tMock := mockTrigger()
	tMock.On("Delete", key).Return(nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/trigger/test_key", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestTriggerDeleteFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	key := "test_key"

	b, tMock := mockTrigger()
	tMock.On("Delete", key).Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/trigger/test_key", nil)
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestTriggerActivateSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	mockTrigger := &mocks.TriggerI{}
	mockContent := &mocks.ContentI{}
	mockReceipt := &mocks.ReceiptI{}
	b := getMockBloodlines()
	b.trigger = &handlers.Trigger{
		Trigger:     mockTrigger,
		Receipt:     mockReceipt,
		Content:     mockContent,
		BaseHandler: &handlers.BaseHandler{Stats: nil},
	}
	InitRouter(b)

	key := "test_key"
	values := make(map[string]string)
	values["first_name"] = "Wololo"

	trigger := &models.Trigger{
		Key:       key,
		ID:        uuid.NewUUID(),
		ContentID: uuid.NewUUID(),
		Values:    values,
	}

	content := &models.Content{
		ID:      trigger.ContentID,
		Type:    models.EMAIL,
		Text:    "$first_name$ $last_name$",
		Params:  []string{"first_name", "last_name"},
		Status:  models.ACTIVE,
		Subject: "Welcome",
	}

	values = make(map[string]string)
	values["last_name"] = "ololo"
	receipt := &models.Receipt{
		Values: values,
		UserID: uuid.NewUUID(),
	}
	s, _ := json.Marshal(receipt)

	mockTrigger.On("Get", key).Return(trigger, nil)
	mockContent.On("Get", trigger.ContentID.String()).Return(content, nil)
	mockReceipt.On("Insert", mock.AnythingOfType("*models.Receipt")).Return(nil)
	mockReceipt.On("Send", mock.AnythingOfType("*models.SendRequest")).Return(nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger/test_key/activate", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(200, w.Code)
}

func TestTriggerActivateInsertFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	mockTrigger := &mocks.TriggerI{}
	mockContent := &mocks.ContentI{}
	mockReceipt := &mocks.ReceiptI{}
	b := getMockBloodlines()
	b.trigger = &handlers.Trigger{
		Trigger:     mockTrigger,
		Receipt:     mockReceipt,
		Content:     mockContent,
		BaseHandler: &handlers.BaseHandler{Stats: nil},
	}
	InitRouter(b)

	key := "test_key"
	values := make(map[string]string)
	values["first_name"] = "Wololo"
	values["last_name"] = "lololo"

	trigger := &models.Trigger{
		Key:       key,
		ID:        uuid.NewUUID(),
		ContentID: uuid.NewUUID(),
		Values:    values,
	}

	content := &models.Content{
		ID:      trigger.ContentID,
		Type:    models.EMAIL,
		Text:    "$first_name$ $last_name$",
		Params:  []string{"first_name", "last_name"},
		Status:  models.ACTIVE,
		Subject: "Welcome",
	}

	values = make(map[string]string)
	receipt := &models.Receipt{
		Values: values,
		UserID: uuid.NewUUID(),
	}
	s, _ := json.Marshal(receipt)

	mockTrigger.On("Get", key).Return(trigger, nil)
	mockContent.On("Get", trigger.ContentID.String()).Return(content, nil)
	mockReceipt.On("Insert", mock.AnythingOfType("*models.Receipt")).Return(fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger/test_key/activate", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestTriggerActivateValueMapError(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	mockTrigger := &mocks.TriggerI{}
	mockContent := &mocks.ContentI{}
	mockReceipt := &mocks.ReceiptI{}
	b := getMockBloodlines()
	b.trigger = &handlers.Trigger{
		Trigger:     mockTrigger,
		Receipt:     mockReceipt,
		Content:     mockContent,
		BaseHandler: &handlers.BaseHandler{Stats: nil},
	}
	InitRouter(b)

	key := "test_key"
	values := make(map[string]string)
	values["first_name"] = "Wololo"

	trigger := &models.Trigger{
		Key:       key,
		ID:        uuid.NewUUID(),
		ContentID: uuid.NewUUID(),
		Values:    values,
	}

	content := &models.Content{
		ID:      trigger.ContentID,
		Type:    models.EMAIL,
		Text:    "$first_name$ $last_name$",
		Params:  []string{"first_name", "last_name"},
		Status:  models.ACTIVE,
		Subject: "Welcome",
	}

	values = make(map[string]string)
	receipt := &models.Receipt{
		Values: values,
		UserID: uuid.NewUUID(),
	}
	s, _ := json.Marshal(receipt)

	mockTrigger.On("Get", key).Return(trigger, nil)
	mockContent.On("Get", trigger.ContentID.String()).Return(content, nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger/test_key/activate", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
}

func TestTriggerActivateTriggerFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	mockTrigger := &mocks.TriggerI{}
	b := getMockBloodlines()
	b.trigger = &handlers.Trigger{
		Trigger:     mockTrigger,
		BaseHandler: &handlers.BaseHandler{Stats: nil},
	}
	InitRouter(b)

	key := "test_key"

	values := make(map[string]string)
	values["last_name"] = "ololo"
	receipt := &models.Receipt{
		Values: values,
		UserID: uuid.NewUUID(),
	}
	s, _ := json.Marshal(receipt)

	mockTrigger.On("Get", key).Return(nil, fmt.Errorf("some error"))
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger/test_key/activate", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestTriggerActivateTriggerNil(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	mockTrigger := &mocks.TriggerI{}
	b := getMockBloodlines()
	b.trigger = &handlers.Trigger{
		Trigger:     mockTrigger,
		BaseHandler: &handlers.BaseHandler{Stats: nil},
	}
	InitRouter(b)

	key := "test_key"

	values := make(map[string]string)
	values["last_name"] = "ololo"
	receipt := &models.Receipt{
		Values: values,
		UserID: uuid.NewUUID(),
	}
	s, _ := json.Marshal(receipt)

	mockTrigger.On("Get", key).Return(nil, nil)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger/test_key/activate", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
}

func TestTriggerActivateContentFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	mockTrigger := &mocks.TriggerI{}
	mockContent := &mocks.ContentI{}
	b := getMockBloodlines()
	b.trigger = &handlers.Trigger{
		Trigger:     mockTrigger,
		Content:     mockContent,
		BaseHandler: &handlers.BaseHandler{Stats: nil},
	}
	InitRouter(b)

	key := "test_key"
	values := make(map[string]string)
	values["first_name"] = "Wololo"

	trigger := &models.Trigger{
		Key:       key,
		ID:        uuid.NewUUID(),
		ContentID: uuid.NewUUID(),
		Values:    values,
	}

	values = make(map[string]string)
	values["last_name"] = "ololo"
	receipt := &models.Receipt{
		Values: values,
		UserID: uuid.NewUUID(),
	}
	s, _ := json.Marshal(receipt)

	mockTrigger.On("Get", key).Return(trigger, nil)
	mockContent.On("Get", trigger.ContentID.String()).Return(nil, fmt.Errorf("some error"))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger/test_key/activate", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(500, w.Code)
}

func TestTriggerActivateInvalid(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	b := getMockBloodlines()
	b.trigger = &handlers.Trigger{
		BaseHandler: &handlers.BaseHandler{Stats: nil},
	}
	InitRouter(b)

	s := []byte("{\"userId\":\"invalid-id\"}")

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/trigger/test_key/activate", bytes.NewReader(s))
	b.router.ServeHTTP(w, r)

	assert.Equal(400, w.Code)
}
