package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/pborman/uuid"
)

/*Job entity data*/
type Job struct {
	ID         uuid.UUID   `json:"id"`
	SendTime   time.Time   `json:"sendTime"`
	SendStatus Status      `json:"sendStatus"`
	Receipts   []uuid.UUID `json:"receipts"`
}

/*NewJob constructs and returns a new Job with a new uuid*/
func NewJob(receipts []uuid.UUID, sendTime time.Time) *Job {
	return &Job{
		ID:         uuid.NewUUID(),
		SendTime:   sendTime,
		SendStatus: READY,
		Receipts:   receipts,
	}
}

/*JobFromSQL returns a Job splice from sql rows*/
func JobFromSQL(rows *sql.Rows) ([]*Job, error) {
	jobs := make([]*Job, 0)

	for rows.Next() {
		j := &Job{}

		var receiptList, jState string
		rows.Scan(&j.ID, &j.SendTime, &jState, receiptList)

		j.Receipts = toUUIDList(receiptList)

		var ok bool
		j.SendStatus, ok = toStatus(jState)
		if !ok {
			return nil, errors.New("invalid send state")
		}

		jobs = append(jobs, j)
	}

	return jobs, nil
}
