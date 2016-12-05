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

func toMap(s []string) map[string]string {
	var m map[string]string
	for _, v := range s {
		kv := strings.Split(v, ":")
		m[kv[0]] = kv[1]
	}

	return m
}

func toUUIDList(s string) []uuid.UUID {
	ids := make([]uuid.UUID, 0)
	if s == "" {
		return ids
	}

	receipts := strings.Split(s, ",")
	for _, receipt := range receipts {
		ids = append(ids, uuid.Parse(receipt))
	}
	return ids
}