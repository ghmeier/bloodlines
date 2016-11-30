package models

import (
	"database/sql"

	"github.com/pborman/uuid"
)

/*Trigger stores trigger entity data*/
type Trigger struct {
	ID        uuid.UUID `json:"id"`
	ContentID uuid.UUID `json:"contentId"`
	Key       string    `json:"key"`
	Params    []string  `json:"params"`
}

/*NewTrigger creates and returns a new trigger entity with id*/
func NewTrigger(contentID uuid.UUID, key string, params []string) *Trigger {
	return &Trigger{
		ID:        uuid.NewUUID(),
		ContentID: contentID,
		Key:       key,
		Params:    params,
	}
}

/*TriggerFromSQL returns a trigger splice from sql rows*/
func TriggerFromSQL(rows *sql.Rows) ([]*Trigger, error) {
	trigger := make([]*Trigger, 0)

	for rows.Next() {
		t := &Trigger{}

		var paramList string
		rows.Scan(&t.ID, &t.ContentID, &t.Key, &paramList)
		t.Params = toList(paramList)

		trigger = append(trigger, t)
	}

	return trigger, nil
}
