package helpers

import (
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

/*Preference is the helper for preference entities*/
type Preference struct {
	*baseHelper
}

/*NewPreference constructs and returns a Preference helper*/
func NewPreference(sql gateways.SQL) *Preference {
	return &Preference{baseHelper: &baseHelper{sql: sql}}
}

/*Insert adds a pereference entity to the database*/
func (p *Preference) Insert(preference *models.Preference) error {
	err := p.sql.Modify("INSERT INTO preference (id, userId, email) VALUES (?, ?, ?)",
		preference.ID,
		preference.UserID,
		string(preference.Email),
	)
	return err
}

/*GetAll returns a list of <limit> preferences starting from <offset>*/
func (p *Preference) GetAll(offset int, limit int) ([]*models.Preference, error) {
	rows, err := p.sql.Select("SELECT id, userId, email FROM preference ORDER BY id ASC LIMIT ?,?", offset, limit)
	if err != nil {
		return nil, err
	}

	preferences, err := models.PreferencesFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return preferences, nil
}

/*GetPreferenceByUserID returns a preference associated with the given user id*/
func (p *Preference) GetPreferenceByUserID(id string) (*models.Preference, error) {
	rows, err := p.sql.Select("SELECT id, userId, email FROM preference where userId=?", id)
	if err != nil {
		return nil, err
	}

	preferences, err := models.PreferencesFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return preferences[0], nil
}

/*Update sets the preference entry to the given values*/
func (p *Preference) Update(preference *models.Preference) error {
	err := p.sql.Modify("UPDATE preference SET email=? WHERE userId=?", string(preference.Email), preference.ID)
	return err
}
