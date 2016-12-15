package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigInitFail(t *testing.T) {
	assert := assert.New(t)
	_, err := Init("not-found")

	assert.Error(err)
}

func TestConfigInitSuccess(t *testing.T) {
	assert := assert.New(t)
	f, _ := os.Create("config-test.json")
	f.Write([]byte("{}"))
	_, err := Init("config-test.json")
	f.Close()
	assert.NoError(err)
}

func TestConfigInitMapFail(t *testing.T) {
	assert := assert.New(t)
	f, _ := os.Create("config-test.json")
	f.Write([]byte(""))
	_, err := Init("config-test.json")
	f.Close()
	assert.Error(err)
}
