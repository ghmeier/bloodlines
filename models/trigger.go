package models

import (
	"database/sql"

	"github.com/pborman/uuid"
)

type Trigger struct {
	Id        uuid.UUID `json:"id"`
	ContentId uuid.UUID `json:"contentId"`
	Key       string    `json:"key"`
	Params    []string  `json:"params"`
}

func NewTrigger(contentId uuid.UUID, key string, params []string) *Trigger {
	return &Trigger{
		Id:        uuid.NewUUID(),
		ContentId: contentId,
		Key:       key,
		Params:    params,
	}
}

func TriggerFromSql(rows *sql.Rows) ([]*Trigger, error) {
	trigger := make([]*Trigger, 0)

	for rows.Next() {
		t := &Trigger{}

		var paramList string
		rows.Scan(&t.Id, &t.ContentId, &t.Key, &paramList)
		t.Params = toList(paramList)

		trigger = append(trigger, t)
	}

	return trigger, nil
}
