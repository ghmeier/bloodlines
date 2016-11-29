package helpers

import (
	"strings"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

type baseHelper struct {
	sql gateways.Sql
}

type Content struct {
	*baseHelper
}

func NewContent(sql gateways.Sql) *Content {
	return &Content{baseHelper: &baseHelper{sql: sql}}
}

func (c *Content) GetById(id string) (*models.Content, error) {
	rows, err := c.sql.Select("SELECT id, contentType, text, parameters, status FROM content WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	content, err := models.ContentFromSql(rows)
	if err != nil {
		return nil, err
	}
	return content[0], nil
}

func (c *Content) GetAll(offset int, limit int) ([]*models.Content, error) {
	rows, err := c.sql.Select("SELECT id, contentType, text, parameters, status FROM content WHERE status=? ORDER BY id ASC LIMIT ?,?", models.ACTIVE, offset, limit)
	if err != nil {
		return nil, err
	}

	content, err := models.ContentFromSql(rows)
	if err != nil {
		return nil, err
	}

	return content, err
}

func (c *Content) Insert(content *models.Content) error {
	err := c.sql.Modify(
		"INSERT INTO content (id, contentType, text, parameters, status)VALUE(?, ?, ?, ?, ?)",
		content.Id,
		content.Type,
		content.Text,
		strings.Join(content.Params, ","),
		models.ACTIVE)
	if err != nil {
		return err
	}

	return nil

}

func (c *Content) Update(content *models.Content) error {
	err := c.sql.Modify("UPDATE content SET contentType=?,text=?,parameters=?,status=? WHERE id=?",
		content.Type,
		content.Text,
		strings.Join(content.Params, ","),
		content.Status,
		content.Id,
	)
	return err
}

func (c *Content) SetStatus(id string, status models.ContentStatus) error {
	err := c.sql.Modify("UPDATE content SET status=? WHERE id=?", status, id)
	return err
}
