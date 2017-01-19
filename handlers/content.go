package handlers

import (
	"github.com/pborman/uuid"
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

/*ContentI has the methods for a content handler*/
type ContentI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
}

/*Content is the handler for all content api calls*/
type Content struct {
	*BaseHandler
	Helper helpers.ContentI
}

/*NewContent returns a content handler*/
func NewContent(ctx *GatewayContext) ContentI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.content"))
	return &Content{
		Helper:      helpers.NewContent(ctx.Sql),
		BaseHandler: NewBaseHandler(stats),
	}
}

/*New adds the given content entry to the database*/
func (c *Content) New(ctx *gin.Context) {
	var json models.Content

	err := ctx.BindJSON(&json)
	if err != nil {
		c.UserError(ctx, "Error: Unable to parse json", err)
		return
	}

	content := models.NewContent("EMAIL", json.Text, json.Subject, json.Params)
	err = c.Helper.Insert(content)
	if err != nil {
		c.ServerError(ctx, err, json)
		return
	}

	c.Success(ctx, content)
}

/*ViewAll returns a list of content with limit and offset
  determining the entries and amount (default 0,20)*/
func (c *Content) ViewAll(ctx *gin.Context) {
	offset, limit := c.GetPaging(ctx)
	content, err := c.Helper.GetAll(offset, limit)
	if err != nil {
		c.ServerError(ctx, err, content)
		return
	}

	c.Success(ctx, content)
}

/*View returns a content described by the given id*/
func (c *Content) View(ctx *gin.Context) {
	id := ctx.Param("contentId")

	content, err := c.Helper.GetByID(id)
	if err != nil {
		c.ServerError(ctx, err, content)
		return
	}

	c.Success(ctx, content)
}

/*Update overwrites content data for the content with the given id*/
func (c *Content) Update(ctx *gin.Context) {
	id := ctx.Param("contentId")

	var json models.Content
	err := ctx.BindJSON(&json)
	if err != nil {
		c.UserError(ctx, "Error: Unable to parse json", err)
		return
	}
	json.ID = uuid.Parse(id)

	err = c.Helper.Update(&json)
	if err != nil {
		c.ServerError(ctx, err, json)
		return
	}

	c.Success(ctx, json)
}

/*Deactivate sets a content's status to INACTIVE*/
func (c *Content) Deactivate(ctx *gin.Context) {
	id := ctx.Param("contentId")

	err := c.Helper.SetStatus(id, models.INACTIVE)
	if err != nil {
		c.ServerError(ctx, err, nil)
		return
	}

	c.Success(ctx, id)
}
