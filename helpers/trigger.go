package helpers

import (
	"strings"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

/*Trigger is the helper for trigger entites*/
type Trigger struct {
	*baseHelper
}

/*NewTrigger constructs and returns a new Trigger helper*/
func NewTrigger(sql gateways.SQL) *Trigger {
	return &Trigger{baseHelper: &baseHelper{sql: sql}}
}

/*Insert creates a new trigger from the model and inserts it into the database*/
func (t *Trigger) Insert(trigger *models.Trigger) error {
	err := t.sql.Modify(
		"INSERT INTO trigger (id, contentId, key, params) VALUES (?, ?, ?, ?)",
		trigger.ID,
		trigger.ContentID,
		trigger.Key,
		strings.Join(trigger.Params, ","),
	)
	return err
}

/*GetAll returns <limit> trigger entities starting at <offset>*/
func (t *Trigger) GetAll(offset int, limit int) ([]*models.Trigger, error) {
	rows, err := t.sql.Select("SELECT id, contentId, key, params FROM trigger ORDER BY id ASC LIMIT ?,? ", offset, limit)
	if err != nil {
		return nil, err
	}

	triggers, err := models.TriggerFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return triggers, nil
}

/*GetByKey given a string key, returns the trigger associated with it*/
func (t *Trigger) GetByKey(key string) (*models.Trigger, error) {
	rows, err := t.sql.Select("SELECT id, contentId, key, params FROM receipt WHERE key=?", key)
	if err != nil {
		return nil, err
	}

	triggers, err := models.TriggerFromSQL(rows)
	if err != nil {
		return nil, err
	}
	return triggers[0], nil
}

/*Update overwrites a trigger's entry with the given data*/
func (t *Trigger) Update(key string, contentID uuid.UUID, params []string) error {
	err := t.sql.Modify(
		"UPDATE trigger SET contentId=?,params=? WHERE key=?",
		contentID,
		strings.Join(params, ","),
		key,
	)
	return err
}

/*Delete removes the trigger entry from the database*/
func (t *Trigger) Delete(key string) error {
	err := t.sql.Modify(
		"DELETE FROM trigger WHERE key=?",
		key,
	)
	return err
}
