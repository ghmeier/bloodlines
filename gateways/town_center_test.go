package gateways

import (
	"testing"

	"github.com/ghmeier/bloodlines/config"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInvalidUser(t *testing.T) {
	assert := assert.New(t)

	s := getTownCenter()

	name, err := s.GetUser(uuid.NewUUID())

	assert.NoError(err)
	assert.EqualValues("INVALID", name)
}

func TestSpecificUser(t *testing.T) {
	assert := assert.New(t)

	s := getTownCenter()

	name, err := s.GetUser(uuid.Parse("afee47b7-4eff-4826-a12c-86affd91a2d9"))

	assert.NoError(err)
	assert.EqualValues("meier.garret@gmail.com", name)
}

func getTownCenter() TownCenterI {
	return NewTownCenter(config.TownCenter{})
}
