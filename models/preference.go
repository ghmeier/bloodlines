package models

import (
	"database/sql"
	"errors"

	"github.com/pborman/uuid"
)

/*Preference entity data*/
type Preference struct {
	ID     uuid.UUID       `json:"id"`
	UserID uuid.UUID       `json:"userId"`
	Email  PreferenceState `json:"email"`
}

/*NewPreference contstructs and returns a new preference entity with it's id*/
func NewPreference(userID uuid.UUID) *Preference {
	return &Preference{
		ID:     uuid.NewUUID(),
		UserID: userID,
		Email:  SUBSCRIBED,
	}
}

/*PreferencesFromSQL returns a preference splice from sql rows*/
func PreferencesFromSQL(rows *sql.Rows) ([]*Preference, error) {
	preferences := make([]*Preference, 0)

	for rows.Next() {
		p := &Preference{}
		var email string
		rows.Scan(&p.ID, &p.UserID, &email)

		var ok bool
		p.Email, ok = toPreferenceState(email)
		if !ok {
			return nil, errors.New("invalid email preference")
		}
		preferences = append(preferences, p)
	}

	return preferences, nil
}

func toPreferenceState(s string) (PreferenceState, bool) {
	switch s {
	case SUBSCRIBED:
		return SUBSCRIBED, true
	case UNSUBSCRIBED:
		return UNSUBSCRIBED, true
	case MINIMAL:
		return MINIMAL, true
	default:
		return "", false
	}
}

/*PreferenceState wraps valid preference state strings*/
type PreferenceState string

/*valid preference states*/
const (
	SUBSCRIBED   = "SUBSCRIBED"
	UNSUBSCRIBED = "UNSUBSCRIBED"
	MINIMAL      = "MINIMAL"
)
