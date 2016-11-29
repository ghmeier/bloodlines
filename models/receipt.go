package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/pborman/uuid"
)

type Receipt struct {
	Id        uuid.UUID `json:"id"`
	Created   time.Time `json:"ts"`
	Values    []string  `json:"vals"`
	SendState Status    `json:"sendState"`
	ContentId uuid.UUID `json:"contentId"`
}

func NewReceipt(Values []string, ContentId uuid.UUID) *Receipt {
	return &Receipt{
		Id:        uuid.NewUUID(),
		Values:    Values,
		SendState: READY,
		Created:   time.Now(),
		ContentId: ContentId,
	}
}

func ReceiptFromSql(rows *sql.Rows) ([]*Receipt, error) {
	receipts := make([]*Receipt, 0)

	for rows.Next() {
		r := &Receipt{}
		var valueList, rState string
		rows.Scan(&r.Id, &r.Created, &valueList, &rState, &r.ContentId)

		r.Values = make([]string, 0)
		if valueList != "" {
			r.Values = strings.Split(valueList, ",")
		}

		var ok bool
		r.SendState, ok = toStatus(rState)
		if !ok {
			return nil, errors.New("Invalid Send State")
		}

		receipts = append(receipts, r)
	}

	return receipts, nil
}

func toStatus(s string) (Status, bool) {
	switch s {
	case READY:
		return READY, true
	case QUEUED:
		return QUEUED, true
	case SUCCESS:
		return SUCCESS, true
	case FAILURE:
		return FAILURE, true
	default:
		return "", false
	}
}

type Status string

const ( // iota is reset to 0
	READY   = "READY"
	QUEUED  = "QUEUED"
	SUCCESS = "SUCCESS"
	FAILURE = "FAILURE"
)
