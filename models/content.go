package models

import(
	"database/sql"

	"github.com/pborman/uuid"
)

type Content struct {
	Id uuid.UUID `json:"id"`
	Type string `json:"contentType"`
	Text string `json:"text"`
	Params []string `json:"parameters"`
	Status bool `json:"status"`
}

func NewContent(Type string, Text string, Params []string) *Content {
	return &Content{
		Id: uuid.NewUUID(),
		Type: Type,
		Text: Text,
		Params: Params,
		Status: true,
	}
}

func ContentFromSql(rows *sql.Rows) ([]*Content, error) {
	content := make([]*Content,0)

	for rows.Next() {
		c := &Content{}
		rows.Scan(&c.Id, &c.Type, &c.Text, &c.Params, &c.Status)
		content = append(content, c)
	}

	return content, nil
}