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
	Helper  helpers.TriggerI
	RHelper helpers.ReceiptI
	CHelper helpers.ContentI
}

/*NewTrigger constructs and returns reference to a Trigger handler*/
func NewTrigger(sql gateways.SQL, sendgrid gateways.SendgridI, towncenter gateways.TownCenterI, rabbit gateways.RabbitI) *Trigger {
	return &Trigger{
		Helper:  helpers.NewTrigger(sql),
		RHelper: helpers.NewReceipt(sql, sendgrid, towncenter, rabbit),
		CHelper: helpers.NewContent(sql),
	}
}

/*New creates a new trigger entity based on the given data */
func (t *Trigger) New(ctx *gin.Context) {
	var json models.Trigger
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse("Invalid Trigger Object"))
		return
	}

	trigger := models.NewTrigger(json.ContentID, json.Key, json.Values)
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

	trigger, err := t.Helper.GetByKey(key)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": trigger})
}

/*Update overwrites the trigger entity with new values provided*/
func (t *Trigger) Update(ctx *gin.Context) {
	key := ctx.Param("key")

	var json models.Trigger
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(400, errResponse(err.Error()))
		return
	}

	err = t.Helper.Update(key, json.ContentID, json.Values)
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

	err := t.Helper.Delete(key)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, empty())
}

/*Activate starts a trigger's action*/
func (t *Trigger) Activate(ctx *gin.Context) {
	key := ctx.Param("key")

	var json models.Receipt
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(400, errResponse(err.Error()))
		return
	}

	trigger, err := t.Helper.GetByKey(key)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	if trigger == nil {
		ctx.JSON(400, errResponse("no trigger found"))
		return
	}

	content, err := t.CHelper.GetByID(trigger.ContentID.String())
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	for k, v := range trigger.Values {
		json.Values[k] = v
	}

	receipt := models.NewReceipt(json.Values, trigger.ContentID, json.UserID)
	resolved, err := content.ResolveText(receipt.Values)
	if err != nil {
		ctx.JSON(400, errResponse(err.Error()))
		return
	}

	request := &models.SendRequest{Receipt: receipt, Subject: content.Subject, Text: resolved}
	t.RHelper.Send(request)

	ctx.JSON(200, gin.H{"data": request})
}
