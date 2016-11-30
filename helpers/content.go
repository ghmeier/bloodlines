package helpers

import (
	"strings"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

type baseHelper struct {
	sql gateways.SQL
}

/*Content is the helper for content entries*/
type Content struct {
	*baseHelper
}

/*NewContent returns a new Content helper*/
func NewContent(sql gateways.SQL) *Content {
	return &Content{baseHelper: &baseHelper{sql: sql}}
}

/*GetByID returns the content referenced by the provided id, otherwise nil*/
func (c *Content) GetByID(id string) (*models.Content, error) {
	rows, err := c.sql.Select("SELECT id, contentType, text, parameters, status FROM content WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	content, err := models.ContentFromSQL(rows)
	if err != nil {
		return nil, err
	}
	return content[0], nil
}

/*GetAll returns <limit> content entries from <offset> number*/
func (c *Content) GetAll(offset int, limit int) ([]*models.Content, error) {
	rows, err := c.sql.Select("SELECT id, contentType, text, parameters, status FROM content WHERE status=? ORDER BY id ASC LIMIT ?,?", models.ACTIVE, offset, limit)
	if err != nil {
		return nil, err
	}

	content, err := models.ContentFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return content, err
}

/*Insert adds the given content entry*/
func (c *Content) Insert(content *models.Content) error {
	err := c.sql.Modify(
		"INSERT INTO content (id, contentType, text, parameters, status)VALUE(?, ?, ?, ?, ?)",
		content.ID,
		content.Type,
		content.Text,
		strings.Join(content.Params, ","),
		models.ACTIVE)
	if err != nil {
		return err
	}

	return nil

}

/*Update upserts the content with the given id*/
func (c *Content) Update(content *models.Content) error {
	err := c.sql.Modify("UPDATE content SET contentType=?,text=?,parameters=?,status=? WHERE id=?",
		content.Type,
		content.Text,
		strings.Join(content.Params, ","),
		content.Status,
		content.ID,
	)
	return err
}

/*SetStatus updates the status of the content with the given id*/
func (c *Content) SetStatus(id string, status models.ContentStatus) error {
	err := c.sql.Modify("UPDATE content SET status=? WHERE id=?", status, id)
	return err
}
