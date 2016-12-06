package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

/*TriggerI describes the methods for a Trigger handler*/
type TriggerI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Remove(ctx *gin.Context)
	Activate(ctx *gin.Context)
}

/*Trigger implements TriggerI and uses a trigger helper*/
type Trigger struct {
	Helper helpers.TriggerI
}

/*NewTrigger constructs and returns reference to a Trigger handler*/
func NewTrigger(sql gateways.SQL) *Trigger {
	return &Trigger{Helper: helpers.NewTrigger(sql)}
}

/*New creates a new trigger entity based on the given data */
func (t *Trigger) New(ctx *gin.Context) {
	var json models.Trigger
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse("Invalid Trigger Object"))
		return
	}

	trigger := models.NewTrigger(json.ContentID, json.Key, json.Params)
	err = t.Helper.Insert(trigger)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": trigger})
}

/*ViewAll returns a list of trigger entites based on the offset and limit (default 0, 20)*/
func (t *Trigger) ViewAll(ctx *gin.Context) {
	offset, limit := getPaging(ctx)
	triggers, err := t.Helper.GetAll(offset, limit)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": triggers})
}

/*View returns a trigger with id provided*/
func (t *Trigger) View(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.JSON(400, errResponse("key is a required parameter"))
	}

	trigger, err := t.Helper.GetByKey(key)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
	}

	ctx.JSON(200, gin.H{"data": trigger})
}

/*Update overwrites the trigger entity with new values provided*/
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

	err = t.Helper.Update(key, json.ContentID, json.Params)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	json.Key = key
	ctx.JSON(200, gin.H{"data": json})
}

/*Remove sets a trigger to inactive*/
func (t *Trigger) Remove(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.JSON(400, errResponse("key is a required parameter"))
	}

	err := t.Helper.Delete(key)
	if err != nil {
		ctx.JSON(200, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, empty())
}

/*Activate starts a trigger's action*/
func (t *Trigger) Activate(ctx *gin.Context) {
	/* TODO: sent an email based on the trigger
	   and posted values */

	ctx.JSON(200, empty())
}
