package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodsil01/solution/baggage/domain/service"
)

type BaggageHandler struct {
	service service.BaggageService
}

func (ch *BaggageHandler) GetBaggageByReservationId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservation_id := vars["reservation_id"]

	clients, err := ch.service.GetBaggageByReservationId(reservation_id)
	createResponse(w, clients, err)
}

func NewBaggageHandler(service service.BaggageService) *BaggageHandler {
	return &BaggageHandler{service: service}
}
