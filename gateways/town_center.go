package gateways

import (
	"github.com/ghmeier/bloodlines/config"

	"github.com/pborman/uuid"
)

/*TownCenterI describes the functions for interacting with town center*/
type TownCenterI interface {
	GetUser(id uuid.UUID) (string, error)
}

/*TownCenter contains instrumentation for accessing TownCenter service*/
type TownCenter struct {
}

/*NewTownCenter creates and returns a TownCenter gateway*/
func NewTownCenter(config config.TownCenter) TownCenterI {
	return &TownCenter{}
}

/*GetUser stub function which returns an email based on a user ID */
func (t *TownCenter) GetUser(id uuid.UUID) (string, error) {
	switch id.String() {
	case "afee47b7-4eff-4826-a12c-86affd91a2d9":
		return "meier.garret@gmail.com", nil
	default:
		return "INVALID", nil
	}

}
