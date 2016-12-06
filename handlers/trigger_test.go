package handlers

import (
	"testing"

	"github.com/ghmeier/bloodlines/_mocks"

	"github.com/stretchr/testify/assert"
)

func TestNewTrigger(t *testing.T) {
	assert := assert.New(t)
	tr := NewTrigger(new(mocks.SQL))

	assert.NotNil(tr)
}
