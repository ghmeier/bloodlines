package models

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/pborman/uuid"
)

type Content struct {
	Id     uuid.UUID     `json:"id"`
	Type   ContentType   `json:"contentType"`
	Text   string        `json:"text"`
	Params []string      `json:"parameters"`
	Status ContentStatus `json:"status"`
}

func NewContent(contentType ContentType, text string, params []string) *Content {
	return &Content{
		Id:     uuid.NewUUID(),
		Type:   contentType,
		Text:   text,
		Params: params,
		Status: ACTIVE,
	}
}

func ContentFromSql(rows *sql.Rows) ([]*Content, error) {
	content := make([]*Content, 0)
	defer rows.Close()

	for rows.Next() {
		c := &Content{}
		var paramList, cType, cStatus string
		rows.Scan(&c.Id, &cType, &c.Text, &paramList, &cStatus)

		c.Params = make([]string, 0)
		if paramList != "" {
			c.Params = strings.Split(paramList, ",")
		}

		var ok bool
		c.Type, ok = toContentType(cType)
		if !ok {
			return nil, errors.New("Invalid contentType string.")
		}

		c.Status, ok = toContentStatus(cStatus)
		if !ok {
			return nil, errors.New("Invalid contentStatus string.")
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

type ContentType string

const (
	EMAIL = "EMAIL"
)

type ContentStatus string

const (
	ACTIVE   = "ACTIVE"
	INACTIVE = "INACTIVE"
)
