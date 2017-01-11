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

	assert.NoError(err)
	assert.NotNil(r)
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

func mockJob() (*Bloodlines, *mocks.JobI) {
	b := getMockBloodlines()
	mock := new(mocks.JobI)
	b.job = &handlers.Job{Helper: mock}
	InitRouter(b)

	return b, mock
}

func mockReceipt() (*Bloodlines, *mocks.ReceiptI, *mocks.ContentI) {
	b := getMockBloodlines()
	mock := new(mocks.ReceiptI)
	cmock := new(mocks.ContentI)
	b.receipt = &handlers.Receipt{Helper: mock, CHelper: cmock}
	mock.On("Consume").Return(nil)
	InitRouter(b)

	return b, mock, cmock
}

func mockPreference() (*Bloodlines, *mocks.PreferenceI) {
	b := getMockBloodlines()
	mock := new(mocks.PreferenceI)
	b.preference = &handlers.Preference{Helper: mock}
	InitRouter(b)

	return b, mock
}
