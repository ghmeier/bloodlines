package helpers

import (
	"strings"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

type Trigger struct {
	*baseHelper
}

func NewTrigger(sql gateways.Sql) *Trigger {
	return &Trigger{baseHelper: &baseHelper{sql: sql}}
}

func (t *Trigger) Insert(trigger *models.Trigger) error {
	err := t.sql.Modify(
		"INSERT INTO trigger (id, contentId, key, params) VALUES (?, ?, ?, ?)",
		trigger.Id,
		trigger.ContentId,
		trigger.Key,
		strings.Join(trigger.Params, ","),
	)
	return err
}

func (t *Trigger) GetAll() ([]*models.Trigger, error) {
	rows, err := t.sql.Select("SELECT id, contentId, key, params FROM trigger")
	if err != nil {
		return nil, err
	}

	triggers, err := models.TriggerFromSql(rows)
	if err != nil {
		return nil, err
	}

	return triggers, nil
}

func (t *Trigger) GetByKey(key string) (*models.Trigger, error) {
	rows, err := t.sql.Select("SELECT id, contentId, key, params FROM receipt WHERE key=?", key)
	if err != nil {
		return nil, err
	}

	triggers, err := models.TriggerFromSql(rows)
	if err != nil {
		return nil, err
	}
	return triggers[0], nil
}

func (t *Trigger) Update(key string, contentId uuid.UUID, params []string) error {
	err := t.sql.Modify(
		"UPDATE trigger SET contentId=?,params=? WHERE key=?",
		contentId,
		strings.Join(params, ","),
		key,
	)
	return err
}

func (t *Trigger) Delete(key string) error {
	err := t.sql.Modify(
		"DELETE FROM trigger WHERE key=?",
		key,
	)
	return err
}
