package models

import (
	"database/sql"
	"errors"

	"github.com/pborman/uuid"
)

type Preference struct {
	Id     uuid.UUID       `json:"id"`
	UserId uuid.UUID       `json:"userId"`
	Email  PreferenceState `json:"email"`
}

func NewPreference(userId uuid.UUID) *Preference {
	return &Preference{
		Id:     uuid.NewUUID(),
		UserId: userId,
		Email:  SUBSCRIBED,
	}
}

func PreferencesFromSql(rows *sql.Rows) ([]*Preference, error) {
	preferences := make([]*Preference, 0)

	for rows.Next() {
		p := &Preference{}
		var email string
		rows.Scan(&p.Id, &p.UserId, &email)

		var ok bool
		p.Email, ok = toPreferenceState(email)
		if !ok {
			return nil, errors.New("Invalid Email Preference")
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

type PreferenceState string

const (
	SUBSCRIBED   = "SUBSCRIBED"
	UNSUBSCRIBED = "UNSUBSCRIBED"
	MINIMAL      = "MINIMAL"
)
