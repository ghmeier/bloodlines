package helpers

import (
//	"strings"

//	"github.com/ghmeier/bloodlines/models"
	"github.com/ghmeier/bloodlines/gateways"
)

type Receipt struct{
	*baseHelper
}

func NewReceipt(sql *gateways.Sql) *Receipt {
	return &Receipt{baseHelper: &baseHelper{sql: sql}}
}