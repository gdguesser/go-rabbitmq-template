package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// Connect establishes a connection to our RabbitMQ instance
// and declares the queue we are going to be using
func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")
	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Sucessfully connected to RabbitMQ")

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}

	return nil
}

// NewRabbitMQService returns a pointer to a new RabbitMQ service
func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
