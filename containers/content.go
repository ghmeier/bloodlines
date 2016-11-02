package containers

import(
	"database/sql"
	"fmt"

	"github.com/pborman/uuid"
)

type Content struct {
	Id uuid.UUID `json:"id"`
	ContentType string `json:"ContentType"`
	Text string `json:"text"`
	Parameters []string `json:"parameters"`
	Active bool `json:"active"`
}

func FromSql(rows *sql.Rows) ([]*Content, error) {
	content := make([]*Content,0)

	for rows.Next() {
		c := &Content{}
		rows.Scan(&c.Id, &c.ContentType, &c.Text, &c.Parameters, &c.Active)
		content = append(content, c)
	}

	return content, nil
}