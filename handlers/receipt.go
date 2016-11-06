package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/models"
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/helpers"
)

type ReceiptIfc interface {
	Send(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
}

type Receipt struct {
	helper *helpers.Receipt
}

func NewReceipt(sql *gateways.Sql) ReceiptIfc {
	return &Receipt{helper: helpers.NewReceipt(sql)}
}

func (r *Receipt) Send(ctx *gin.Context) {
	var json models.Receipt
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(400, errResponse("Invalid vals or contentId"))
		return
	}

	receipt := models.NewReceipt(json.Values, json.ContentId)

	err = r.helper.Insert(receipt)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	// SEND SOMETHING

	ctx.JSON(200, gin.H{"Data": receipt})
}

func (r *Receipt) ViewAll(ctx *gin.Context) {
	receipts, err := r.helper.GetReceipts()
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": receipts})
}

func (r *Receipt) View(ctx *gin.Context) {
	id := ctx.Param("receiptId")
	if id == "" {
		ctx.JSON(400, errResponse("receiptId is a required parameter"))
		return
	}

	receipt, err := r.helper.GetReceiptById(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
	}

	ctx.JSON(200, gin.H{"data":receipt})
}
