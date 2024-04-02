package producer

import (
	"context"
	"strconv"

	"github.com/rabbitmq/amqp091-go"
)

func Producer() {
	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	for i := 0; i < 3; i++ {

		message := amqp091.Publishing{
			Headers: amqp091.Table{
				"key": "value",
			},
			Body: []byte("Hello World " + strconv.Itoa(i)),
		}
		err := channel.PublishWithContext(ctx, "notification", "email", false, false, message)
		if err != nil {
			panic(err)
		}
	}
}
