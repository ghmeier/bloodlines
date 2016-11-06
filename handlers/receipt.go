package handlers

import (
	"fmt"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/models"
	"github.com/ghmeier/bloodlines/gateways"
)

type ReceiptIfc interface {
	Send(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
}

type Receipt struct {
	sql *gateways.Sql
}

func NewReceipt(sql *gateways.Sql) ReceiptIfc {
	return &Receipt{sql: sql}
}

func (r *Receipt) Send(ctx *gin.Context) {
	var json models.Receipt
	err := ctx.BindJSON(&json)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, errResponse("Invalid vals or contentId"))
		return
	}

	receipt := models.NewReceipt(json.Values, json.ContentId)
	err = r.sql.Modify(
		"INSERT INTO receipt (id, ts, vals, sendState, contentId) VALUES (?, ?, ?, ?, ?)",
		receipt.Id,
		receipt.Created,
		strings.Join(receipt.Values,","),
		receipt.SendState,
		receipt.ContentId,
	)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"Data": receipt})
}

func (r *Receipt) ViewAll(ctx *gin.Context) {
	rows, err := r.sql.Select("Select id, ts, vals, sendState, contentId from receipt")
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	receipts, err := models.ReceiptFromSql(rows)
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

	receipt, err := r.GetReceiptById(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
	}

	ctx.JSON(200, gin.H{"data":receipt})
}


func (r *Receipt) GetReceiptById(id string) (*models.Receipt, error) {
	rows, err := r.sql.Select("SELECT id, ts, vals, sendState, contentId FROM receipt WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	receipts, err := models.ReceiptFromSql(rows)
	if err != nil {
		return nil, err
	}
	return receipts[0], nil
}

func (r *Receipt) SetSendState(id string, state models.Status) error {
	err := r.sql.Modify("UPDATE receipt SET sendState=? where id=?", state, id)
	return err
}