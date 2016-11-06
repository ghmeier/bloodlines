package models

import (
	"database/sql"
	"time"

	"github.com/pborman/uuid"
)

type Job struct {
	Id uuid.UUID `json:"id"`
	SendTime time.Time `json:"sendTime"`
	SendStatus Status `json:"sendStatus"`
	Receipts []uuid.UUID `json:"receipts"`
}

func NewJob(receipts []uuid.UUID, sendTime time.Time) *Job {
	return &Job{
		Id: uuid.NewUUID(),
		SendTime: sendTime,
		SendStatus: READY,
		Receipts: receipts,
	}
}

func JobFromSql(rows *sql.Rows) ([]*Job, error) {
	jobs := make([]*Job, 0)

	for rows.Next() {
		j := &Job{}
		rows.Scan(&j.Id, &j.SendTime, &j.SendStatus, &j.Receipts)
		jobs = append(jobs, j)
	}

	return jobs, nil
}

/*


Job
[]uuid		receipts
timestamp	time_sent
uuid		id
send_status	status
*/