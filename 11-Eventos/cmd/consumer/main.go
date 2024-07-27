package main

import (
	"fmt"
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

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		err := msg.Ack(false)
		if err != nil {
			panic(err)
		}
	}
}
