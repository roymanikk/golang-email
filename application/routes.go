package application

import (
	"net/http"

	"golang-email/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.Route("/notification", loadNotificationRoutes)

	return router
}

func loadNotificationRoutes(router chi.Router) {
	notificationHandler := &handler.Notification{}

	router.Post("/create-email", notificationHandler.CreateNotificationEmail)
}
