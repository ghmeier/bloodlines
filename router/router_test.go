package router

import (
	"testing"

	"github.com/ghmeier/bloodlines/_mocks"
	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/handlers"

	"github.com/stretchr/testify/assert"
)

func TestNewSuccess(t *testing.T) {
	assert := assert.New(t)

	r, err := New(&config.Root{SQL: config.MySQL{}})

	assert.NoError(err)
	assert.NotNil(r)
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

func mockContent() (*Bloodlines, *mocks.ContentI) {
	b := getMockBloodlines()
	mock := new(mocks.ContentI)
	b.content = &handlers.Content{Helper: mock}
	InitRouter(b)

	return b, mock
}
