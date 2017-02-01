package router

import (
	"testing"

	mocks "github.com/ghmeier/bloodlines/_mocks"
	mockg "github.com/ghmeier/bloodlines/_mocks/gateways"
	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/handlers"
	tmocks "github.com/jakelong95/TownCenter/_mocks"

	"github.com/stretchr/testify/assert"
	"gopkg.in/alexcesaro/statsd.v2"
)

func TestNewSuccess(t *testing.T) {
	assert := assert.New(t)

	r, err := New(&config.Root{SQL: config.MySQL{}})

	assert.NoError(err)
	assert.NotNil(r)
}

func getMockBloodlines() *Bloodlines {
	sql := new(mockg.SQL)
	towncenter := new(tmocks.TownCenterI)
	sendgrid := new(mockg.SendgridI)
	rabbit := new(mockg.RabbitI)
	rabbit.On("Consume").Return(nil, nil)
	stats, _ := statsd.New()
	ctx := &handlers.GatewayContext{
		Sql:        sql,
		Sendgrid:   sendgrid,
		TownCenter: towncenter,
		Rabbit:     rabbit,
		Stats:      stats,
	}
	return &Bloodlines{
		content:    handlers.NewContent(ctx),
		receipt:    handlers.NewReceipt(ctx),
		job:        handlers.NewJob(ctx),
		trigger:    handlers.NewTrigger(ctx),
		preference: handlers.NewPreference(ctx),
	}
}

func mockTrigger() (*Bloodlines, *mocks.TriggerI) {
	b := getMockBloodlines()
	mock := new(mocks.TriggerI)
	b.trigger = &handlers.Trigger{Helper: mock, BaseHandler: &handlers.BaseHandler{Stats: nil}}
	InitRouter(b)

	return b, mock
}

func mockContent() (*Bloodlines, *mocks.ContentI) {
	b := getMockBloodlines()
	mock := new(mocks.ContentI)
	b.content = &handlers.Content{Helper: mock, BaseHandler: &handlers.BaseHandler{Stats: nil}}
	InitRouter(b)

	return b, mock
}

func mockJob() (*Bloodlines, *mocks.JobI) {
	b := getMockBloodlines()
	mock := new(mocks.JobI)
	b.job = &handlers.Job{Helper: mock, BaseHandler: &handlers.BaseHandler{Stats: nil}}
	InitRouter(b)

	return b, mock
}

func mockReceipt() (*Bloodlines, *mocks.ReceiptI, *mocks.ContentI) {
	b := getMockBloodlines()
	mock := new(mocks.ReceiptI)
	cmock := new(mocks.ContentI)
	b.receipt = &handlers.Receipt{Helper: mock, CHelper: cmock, BaseHandler: &handlers.BaseHandler{Stats: nil}}
	mock.On("Consume").Return(nil)
	InitRouter(b)

	return b, mock, cmock
}

func mockPreference() (*Bloodlines, *mocks.PreferenceI) {
	b := getMockBloodlines()
	mock := new(mocks.PreferenceI)
	b.preference = &handlers.Preference{Helper: mock, BaseHandler: &handlers.BaseHandler{Stats: nil}}
	InitRouter(b)

	return b, mock
}
