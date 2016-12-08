package handlers

import (
	"github.com/pborman/uuid"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
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
	Helper helpers.ContentI
}

/*NewContent returns a content handler*/
func NewContent(sql gateways.SQL) ContentI {
	return &Content{Helper: helpers.NewContent(sql)}
}

/*New adds the given content entry to the database*/
func (c *Content) New(ctx *gin.Context) {
	var json models.Content
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse("Invalid Content Object"))
		return
	}

	content := models.NewContent("EMAIL", json.Text, json.Subject, json.Params)
	err = c.Helper.Insert(content)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": content})
}

/*ViewAll returns a list of content with limit and offset
  determining the entries and amount (default 0,20)*/
func (c *Content) ViewAll(ctx *gin.Context) {
	offset, limit := getPaging(ctx)
	content, err := c.Helper.GetAll(offset, limit)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": content})
}

/*View returns a content described by the given id*/
func (c *Content) View(ctx *gin.Context) {
	id := ctx.Param("contentId")

	content, err := c.Helper.GetByID(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": content})
}

/*Update overwrites content data for the content with the given id*/
func (c *Content) Update(ctx *gin.Context) {
	id := ctx.Param("contentId")

	var json models.Content
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(400, errResponse(err.Error()))
		return
	}
	json.ID = uuid.Parse(id)

	err = c.Helper.Update(&json)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": json})
}

/*Deactivate sets a content's status to INACTIVE*/
func (c *Content) Deactivate(ctx *gin.Context) {
	id := ctx.Param("contentId")

	err := c.Helper.SetStatus(id, models.INACTIVE)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, empty())
}
