package handler

import (
	"fmt"
	"net/http"

	"golang-email/producer"
)

type Notification struct{}

func (n *Notification) CreateNotificationEmail(w http.ResponseWriter, r *http.Request) {
	producer.Producer()
	fmt.Println("Email queued to be sent")
}
