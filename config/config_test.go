package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigInitFail(t *testing.T) {
	assert := assert.New(t)
	_, err := Init("not-found")

	assert.Error(err)
}
