package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/pborman/uuid"
)

/*Receipt stores data for receipts*/
type Receipt struct {
	ID        uuid.UUID         `json:"id"`
	Created   time.Time         `json:"ts"`
	Values    map[string]string `json:"values"`
	SendState Status            `json:"sendState"`
	ContentID uuid.UUID         `json:"contentId"`
	UserID    uuid.UUID         `json:"userId"`
}

/*SendRequest has the info for sending a message through the sendgrid api*/
type SendRequest struct {
	ReceiptID uuid.UUID `json:"receiptId"`
	ContentID uuid.UUID `json:"contentId"`
}

/*NewReceipt creates and returns a new receipt with a new id*/
func NewReceipt(values map[string]string, contentID uuid.UUID, userID uuid.UUID) *Receipt {
	return &Receipt{
		ID:        uuid.NewUUID(),
		Values:    values,
		SendState: READY,
		Created:   time.Now(),
		ContentID: contentID,
		UserID:    userID,
	}
}

/*ReceiptFromSQL returns a receipt splice from sql rows*/
func ReceiptFromSQL(rows *sql.Rows) ([]*Receipt, error) {
	receipts := make([]*Receipt, 0)

	for rows.Next() {
		r := &Receipt{}
		var valueList, rState string
		rows.Scan(&r.ID, &r.Created, &valueList, &rState, &r.ContentID, &r.UserID)

		err := json.Unmarshal([]byte(valueList), &r.Values)
		if err != nil {
			return nil, errors.New("invalid value list")
		}

		var ok bool
		r.SendState, ok = toStatus(rState)
		if !ok {
			return nil, errors.New("invalid send state")
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

/*Status wraps valid receipt status strings*/
type Status string

/*valid Statuses*/
const ( // iota is reset to 0
	READY   = "READY"
	QUEUED  = "QUEUED"
	SUCCESS = "SUCCESS"
	FAILURE = "FAILURE"
)
