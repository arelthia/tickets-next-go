package controller

import (
	"tickets-next-go/api/handler"

	"github.com/go-chi/chi/v5"
)

func Tickets(r chi.Router) {

	r.Route("/api/v1/tickets", func(r chi.Router) {
		// set up cors headers
		r.Get("/", handler.GetTickets)
		r.Get("/{id}", handler.GetTicket)
		r.Post("/", handler.CreateTicket)
		r.Put("/{id}", handler.UpdateTicket)
		r.Delete("/{id}", handler.DeleteTicket)
	})

}
