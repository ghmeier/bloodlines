package gateways

import (
	"encoding/json"
	"fmt"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/models"

	"github.com/streadway/amqp"
)

/*RabbitI describes the gateway for interacting with RabbitMQ*/
type RabbitI interface {
	Produce(*models.SendRequest) error
	Consume() (<-chan amqp.Delivery, error)
	Destroy()
}

/*Rabbit stores references to RabbitMQ connections, channels, queues, and config*/
type Rabbit struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
	Config  config.Rabbit
}

/*NewRabbit creates and returns a new Rabbit struct*/
func NewRabbit(config config.Rabbit) (*Rabbit, error) {
	r := &Rabbit{Config: config}

	err := r.connect()
	return r, err
}

/*Produce adds a message to the queue based on a send request*/
func (r *Rabbit) Produce(request *models.SendRequest) error {
	err := r.connect()
	if err != nil {
		return err
	}

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

/*Consume starts a worker that returns a channel from the queue to read from*/
func (r *Rabbit) Consume() (<-chan amqp.Delivery, error) {
	err := r.connect()
	if err != nil {
		return nil, err
	}

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

/*Destroy ends the rabbitmq connection*/
func (r *Rabbit) Destroy() {
	r.Channel.Close()
	r.Conn.Close()
}

func (r *Rabbit) connect() error {
	if r.Conn != nil && r.Channel != nil && r.Queue != nil {
		return nil
	}
	conn, err := amqp.Dial("amqp://guest:guest@" + r.Config.Host + ":" + r.Config.Port + "/")
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		r.Config.PubQ, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	r.Conn = conn
	r.Channel = ch
	r.Queue = &q
	return nil
}
