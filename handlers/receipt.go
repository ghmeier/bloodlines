package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
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
	Helper  helpers.ReceiptI
	CHelper helpers.ContentI
}

/*NewReceipt constructs and returns a new receipt handler*/
func NewReceipt(sql gateways.SQL, sendgrid gateways.SendgridI, towncenter gateways.TownCenterI, rabbit gateways.RabbitI) *Receipt {
	return &Receipt{
		Helper:  helpers.NewReceipt(sql, sendgrid, towncenter, rabbit),
		CHelper: helpers.NewContent(sql),
	}
}

/*Send creates a message based on the receipt provided*/
func (r *Receipt) Send(ctx *gin.Context) {
	var json models.Receipt
	err := ctx.BindJSON(&json)
	if err != nil {
		UserError(ctx, "Error: unable to parse json", err)
		return
	}

	receipt := models.NewReceipt(json.Values, json.ContentID, json.UserID)

	err = r.Helper.Insert(receipt)
	if err != nil {
		ServerError(ctx, err, json)
		return
	}

	content, err := r.CHelper.GetByID(receipt.ContentID.String())
	if err != nil {
		ServerError(ctx, err, receipt)
		return
	}

	request := &models.SendRequest{
		ReceiptID: receipt.ID,
		ContentID: content.ID,
	}
	err = r.Helper.Send(request)
	if err != nil {
		ServerError(ctx, err, request)
		return
	}

	Success(ctx, receipt)
}

/*ViewAll returns a list of Receipt entities starting at offset up to limit*/
func (r *Receipt) ViewAll(ctx *gin.Context) {
	offset, limit := GetPaging(ctx)
	receipts, err := r.Helper.GetAll(offset, limit)
	if err != nil {
		ServerError(ctx, err, nil)
		return
	}

	Success(ctx, receipts)
}

/*View returns the receipt with the given id*/
func (r *Receipt) View(ctx *gin.Context) {
	id := ctx.Param("receiptId")

	receipt, err := r.Helper.GetByID(id)
	if err != nil {
		ServerError(ctx, err, nil)
		return
	}

	Success(ctx, receipt)
}
