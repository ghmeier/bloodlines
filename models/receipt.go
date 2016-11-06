package models

import(
	"database/sql"
	"time"

	"github.com/pborman/uuid"
)

type Receipt struct {
	Id uuid.UUID `json:"id"`
	Created time.Time `json:"ts"`
	Values []string `json:"vals"`
	SendState Status `json:"sendState"`
	ContentId uuid.UUID `json:"contentId"`
}

func NewReceipt(Values []string, ContentId uuid.UUID) *Receipt {
	return &Receipt{
		Id: uuid.NewUUID(),
		Values: Values,
		SendState: READY,
		Created: time.Now(),
		ContentId: ContentId,
	}
}

func ReceiptFromSql(rows *sql.Rows) ([]*Receipt, error) {
	receipts := make([]*Receipt,0)

	for rows.Next() {
		r := &Receipt{}
		rows.Scan(&r.Id, &r.Created, &r.Values, &r.SendState, &r.ContentId)
		receipts = append(receipts, r)
	}

	return receipts, nil
}

type Status string
const (  // iota is reset to 0
        READY Status = "READY"
        QUEUED = "QUEUED"
        SUCCESS  = "SUCCESS"
        FAILURE = "FAILURE"
)


/*
	uuid		recipient
	[]string	values
	uuid		content_id
	uuid		id
	timestamp	time_sent
	send_status	status
*/