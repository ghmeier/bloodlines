package handlers

import(
	"fmt"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"
	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/containers"
	"github.com/ghmeier/bloodlines/gateways"

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
	var json containers.Content
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse("Invalid Content Object"))
		fmt.Printf("%s",err.Error())
		return
	}

	err = c.sql.Modify(
		"INSERT INTO content VALUE(?, ?, ?, ?, ?)",
		uuid.New(),
		"EMAIL",
		json.Text,
		strings.Join(json.Parameters,","),
		true)
	if err != nil {
		ctx.JSON(500, &gin.H{"error": err, "message": err.Error()})
		return
	}
	ctx.JSON(200, empty())
}

func (c *Content) ViewAll(ctx *gin.Context) {
	rows, err := c.sql.Select("SELECT * FROM content")
	if err != nil {
		 ctx.JSON(500, errResponse(err.Error()))
		 return
	}
	content, err := containers.FromSql(rows)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": content})
}

func (c *Content) View(ctx *gin.Context) {
	//var json models.Content
	id := ctx.Param("contentId")
	if id == "" {
		ctx.JSON(500, errResponse("contentId is a required parameter"))
		return
	}

	rows, err := c.sql.Select("SELECT * FROM content WHERE id=?", id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	content, err := containers.FromSql(rows)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": content})
}

func (c *Content) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (c *Content) Deactivate(ctx *gin.Context) {
	ctx.JSON(200, empty())
}
