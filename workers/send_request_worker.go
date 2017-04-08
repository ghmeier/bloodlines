package workers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/handlers"
	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

/*Worker describes the functions a queue worker uses*/
type Worker interface {
	Consume() error
	Handle([]byte)
}

type HandleFunc func([]byte)
type BaseWorker struct {
	RB gateways.RabbitI
	HandleFunc
}

/*SendRequest has acces to receipt and content helpers as well as rabbitmq for publishing*/
type sendWorker struct {
	Receipt helpers.ReceiptI
	Content helpers.ContentI
}

/*NewSend creates and returns a new SendWorker*/
func NewSend(ctx *handlers.GatewayContext) Worker {
	worker := &sendWorker{
		Receipt: helpers.NewReceipt(ctx.Sql, ctx.Sendgrid, ctx.TownCenter, ctx.Rabbit),
		Content: helpers.NewContent(ctx.Sql),
	}

	return &BaseWorker{
		HandleFunc: HandleFunc(worker.handle),
		RB:         ctx.Rabbit,
	}
}

func (f HandleFunc) Handle(body []byte) {
	f(body)
}

func (s *sendWorker) handle(body []byte) {
	var request models.SendRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		fmt.Printf("ERROR: unable to parse body\n")
		fmt.Println(err.Error())
		return
	}

	receipt, err := s.Receipt.GetByID(request.ReceiptID.String())
	if err != nil {
		fmt.Printf("ERROR: unable to get receipt: %s\n", request.ReceiptID.String())
		fmt.Println(err.Error())
		return
	}

	content, err := s.Content.Get(request.ContentID.String())
	if err != nil {
		fmt.Printf("ERROR: unable to get content %s\n", request.ContentID.String())
		fmt.Println(err.Error())
		s.Receipt.SetStatus(receipt.ID, models.FAILURE)
		return
	}

	err = s.Receipt.DeliverContent(receipt, content)
	if err != nil {
		//TODO: resend?
		fmt.Println("ERROR: unable to complete request")
		fmt.Println(err.Error())
		s.Receipt.SetStatus(receipt.ID, models.FAILURE)
		return
	}
	s.Receipt.SetStatus(receipt.ID, models.SUCCESS)
}

/*Consume starts a channel that consumes messages from the front of the queue*/
func (w *BaseWorker) Consume() error {

	go func() {
		msgs, err := w.RB.Consume()
		for err != nil {
			fmt.Printf("Rabbit ERROR: %s\n", err.Error())
			time.Sleep(time.Duration(5) * time.Second)
			msgs, err = w.RB.Consume()
		}

		//forever := make(chan bool)

		for d := range msgs {
			w.Handle(d.Body)
		}
	}()

	//<-forever
	return nil
}
