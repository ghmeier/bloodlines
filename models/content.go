package models

import (
	"database/sql"
	"errors"

	"github.com/pborman/uuid"
)

/*Content is the representation of content entries in bloodlines*/
type Content struct {
	ID     uuid.UUID     `json:"id"`
	Type   ContentType   `json:"contentType"`
	Text   string        `json:"text"`
	Params []string      `json:"parameters"`
	Status ContentStatus `json:"status"`
}

/*NewContent constructs and returns a new content entity with a new uuid*/
func NewContent(contentType ContentType, text string, params []string) *Content {
	return &Content{
		ID:     uuid.NewUUID(),
		Type:   contentType,
		Text:   text,
		Params: params,
		Status: ACTIVE,
	}
}

/*ContentFromSQL returns a new content slice from a group of sql rows*/
func ContentFromSQL(rows *sql.Rows) ([]*Content, error) {
	content := make([]*Content, 0)
	defer rows.Close()

	for rows.Next() {
		c := &Content{}
		var paramList, cType, cStatus string
		rows.Scan(&c.ID, &cType, &c.Text, &paramList, &cStatus)

		c.Params = toList(paramList)

		var ok bool
		c.Type, ok = toContentType(cType)
		if !ok {
			return nil, errors.New("invalid contentType string")
		}

		c.Status, ok = toContentStatus(cStatus)
		if !ok {
			return nil, errors.New("invalid contentStatus string")
		}

		content = append(content, c)
	}

	return content, nil
}

func toContentType(s string) (ContentType, bool) {
	switch s {
	case EMAIL:
		return EMAIL, true
	default:
		return "", false
	}
}

func toContentStatus(s string) (ContentStatus, bool) {
	switch s {
	case ACTIVE:
		return ACTIVE, true
	case INACTIVE:
		return INACTIVE, true
	default:
		return "", false
	}
}

/*ContentType is an enum wrapper for valid content type*/
type ContentType string

/*valid ContentTypes*/
const (
	EMAIL = "EMAIL"
)

/*ContentStatus is an enum wrapper for valid ContentStatus strings*/
type ContentStatus string

/*valid ContentStatuses*/
const (
	ACTIVE   = "ACTIVE"
	INACTIVE = "INACTIVE"
)
