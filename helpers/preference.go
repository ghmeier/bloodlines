package helpers

import (
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

type Preference struct {
	*baseHelper
}

func NewPreference(sql gateways.Sql) *Preference {
	return &Preference{baseHelper: &baseHelper{sql: sql}}
}

func (p *Preference) Insert(preference *models.Preference) error {
	err := p.sql.Modify("INSERT INTO preference (id, userId, email) VALUES (?, ?, ?)",
		preference.Id,
		preference.UserId,
		preference.Email,
	)
	return err
}

func (p *Preference) GetAll() ([]*models.Preference, error) {
	rows, err := p.sql.Select("SELECT id, userId, email FROM preference")
	if err != nil {
		return nil, err
	}

	preferences, err := models.PreferencesFromSql(rows)
	if err != nil {
		return nil, err
	}

	return preferences, nil
}

func (p *Preference) GetPreferenceByUserId(id string) (*models.Preference, error) {
	rows, err := p.sql.Select("SELECT id, userId, email from preference where userId=?", id)
	if err != nil {
		return nil, err
	}

	preferences, err := models.PreferencesFromSql(rows)
	if err != nil {
		return nil, err
	}

	return preferences[0], nil
}

func (p *Preference) Update(preference *models.Preference) error {
	err := p.sql.Modify("UPDATE preference SET email=? WHERE userId=?", preference.Email, preference.Id)
	return err
}
