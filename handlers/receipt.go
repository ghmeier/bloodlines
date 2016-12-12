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
	StartConsumer()
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

	receipt := models.NewReceipt(json.Values, json.ContentID, json.UserID)

	err = r.Helper.Insert(receipt)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	content, err := r.CHelper.GetByID(receipt.ContentID.String())
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	resolved, err := content.ResolveText(receipt.Values)
	if err != nil {
		ctx.JSON(400, errResponse(err.Error()))
		return
	}

	err = r.Helper.Send(&models.SendRequest{
		Receipt: receipt,
		Subject: content.Subject,
		Text:    resolved,
	})
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": receipt})
}

/*ViewAll returns a list of Receipt entities starting at offset up to limit*/
func (r *Receipt) ViewAll(ctx *gin.Context) {
	offset, limit := getPaging(ctx)
	receipts, err := r.Helper.GetAll(offset, limit)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": receipts})
}

/*View returns the receipt with the given id*/
func (r *Receipt) View(ctx *gin.Context) {
	id := ctx.Param("receiptId")
	if id == "" {
		ctx.JSON(400, errResponse("receiptId is a required parameter"))
		return
	}

	receipt, err := r.Helper.GetByID(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
	}

	ctx.JSON(200, gin.H{"data": receipt})
}

func (r *Receipt) StartConsumer() {
	go r.Helper.Consume()
}
