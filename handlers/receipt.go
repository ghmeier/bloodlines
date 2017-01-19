package handlers

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

/*ReceiptI describes the Receipt interface interactions */
type ReceiptI interface {
	Send(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
}

/*Receipt implements ReceiptI for the receipt router*/
type Receipt struct {
	*BaseHandler
	Helper  helpers.ReceiptI
	CHelper helpers.ContentI
}

/*NewReceipt constructs and returns a new receipt handler*/
func NewReceipt(ctx *GatewayContext) *Receipt {
	stats := ctx.Stats.Clone(statsd.Prefix("api.receipt"))
	return &Receipt{
		Helper:      helpers.NewReceipt(ctx.Sql, ctx.Sendgrid, ctx.TownCenter, ctx.Rabbit),
		CHelper:     helpers.NewContent(ctx.Sql),
		BaseHandler: NewBaseHandler(stats),
	}
}

/*Send creates a message based on the receipt provided*/
func (r *Receipt) Send(ctx *gin.Context) {
	var json models.Receipt
	err := ctx.BindJSON(&json)
	if err != nil {
		r.UserError(ctx, "Error: unable to parse json", err)
		return
	}

	receipt := models.NewReceipt(json.Values, json.ContentID, json.UserID)

	err = r.Helper.Insert(receipt)
	if err != nil {
		r.ServerError(ctx, err, json)
		return
	}

	content, err := r.CHelper.GetByID(receipt.ContentID.String())
	if err != nil {
		r.ServerError(ctx, err, receipt)
		return
	}

	request := &models.SendRequest{
		ReceiptID: receipt.ID,
		ContentID: content.ID,
	}
	err = r.Helper.Send(request)
	if err != nil {
		r.ServerError(ctx, err, request)
		return
	}

	r.Success(ctx, receipt)
}

/*ViewAll returns a list of Receipt entities starting at offset up to limit*/
func (r *Receipt) ViewAll(ctx *gin.Context) {
	offset, limit := r.GetPaging(ctx)
	receipts, err := r.Helper.GetAll(offset, limit)
	if err != nil {
		r.ServerError(ctx, err, nil)
		return
	}

	r.Success(ctx, receipts)
}

/*View returns the receipt with the given id*/
func (r *Receipt) View(ctx *gin.Context) {
	id := ctx.Param("receiptId")

	receipt, err := r.Helper.GetByID(id)
	if err != nil {
		r.ServerError(ctx, err, nil)
		return
	}

	r.Success(ctx, receipt)
}
