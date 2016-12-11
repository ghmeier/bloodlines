package gateways

import (
	"encoding/json"
	"fmt"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/models"

	"github.com/streadway/amqp"
)

type RabbitI interface {
	Produce(*models.SendRequest) error
	Consume() (<-chan amqp.Delivery, error)
	Destroy()
}

type Rabbit struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func NewRabbit(config config.Rabbit) (*Rabbit, error) {
	conn, err := amqp.Dial("amqp://guest:guest@" + config.Host + ":" + config.Port + "/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		config.PubQ, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return nil, err
	}

	return &Rabbit{Conn: conn, Channel: ch, Queue: &q}, nil
}

func (r *Rabbit) Produce(request *models.SendRequest) error {
	s, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = r.Channel.Publish(
		"",           // exchange
		r.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(s),
		})
	return err
}

func (r *Rabbit) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := r.Channel.Consume(
		r.Queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return msgs, nil
}

func (r *Rabbit) Destroy() {
	r.Channel.Close()
	r.Conn.Close()
}
