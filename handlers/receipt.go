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
	helper helpers.ReceiptI
}

/*NewReceipt constructs and returns a new receipt handler*/
func NewReceipt(sql gateways.SQL) *Receipt {
	return &Receipt{helper: helpers.NewReceipt(sql)}
}

/*Send creates a message based on the receipt provided*/
func (r *Receipt) Send(ctx *gin.Context) {
	var json models.Receipt
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(400, errResponse("Invalid vals or contentId"))
		return
	}

	receipt := models.NewReceipt(json.Values, json.ContentID)

	err = r.helper.Insert(receipt)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	// SEND SOMETHING

	ctx.JSON(200, gin.H{"Data": receipt})
}

/*ViewAll returns a list of Receipt entities starting at offset up to limit*/
func (r *Receipt) ViewAll(ctx *gin.Context) {
	offset, limit := getPaging(ctx)
	receipts, err := r.helper.GetAll(offset, limit)
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

	receipt, err := r.helper.GetByID(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
	}

	ctx.JSON(200, gin.H{"data": receipt})
}
