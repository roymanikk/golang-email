package consumer

import (
	"context"
	"fmt"

	"golang-email/helper"

	"github.com/rabbitmq/amqp091-go"
)

func Consumer() {
	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	emailConsumer, err := channel.ConsumeWithContext(ctx, "email", "consumer-email", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for message := range emailConsumer {
		helper.SendMailHTML("Hello, this is a notification email", "./test.html", []string{"roymanlanjutmanik@gmail.com"}, struct{ Name string }{Name: string(message.Body)})
	}
}
