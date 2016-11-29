package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/pborman/uuid"
)

type Job struct {
	Id         uuid.UUID   `json:"id"`
	SendTime   time.Time   `json:"sendTime"`
	SendStatus Status      `json:"sendStatus"`
	Receipts   []uuid.UUID `json:"receipts"`
}

func NewJob(receipts []uuid.UUID, sendTime time.Time) *Job {
	return &Job{
		Id:         uuid.NewUUID(),
		SendTime:   sendTime,
		SendStatus: READY,
		Receipts:   receipts,
	}
}

func JobFromSql(rows *sql.Rows) ([]*Job, error) {
	jobs := make([]*Job, 0)

	for rows.Next() {
		j := &Job{}

		var receiptList, jState string
		rows.Scan(&j.Id, &j.SendTime, &jState, receiptList)

		j.Receipts = make([]uuid.UUID, 0)
		receipts := strings.Split(receiptList, ",")
		for _, receipt := range receipts {
			j.Receipts = append(j.Receipts, uuid.Parse(receipt))
		}

		var ok bool
		j.SendStatus, ok = toStatus(jState)
		if !ok {
			return nil, errors.New("Invalid Send State.")
		}

		jobs = append(jobs, j)
	}

	return jobs, nil
}
