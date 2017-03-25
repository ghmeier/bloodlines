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
	Time() gin.HandlerFunc
	GetJWT() gin.HandlerFunc
}

/*Trigger implements TriggerI and uses a trigger helper*/
type Trigger struct {
	*BaseHandler
	Trigger helpers.TriggerI
	Receipt helpers.ReceiptI
	Content helpers.ContentI
}

/*NewTrigger constructs and returns reference to a Trigger handler*/
func NewTrigger(ctx *GatewayContext) *Trigger {
	stats := ctx.Stats.Clone(statsd.Prefix("api.trigger"))
	return &Trigger{
		Trigger:     helpers.NewTrigger(ctx.Sql),
		Receipt:     helpers.NewReceipt(ctx.Sql, ctx.Sendgrid, ctx.TownCenter, ctx.Rabbit),
		Content:     helpers.NewContent(ctx.Sql),
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

	trigger, err := t.Trigger.Get(json.Key)
	if err != nil {
		t.ServerError(ctx, err, json)
		return
	} else if trigger != nil {
		t.UserError(ctx, "ERROR: trigger already exists with that key", json)
		return
	}

	trigger = models.NewTrigger(json.ContentID, json.Key, json.Values)
	err = t.Trigger.Insert(trigger)
	if err != nil {
		t.ServerError(ctx, err, json)
		return
	}

	t.Success(ctx, trigger)
}

/*ViewAll returns a list of trigger entites based on the offset and limit (default 0, 20)*/
func (t *Trigger) ViewAll(ctx *gin.Context) {
	offset, limit := t.GetPaging(ctx)
	triggers, err := t.Trigger.GetAll(offset, limit)
	if err != nil {
		t.ServerError(ctx, err, nil)
		return
	}

	t.Success(ctx, triggers)
}

/*View returns a trigger with id provided*/
func (t *Trigger) View(ctx *gin.Context) {
	key := ctx.Param("key")

	trigger, err := t.Trigger.Get(key)
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

	err = t.Trigger.Update(key, json.ContentID, json.Values)
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

	err := t.Trigger.Delete(key)
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

	trigger, err := t.Trigger.Get(key)
	if err != nil {
		t.ServerError(ctx, err, json)
		return
	}

	if trigger == nil {
		t.NotFoundError(ctx, "Error: no trigger found")
		return
	}

	content, err := t.Content.Get(trigger.ContentID.String())
	if err != nil {
		t.ServerError(ctx, err, trigger)
		return
	}

	params := make(map[string]string)
	for _, v := range content.Params {
		if json.Values[v] != "" {
			params[v] = json.Values[v]
		} else if trigger.Values[v] != "" {
			params[v] = trigger.Values[v]
		} else {
			t.UserError(ctx, "Error: no value for parameter "+v, content)
			return
		}
	}

	receipt := models.NewReceipt(json.Values, trigger.ContentID, json.UserID)
	err = t.Receipt.Insert(receipt)
	if err != nil {
		t.ServerError(ctx, err, content)
		return
	}

	request := &models.SendRequest{
		ReceiptID: receipt.ID,
		ContentID: content.ID,
	}
	t.Receipt.Send(request)

	t.Success(ctx, request)
}
