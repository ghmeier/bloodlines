package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

type TriggerI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Remove(ctx *gin.Context)
	Activate(ctx *gin.Context)
}

type Trigger struct {
	helper *helpers.Trigger
}

func NewTrigger(sql gateways.Sql) *Trigger {
	return &Trigger{helper: helpers.NewTrigger(sql)}
}

func (t *Trigger) New(ctx *gin.Context) {
	var json models.Trigger
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse("Invalid Trigger Object"))
		return
	}

	trigger := models.NewTrigger(json.ContentId, json.Key, json.Params)
	err = t.helper.Insert(trigger)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": trigger})
}

func (t *Trigger) ViewAll(ctx *gin.Context) {
	offset, limit := getPaging(ctx)
	triggers, err := t.helper.GetAll(offset, limit)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": triggers})
}

func (t *Trigger) View(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.JSON(400, errResponse("key is a required parameter"))
	}

	trigger, err := t.helper.GetByKey(key)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
	}

	ctx.JSON(200, gin.H{"data": trigger})
}

func (t *Trigger) Update(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.JSON(400, errResponse("key is a required parameter"))
	}

	var json models.Trigger
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(400, errResponse(err.Error()))
		return
	}

	err = t.helper.Update(key, json.ContentId, json.Params)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	json.Key = key
	ctx.JSON(200, gin.H{"data": json})
}

func (t *Trigger) Remove(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.JSON(400, errResponse("key is a required parameter"))
	}

	err := t.helper.Delete(key)
	if err != nil {
		ctx.JSON(200, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, empty())
}

func (t *Trigger) Activate(ctx *gin.Context) {
	/* TODO: sent an email based on the trigger
	   and posted values */

	ctx.JSON(200, empty())
}
