package models

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/pborman/uuid"
)

/*Trigger stores trigger entity data*/
type Trigger struct {
	ID        uuid.UUID         `json:"id"`
	ContentID uuid.UUID         `json:"contentId"`
	Key       string            `json:"tkey"`
	Values    map[string]string `json:"values"`
}

/*NewTrigger creates and returns a new trigger entity with id*/
func NewTrigger(contentID uuid.UUID, key string, values map[string]string) *Trigger {
	return &Trigger{
		ID:        uuid.NewUUID(),
		ContentID: contentID,
		Key:       key,
		Values:    values,
	}
}

/*TriggerFromSQL returns a trigger splice from sql rows*/
func TriggerFromSQL(rows *sql.Rows) ([]*Trigger, error) {
	trigger := make([]*Trigger, 0)
	defer rows.Close()

	for rows.Next() {
		t := &Trigger{}

		var valueList string
		rows.Scan(&t.ID, &t.ContentID, &t.Key, &valueList)
		err := json.Unmarshal([]byte(valueList), &t.Values)
		if err != nil {
			return nil, errors.New("invalid value list")
		}

		trigger = append(trigger, t)
	}

	return trigger, nil
}
