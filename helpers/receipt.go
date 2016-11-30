package helpers

import (
	"strings"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

/*Receipt helps with managing receipt entities and fetching them*/
type Receipt struct {
	*baseHelper
}

/*NewReceipt constructs and returns a receipt helper*/
func NewReceipt(sql gateways.SQL) *Receipt {
	return &Receipt{baseHelper: &baseHelper{sql: sql}}
}

/*Insert takes a reciept model and appends it to the entity*/
func (r *Receipt) Insert(receipt *models.Receipt) error {
	err := r.sql.Modify(
		"INSERT INTO receipt (id, ts, vals, sendState, contentId) VALUES (?, ?, ?, ?, ?)",
		receipt.ID,
		receipt.Created,
		strings.Join(receipt.Values, ","),
		string(receipt.SendState),
		receipt.ContentID,
	)
	return err
}

/*GetReceipts returns a list of receipts of length <limit> starting at <offset>*/
func (r *Receipt) GetReceipts(offset int, limit int) ([]*models.Receipt, error) {
	rows, err := r.sql.Select("SELECT id, ts, vals, sendState, contentId FROM receipt ORDER BY id ASC LIMIT ?,? ", offset, limit)
	if err != nil {
		return nil, err
	}

	receipts, err := models.ReceiptFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return receipts, nil
}

/*GetReceiptByID returns the receipt entitiy with the given id*/
func (r *Receipt) GetReceiptByID(id string) (*models.Receipt, error) {
	rows, err := r.sql.Select("SELECT id, ts, vals, sendState, contentId FROM receipt WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	receipts, err := models.ReceiptFromSQL(rows)
	if err != nil {
		return nil, err
	}
	return receipts[0], nil
}

/*SetSendState updates the status of the receipt with the given id*/
func (r *Receipt) SetSendState(id uuid.UUID, state models.Status) error {
	err := r.sql.Modify("UPDATE receipt SET sendState=? where id=?", string(state), id)
	return err
}
