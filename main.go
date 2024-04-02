package main

import (
	"context"
	"fmt"

	"golang-email/application"
	"golang-email/consumer"
)

func main() {
	go consumer.Consumer()
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}
