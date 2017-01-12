package workers

import (
	"encoding/json"
	"fmt"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

/*Send describes the functions a send request uses*/
type Send interface {
	Consume() error
}

/*SendRequest has acces to receipt and content helpers as well as rabbitmq for publishing*/
type SendRequest struct {
	RHelper helpers.ReceiptI
	CHelper helpers.ContentI
	RB      gateways.RabbitI
}

/*NewSend creates and returns a new SendRequest*/
func NewSend(sql gateways.SQL, sendgrid gateways.SendgridI, towncenter gateways.TownCenterI, rabbit gateways.RabbitI) Send {
	return &SendRequest{
		RHelper: helpers.NewReceipt(sql, sendgrid, towncenter, rabbit),
		CHelper: helpers.NewContent(sql),
		RB:      rabbit,
	}
}

func (s *SendRequest) handleRequest(request *models.SendRequest) {
	receipt, err := s.RHelper.GetByID(request.ReceiptID.String())
	if err != nil {
		fmt.Printf("ERROR: unable to get receipt: %s\n", request.ReceiptID.String())
		fmt.Println(err.Error())
		return
	}

	content, err := s.CHelper.GetByID(request.ContentID.String())
	if err != nil {
		fmt.Printf("ERROR: unable to get content %s\n", request.ContentID.String())
		fmt.Println(err.Error())
		s.RHelper.SetStatus(receipt.ID, models.FAILURE)
		return
	}

	err = s.RHelper.DeliverContent(receipt, content)
	if err != nil {
		//TODO: resend?
		fmt.Println("ERROR: unable to complete request")
		fmt.Println(err.Error())
		s.RHelper.SetStatus(receipt.ID, models.FAILURE)
		return
	}
	s.RHelper.SetStatus(receipt.ID, models.SUCCESS)
}

/*Consume starts a channel that consumes messages from the front of the queue*/
func (s *SendRequest) Consume() error {

	msgs, err := s.RB.Consume()
	if err != nil {
		return err
	}

	//forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)
			var request models.SendRequest
			err := json.Unmarshal(d.Body, &request)
			if err != nil {
				fmt.Println("ERROR: unable to unmarshal body")
				// Resend message??
			}

			s.handleRequest(&request)
		}
	}()

	//<-forever
	return nil
}
