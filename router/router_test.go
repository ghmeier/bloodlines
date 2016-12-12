package router

import (
	"testing"

	mocks "github.com/ghmeier/bloodlines/_mocks"
	mockg "github.com/ghmeier/bloodlines/_mocks/gateways"
	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/handlers"

	"github.com/stretchr/testify/assert"
)

func TestNewSuccess(t *testing.T) {
	assert := assert.New(t)

	r, err := New(&config.Root{SQL: config.MySQL{}})

	assert.Error(err)
	assert.Nil(r)
}

func getMockBloodlines() *Bloodlines {
	sql := new(mockg.SQL)
	towncenter := new(mockg.TownCenterI)
	sendgrid := new(mockg.SendgridI)
	rabbit := new(mockg.RabbitI)
	rabbit.On("Consume").Return(nil, nil)
	return &Bloodlines{
		content:    handlers.NewContent(sql),
		receipt:    handlers.NewReceipt(sql, sendgrid, towncenter, rabbit),
		job:        handlers.NewJob(sql),
		trigger:    handlers.NewTrigger(sql, sendgrid, towncenter, rabbit),
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
