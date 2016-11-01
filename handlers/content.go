package handlers

import (
	"fmt"
	"strings"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"

	"gopkg.in/gin-gonic/gin.v1"
	"github.com/pborman/uuid"
)

type ContentIfc interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
}

type Content struct {
	sql *gateways.Sql
}

func NewContent(sql *gateways.Sql) ContentIfc {
	return &Content{sql: sql}
}

func (c *Content) New(ctx *gin.Context) {
	var json models.Content
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, gin.H{"error":"Invalid Content Object"})
		fmt.Printf("%s",err.Error())
		return
	}

	err = c.sql.Modify(
		"INSERT INTO content VALUE(?, ?, ?, ?, ?)",
		uuid.New(),
		"EMAIL",
		json.Content,
		strings.Join(json.Parameters,","),
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
	//var json models.Content

	ctx.JSON(200, empty())
}

func (c *Content) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (c *Content) Deactivate(ctx *gin.Context) {
	ctx.JSON(200, empty())
}
