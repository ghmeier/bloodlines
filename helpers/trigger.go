package helpers

import (
	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

/*TriggerI describes methods of the Trigger Helper*/
type TriggerI interface {
	GetAll(int, int) ([]*models.Trigger, error)
	GetByKey(string) (*models.Trigger, error)
	Update(string, uuid.UUID, map[string]string) error
	Insert(*models.Trigger) error
	Delete(string) error
}

/*Trigger is the helper for trigger entites*/
type Trigger struct {
	*baseHelper
}

/*NewTrigger constructs and returns a new Trigger helper*/
func NewTrigger(sql gateways.SQL) TriggerI {
	return &Trigger{baseHelper: &baseHelper{sql: sql}}
}

/*Insert creates a new trigger from the model and inserts it into the database*/
func (t *Trigger) Insert(trigger *models.Trigger) error {
	err := t.sql.Modify(
		"INSERT INTO b_trigger (id, contentId, tkey, vals) VALUES (?, ?, ?, ?)",
		trigger.ID,
		trigger.ContentID,
		trigger.Key,
		models.SerializeValues(trigger.Values),
	)
	return err
}

/*GetAll returns <limit> trigger entities starting at <offset>*/
func (t *Trigger) GetAll(offset int, limit int) ([]*models.Trigger, error) {
	rows, err := t.sql.Select("SELECT id, contentId, tkey, vals FROM b_trigger ORDER BY id ASC LIMIT ?,? ", offset, limit)
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
	rows, err := t.sql.Select("SELECT id, contentId, tkey, vals FROM b_trigger WHERE tkey=?", key)
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
func (t *Trigger) Update(key string, contentID uuid.UUID, values map[string]string) error {
	err := t.sql.Modify(
		"UPDATE b_trigger SET contentId=?,vals=? WHERE tkey=?",
		contentID,
		models.SerializeValues(values),
		key,
	)
	return err
}

/*Delete removes the trigger entry from the database*/
func (t *Trigger) Delete(key string) error {
	err := t.sql.Modify(
		"DELETE FROM b_trigger WHERE tkey=?",
		key,
	)
	return err
}
