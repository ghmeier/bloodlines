package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/models"
)

type ReceiptI interface {
	GetAll(int, int) ([]*models.Receipt, error)
	GetByID(string) (*models.Receipt, error)
	Insert(*models.Receipt) error
	SetStatus(uuid.UUID, models.Status) error
	Send(*models.SendRequest) error
	Consume() error
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

	go helper.Consume()
	// if err != nil {
	// 	fmt.Println("ERROR: unable to start consumer")
	// 	fmt.Println(err.Error())
	// 	return nil
	// }

	return helper
}

/*Insert takes a receipt model and appends it to the entity*/
func (r *Receipt) Insert(receipt *models.Receipt) error {
	err := r.sql.Modify(
		"INSERT INTO receipt (id, ts, vals, sendState, contentId, userId) VALUES (?, ?, ?, ?, ?, ?)",
		receipt.ID,
		receipt.Created,
		receipt.SerializeValues(),
		string(receipt.SendState),
		receipt.ContentID,
		receipt.UserID,
	)
	return err
}

/*GetReceipts returns a list of receipts of length <limit> starting at <offset>*/
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

/*GetReceiptByID returns the receipt entitiy with the given id*/
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

/*SetSendState updates the status of the receipt with the given id*/
func (r *Receipt) SetStatus(id uuid.UUID, state models.Status) error {
	err := r.sql.Modify("UPDATE receipt SET sendState=? where id=?", string(state), id)
	return err
}

/*Send attempts to send the text to the recipient in the receipt*/
//receipt *models.Receipt, subject string, text string
func (r *Receipt) Send(request *models.SendRequest) error {
	return r.RB.Produce(request)
}

func (r *Receipt) HandleRequest(request *models.SendRequest) error {
	target, _ := r.TC.GetUser(request.Receipt.UserID)
	err := r.SG.SendEmail(target, request.Subject, request.Text)
	return err
}

/*Consume starts a channel that consumes messages from the front of the queue*/
func (r *Receipt) Consume() error {

	msgs, err := r.RB.Consume()
	if err != nil {
		return err

	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s", d.Body)
			var request models.SendRequest
			err := json.Unmarshal(d.Body, &request)
			if err != nil {
				fmt.Println("ERROR: unable to unmarshal body")
				// Resend message??
			}

			err = r.HandleRequest(&request)
			if err != nil {
				fmt.Println("ERROR: unable to complete request")
				fmt.Println(err.Error())
				// resent receipt?
			}
		}
	}()

	<-forever
	return nil
}
