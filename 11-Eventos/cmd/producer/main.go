package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/willbrr.dev/goexpert/11-Eventos/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			panic(err)
		}
	}(ch)

	rabbitmq.Publish(ch, "Hello, World!", "amq.direct")
}
