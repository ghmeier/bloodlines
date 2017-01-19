package handlers

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

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
	*BaseHandler
	Helper  helpers.TriggerI
	RHelper helpers.ReceiptI
	CHelper helpers.ContentI
}

/*NewTrigger constructs and returns reference to a Trigger handler*/
func NewTrigger(ctx *GatewayContext) *Trigger {
	stats := ctx.Stats.Clone(statsd.Prefix("api.trigger"))
	return &Trigger{
		Helper:      helpers.NewTrigger(ctx.Sql),
		RHelper:     helpers.NewReceipt(ctx.Sql, ctx.Sendgrid, ctx.TownCenter, ctx.Rabbit),
		CHelper:     helpers.NewContent(ctx.Sql),
		BaseHandler: NewBaseHandler(stats),
	}
}

/*New creates a new trigger entity based on the given data */
func (t *Trigger) New(ctx *gin.Context) {
	var json models.Trigger
	err := ctx.BindJSON(&json)

	if err != nil {
		t.UserError(ctx, "Error: unable to parse json", err)
		return
	}

	trigger := models.NewTrigger(json.ContentID, json.Key, json.Values)
	err = t.Helper.Insert(trigger)
	if err != nil {
		t.ServerError(ctx, err, json)
		return
	}

	t.Success(ctx, trigger)
}

/*ViewAll returns a list of trigger entites based on the offset and limit (default 0, 20)*/
func (t *Trigger) ViewAll(ctx *gin.Context) {
	offset, limit := t.GetPaging(ctx)
	triggers, err := t.Helper.GetAll(offset, limit)
	if err != nil {
		t.ServerError(ctx, err, nil)
		return
	}

	t.Success(ctx, triggers)
}

/*View returns a trigger with id provided*/
func (t *Trigger) View(ctx *gin.Context) {
	key := ctx.Param("key")

	trigger, err := t.Helper.GetByKey(key)
	if err != nil {
		t.ServerError(ctx, err, nil)
		return
	}

	t.Success(ctx, trigger)
}

/*Update overwrites the trigger entity with new values provided*/
func (t *Trigger) Update(ctx *gin.Context) {
	key := ctx.Param("key")

	var json models.Trigger
	err := ctx.BindJSON(&json)
	if err != nil {
		t.UserError(ctx, "Error: unable to parse json", err)
		return
	}

	err = t.Helper.Update(key, json.ContentID, json.Values)
	if err != nil {
		t.ServerError(ctx, err, json)
		return
	}

	json.Key = key
	t.Success(ctx, json)
}

/*Remove sets a trigger to inactive*/
func (t *Trigger) Remove(ctx *gin.Context) {
	key := ctx.Param("key")

	err := t.Helper.Delete(key)
	if err != nil {
		t.ServerError(ctx, err, nil)
		return
	}

	t.Success(ctx, nil)
}

/*Activate starts a trigger's action*/
func (t *Trigger) Activate(ctx *gin.Context) {
	key := ctx.Param("key")

	var json models.Receipt
	err := ctx.BindJSON(&json)
	if err != nil {
		t.UserError(ctx, "Error: unable to parse json", err)
		return
	}

	trigger, err := t.Helper.GetByKey(key)
	if err != nil {
		t.ServerError(ctx, err, json)
		return
	}

	if trigger == nil {
		t.UserError(ctx, "Error: no trigger found", nil)
		return
	}

	content, err := t.CHelper.GetByID(trigger.ContentID.String())
	if err != nil {
		t.ServerError(ctx, err, trigger)
		return
	}

	for k, v := range trigger.Values {
		json.Values[k] = v
	}

	receipt := models.NewReceipt(json.Values, trigger.ContentID, json.UserID)
	err = t.RHelper.Insert(receipt)
	if err != nil {
		t.ServerError(ctx, err, content)
		return
	}

	request := &models.SendRequest{
		ReceiptID: receipt.ID,
		ContentID: content.ID,
	}
	t.RHelper.Send(request)

	t.Success(ctx, request)
}
