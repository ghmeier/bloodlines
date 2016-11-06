package helpers

import (
	"strings"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/models"
	"github.com/ghmeier/bloodlines/gateways"
)

type Receipt struct{
	*baseHelper
}

func NewReceipt(sql *gateways.Sql) *Receipt {
	return &Receipt{baseHelper: &baseHelper{sql: sql}}
}

func (r *Receipt) Insert(receipt *models.Receipt) error {
	err := r.sql.Modify(
		"INSERT INTO receipt (id, ts, vals, sendState, contentId) VALUES (?, ?, ?, ?, ?)",
		receipt.Id,
		receipt.Created,
		strings.Join(receipt.Values,","),
		receipt.SendState,
		receipt.ContentId,
	)
	return err
}

func (r *Receipt) GetReceipts() ([]*models.Receipt, error) {
	rows, err := r.sql.Select("Select id, ts, vals, sendState, contentId from receipt")
	if err != nil {
		return nil, err
	}

	receipts, err := models.ReceiptFromSql(rows)
	if err != nil {
		return nil, err
	}

	return receipts, nil
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

func (r *Receipt) SetSendState(id uuid.UUID, state models.Status) error {
	err := r.sql.Modify("UPDATE receipt SET sendState=? where id=?", state, id)
	return err
}