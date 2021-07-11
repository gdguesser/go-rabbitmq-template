package main

import (
	"fmt"
	"github.com/gdguesser/go-rabbitmq-template/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Go RabbitMQ Template")

	rmq := rabbitmq.NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up application")
		fmt.Println(err)
	}
}
