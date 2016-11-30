package models

import (
	"strings"

	"github.com/pborman/uuid"
)

func toList(s string) []string {
	if s == "" {
		return make([]string, 0)
	}

	return strings.Split(s, ",")
}

func toUUIDList(s string) []uuid.UUID {
	ids := make([]uuid.UUID, 0)
	receipts := strings.Split(s, ",")
	for _, receipt := range receipts {
		ids = append(ids, uuid.Parse(receipt))
	}
	return ids
}
