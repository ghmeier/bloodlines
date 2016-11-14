package models

import(
	"database/sql"

	"github.com/pborman/uuid"
)

type Preference struct {
	Id uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
	Email PreferenceState `json:"email"`
}

func NewPreference(userId uuid.UUID) *Preference {
	return &Preference{
		Id: uuid.NewUUID(),
		UserId: userId,
		Email: SUBSCRIBED,
	}
}

func PreferencesFromSql(rows *sql.Rows) ([]*Preference, error) {
	preferences := make([]*Preference,0)

	for rows.Next() {
		p := &Preference{}
		rows.Scan(&p.Id, &p.UserId, &p.Email)
		preferences = append(preferences, p)
	}

	return preferences, nil
}


type PreferenceState string
const(
	SUBSCRIBED = "SUBSCRIBED"
	UNSUBSCRIBED = "UNSUBSCRIBED"
	MINIMAL = "MINIMAL"
)