package handlers

import (
	"testing"

	mocks "github.com/ghmeier/bloodlines/_mocks/gateways"
	tmocks "github.com/jakelong95/TownCenter/_mocks"

	"github.com/stretchr/testify/assert"
	"gopkg.in/alexcesaro/statsd.v2"
)

func TestNewTrigger(t *testing.T) {
	assert := assert.New(t)
	tr := NewTrigger(GetTestContext())

	assert.NotNil(tr)
}

func GetTestContext() *GatewayContext {
	stats, _ := statsd.New()
	return &GatewayContext{
		Sql:        new(mocks.SQL),
		Sendgrid:   &mocks.SendgridI{},
		TownCenter: &tmocks.TownCenterI{},
		Rabbit:     &mocks.RabbitI{},
		Stats:      stats,
	}
}
