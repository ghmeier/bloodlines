package handlers

import (
	"testing"

	mocks "github.com/ghmeier/bloodlines/_mocks/gateways"

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
		TownCenter: &mocks.TownCenterI{},
		Rabbit:     &mocks.RabbitI{},
		Stats:      stats,
	}
}
