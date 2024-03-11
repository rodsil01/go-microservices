package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodsil01/solution/contracts"
	"github.com/rodsil01/solution/messagebroker"
	"github.com/rodsil01/solution/reservations/domain/service"
	"github.com/rodsil01/solution/reservations/dto"
)

type ReservationHandler struct {
	service      service.ReservationService
	eventManager messagebroker.EventManager
}

func (ch *ReservationHandler) GetAllClients(w http.ResponseWriter, r *http.Request) {
	reservations, err := ch.service.GetAllReservations()
	createResponse(w, reservations, err)
}

func (ch *ReservationHandler) GetReservationById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["reservation_id"]

	reservation, err := ch.service.GetReservationById(id)

	createResponse(w, reservation, err)
}

func (ch *ReservationHandler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var request contracts.ReservationStarted
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err := ch.eventManager.PublishEvent(contracts.RESERVATION_STARTED, request)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dto.MessageDto{
				Message: "The reservation proccess has been initiated",
			})
		}
	}
}

func NewReservationHandler(service service.ReservationService, eventManager messagebroker.EventManager) *ReservationHandler {
	return &ReservationHandler{service: service, eventManager: eventManager}
}
