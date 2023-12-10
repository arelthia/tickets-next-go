package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tickets-next-go/api/data"
	"tickets-next-go/api/model"

	"github.com/go-chi/chi/v5"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
	dbRunner := data.GetDb()
	tickets, err := dbRunner.GetTickets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]string{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)

}

func GetTicket(w http.ResponseWriter, r *http.Request) {
	dbRunner := data.GetDb()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]string{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	ticket, err := dbRunner.GetTicket(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]string{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)

}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	dbRunner := data.GetDb()
	var ticket model.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)
	result, err := dbRunner.CreateTicket(ticket)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]string{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	res := map[string]int64{
		"created": result,
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	dbRunner := data.GetDb()
	ticketId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]string{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	var ticket model.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)
	result, err2 := dbRunner.UpdateTicket(ticketId, ticket)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]string{
			"error": err2.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	res := map[string]int64{
		"updated": result,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	dbRunner := data.GetDb()
	ticketId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]string{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	result, err2 := dbRunner.DeleteTicket(ticketId)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]string{
			"error": err2.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	res := map[string]int64{
		"deleted": result,
	}
	json.NewEncoder(w).Encode(res)
}
