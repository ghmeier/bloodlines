package handlers

import(
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"

)

type ContentIfc interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
}

type Content struct {
	helper *helpers.Content
}

func NewContent(sql *gateways.Sql) ContentIfc {
	return &Content{helper: helpers.NewContent(sql)}
}

func (c *Content) New(ctx *gin.Context) {
	var json models.Content
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse("Invalid Content Object"))
		return
	}

	content := models.NewContent("EMAIL", json.Text, json.Params)
	err  = c.helper.Insert(content)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data":content})
}

func (c *Content) ViewAll(ctx *gin.Context) {
	content, err := c.helper.GetAll()
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": content})
}

func (c *Content) View(ctx *gin.Context) {
	id := ctx.Param("contentId")
	if id == "" {
		ctx.JSON(400, errResponse("contentId is a required parameter"))
		return
	}

	content, err := c.helper.GetById(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": content})
}

func (c *Content) Update(ctx *gin.Context) {
	id := ctx.Param("contentId")
	if id == "" {
		ctx.JSON(400, errResponse("contentId is a required parameter"))
	}

	var json models.Content
	err := ctx.BindJSON(&json)
	json.Id = uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, errResponse(err.Error()))
		return
	}

	err = c.helper.Update(&json)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data":json})
}

func (c *Content) Deactivate(ctx *gin.Context) {
	id := ctx.Param("contentId")
	if id == "" {
		ctx.JSON(400, errResponse("contentId is a required parameter"))
	}

	err := c.helper.SetStatus(id, false)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, empty())
}

