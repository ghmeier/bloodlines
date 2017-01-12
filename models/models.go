package models

import (
	"encoding/json"
	"strings"

	"github.com/pborman/uuid"
)

/*SerializeValues returns a string representation of the given map like in JSON*/
func SerializeValues(values map[string]string) string {
	s, _ := json.Marshal(values)
	return string(s)
}

func toList(s string) []string {
	if s == "" {
		return make([]string, 0)
	}

	return strings.Split(s, ",")
}

func toMap(s []string) map[string]string {
	m := make(map[string]string)
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
