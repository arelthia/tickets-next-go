package main

// import chi
import (
	"net/http"

	"tickets-next-go/api/controller"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	controller.Tickets(r)
	http.ListenAndServe(":8000", r)
}
