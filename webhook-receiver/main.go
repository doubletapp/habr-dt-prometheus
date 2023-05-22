package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"webhook-receiver/alerts"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/alert/{channelId}", alerts.HandleAlert)
	http.ListenAndServe(":3000", router)
}
