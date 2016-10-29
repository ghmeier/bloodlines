package handlers

import (
	"github.com/ghmeier/bloodlines/gateways"

	"gopkg.in/gin-gonic/gin.v1"
	"github.com/pborman/uuid"
)

type ContentI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
}

type Content struct {
	sql *gateways.Sql
}

func NewContent(sql *gateways.Sql) ContentI {
	return &Content{sql: sql}
}

func (c *Content) New(ctx *gin.Context) {
	err := c.sql.Modify(
		"INSERT INTO content VALUE(?, ?, ?, ?, ?)",
		uuid.New(),
		"EMAIL",
		"TEST",
		nil,
		true)
	if err != nil {
		ctx.JSON(500, &gin.H{"error": err, "message": err.Error()})
		return
	}
	ctx.JSON(200, empty())
}

func (c *Content) ViewAll(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (c *Content) View(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (c *Content) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (c *Content) Deactivate(ctx *gin.Context) {
	ctx.JSON(200, empty())
}
