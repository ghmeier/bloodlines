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
	helper helpers.ContentI
}

/*NewContent returns a content handler*/
func NewContent(sql gateways.SQL) ContentI {
	return &Content{helper: helpers.NewContent(sql)}
}

/*New adds the given content entry to the database*/
func (c *Content) New(ctx *gin.Context) {
	var json models.Content
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse("Invalid Content Object"))
		return
	}

	content := models.NewContent("EMAIL", json.Text, json.Params)
	err = c.helper.Insert(content)
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
	content, err := c.helper.GetAll(offset, limit)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": content})
}

/*View returns a content described by the given id*/
func (c *Content) View(ctx *gin.Context) {
	id := ctx.Param("contentId")
	if id == "" {
		ctx.JSON(400, errResponse("contentId is a required parameter"))
		return
	}

	content, err := c.helper.GetByID(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": content})
}

/*Update overwrites content data for the content with the given id*/
func (c *Content) Update(ctx *gin.Context) {
	id := ctx.Param("contentId")
	if id == "" {
		ctx.JSON(400, errResponse("contentId is a required parameter"))
	}

	var json models.Content
	err := ctx.BindJSON(&json)
	json.ID = uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, errResponse(err.Error()))
		return
	}

	err = c.helper.Update(&json)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": json})
}

/*Deactivate sets a content's status to INACTIVE*/
func (c *Content) Deactivate(ctx *gin.Context) {
	id := ctx.Param("contentId")
	if id == "" {
		ctx.JSON(400, errResponse("contentId is a required parameter"))
	}

	err := c.helper.SetStatus(id, models.INACTIVE)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, empty())
}
