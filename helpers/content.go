package helpers

import (
	"strings"

	"gopkg.in/alexcesaro/statsd.v2"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

type baseHelper struct {
	sql   gateways.SQL
	stats *statsd.Client
}

/*ContentI describes the functions for manipulating content models*/
type ContentI interface {
	GetByID(string) (*models.Content, error)
	GetAll(int, int) ([]*models.Content, error)
	Insert(*models.Content) error
	Update(*models.Content) error
	SetStatus(string, models.ContentStatus) error
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
	rows, err := c.sql.Select("SELECT id, contentType, text, parameters, status, subject FROM content WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	content, err := models.ContentFromSQL(rows)
	if err != nil {
		return nil, err
	}
	return content[0], err
}

/*GetAll returns <limit> content entries from <offset> number*/
func (c *Content) GetAll(offset int, limit int) ([]*models.Content, error) {
	rows, err := c.sql.Select("SELECT id, contentType, text, parameters, status, subject FROM content WHERE status=? ORDER BY id ASC LIMIT ?,?", models.ACTIVE, offset, limit)
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
		"INSERT INTO content (id, contentType, text, parameters, status, subject)VALUE(?, ?, ?, ?, ?, ?)",
		content.ID,
		string(content.Type),
		content.Text,
		strings.Join(content.Params, ","),
		string(models.ACTIVE),
		content.Subject,
	)
	return err

}

/*Update upserts the content with the given id*/
func (c *Content) Update(content *models.Content) error {
	err := c.sql.Modify("UPDATE content SET contentType=?,text=?,parameters=?,status=?,subject=? WHERE id=?",
		string(content.Type),
		content.Text,
		strings.Join(content.Params, ","),
		string(content.Status),
		content.Subject,
		content.ID,
	)
	return err
}

/*SetStatus updates the status of the content with the given id*/
func (c *Content) SetStatus(id string, status models.ContentStatus) error {
	err := c.sql.Modify("UPDATE content SET status=? WHERE id=?", string(status), id)
	return err
}
