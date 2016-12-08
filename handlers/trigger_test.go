package handlers

import (
	"testing"

	mocks "github.com/ghmeier/bloodlines/_mocks/gateways"

	"github.com/stretchr/testify/assert"
)

func TestNewTrigger(t *testing.T) {
	assert := assert.New(t)
	tr := NewTrigger(new(mocks.SQL))

	assert.NotNil(tr)
}
