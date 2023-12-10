package main

// import chi
import (
	"net/http"

	"tickets-next-go/api/controller"
	"tickets-next-go/api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	corsHandler := cors.New(middleware.Cors())
	r.Use(corsHandler.Handler)
	controller.Tickets(r)
	http.ListenAndServe(":8000", r)
}
