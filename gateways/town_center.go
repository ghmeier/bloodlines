package gateways

import (
	"github.com/ghmeier/bloodlines/config"

	"github.com/pborman/uuid"
)

type TownCenterI interface {
	GetUser(id uuid.UUID) (string, error)
}

type TownCenter struct {
}

func NewTownCenter(config config.TownCenter) TownCenterI {
	return &TownCenter{}
}

/* Stub function which returns an email based on a user ID */
func (t *TownCenter) GetUser(id uuid.UUID) (string, error) {
	switch id.String() {
	case "afee47b7-4eff-4826-a12c-86affd91a2d9":
		return "meier.garret@gmail.com", nil
	default:
		return "INVALID", nil
	}

}
