package gateways

import (
	"testing"

	"github.com/ghmeier/bloodlines/config"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	assert := assert.New(t)

	s := getSendGrid()

	err := s.SendEmail("test", "test subject", "test text")

	assert.Error(err)
}

func getSendGrid() SendgridI {
	return NewSendgrid(config.Sendgrid{})
}
