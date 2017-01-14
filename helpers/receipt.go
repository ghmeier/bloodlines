package helpers

import (
	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

/*ReceiptI describes methods for the Receipt Helper*/
type ReceiptI interface {
	GetAll(int, int) ([]*models.Receipt, error)
	GetByID(string) (*models.Receipt, error)
	Insert(*models.Receipt) error
	SetStatus(uuid.UUID, models.Status) error
	Send(*models.SendRequest) error
	//Consume() error
	DeliverContent(*models.Receipt, *models.Content) error
}

/*Receipt helps with managing receipt entities and fetching them*/
type Receipt struct {
	*baseHelper
	SG gateways.SendgridI
	TC gateways.TownCenterI
	RB gateways.RabbitI
}

/*NewReceipt constructs and returns a receipt helper*/
func NewReceipt(sql gateways.SQL, sendgrid gateways.SendgridI, townCenter gateways.TownCenterI, rabbit gateways.RabbitI) ReceiptI {
	helper := &Receipt{
		baseHelper: &baseHelper{sql: sql},
		SG:         sendgrid,
		TC:         townCenter,
		RB:         rabbit,
	}
	return helper
}

/*Insert takes a receipt model and appends it to the entity*/
func (r *Receipt) Insert(receipt *models.Receipt) error {
	err := r.sql.Modify(
		"INSERT INTO receipt (id, ts, vals, sendState, contentId, userId) VALUES (?, ?, ?, ?, ?, ?)",
		receipt.ID,
		receipt.Created,
		models.SerializeValues(receipt.Values),
		string(receipt.SendState),
		receipt.ContentID,
		receipt.UserID,
	)
	return err
}

/*GetAll returns a list of receipts of length <limit> starting at <offset>*/
func (r *Receipt) GetAll(offset int, limit int) ([]*models.Receipt, error) {
	rows, err := r.sql.Select("SELECT id, ts, vals, sendState, contentId, userId FROM receipt ORDER BY id ASC LIMIT ?,? ", offset, limit)
	if err != nil {
		return nil, err
	}

	receipts, err := models.ReceiptFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return receipts, nil
}

/*GetByID returns the receipt entitiy with the given id*/
func (r *Receipt) GetByID(id string) (*models.Receipt, error) {
	rows, err := r.sql.Select("SELECT id, ts, vals, sendState, contentId, userId FROM receipt WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	receipts, err := models.ReceiptFromSQL(rows)
	if err != nil {
		return nil, err
	}
	return receipts[0], nil
}

/*SetStatus updates the status of the receipt with the given id*/
func (r *Receipt) SetStatus(id uuid.UUID, state models.Status) error {
	err := r.sql.Modify("UPDATE receipt SET sendState=? where id=?", string(state), id)
	return err
}

/*Send attempts to send the text to the recipient in the receipt*/
//receipt *models.Receipt, subject string, text string
func (r *Receipt) Send(request *models.SendRequest) error {
	err := r.RB.Produce(request)
	if err != nil {
		r.SetStatus(request.ReceiptID, models.FAILURE)
		return err
	}

	err = r.SetStatus(request.ReceiptID, models.QUEUED)
	return err
}

/*DeliverContent combines receipt and content to send a message to the given ContentType */
func (r *Receipt) DeliverContent(receipt *models.Receipt, content *models.Content) error {
	//ignoring error until TC is actually implemented
	switch content.Type {
	case models.EMAIL:
		return r.deliverEmail(receipt, content)
	default:
		return r.deliverNoop(receipt, content)
	}
}

func (r *Receipt) deliverEmail(receipt *models.Receipt, content *models.Content) error {
	target, _ := r.TC.GetUser(receipt.UserID)

	text, err := content.ResolveText(receipt.Values)
	if err != nil {
		return err
	}

	err = r.SG.SendEmail(target, content.Subject, text)
	return err
}

func (r *Receipt) deliverNoop(receipt *models.Receipt, content *models.Content) error {
	return nil
}
