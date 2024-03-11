package eventhandlers

import (
	"encoding/json"

	"github.com/rodsil01/solution/contracts"
	"github.com/rodsil01/solution/messagebroker"
	"github.com/rodsil01/solution/reservations/domain/service"
)

type ReservationStartedHandler struct {
	service      service.ReservationService
	eventManager messagebroker.EventManager
}

func (h *ReservationStartedHandler) HandleEvent(eventType string, event interface{}) {
	jsonData, _ := json.Marshal(event)

	var payload contracts.ReservationStarted
	marshalErr := json.Unmarshal(jsonData, &payload)

	if marshalErr != nil {
		return
	}

	reservation, err := h.service.CreateReservation(&contracts.NewReservationRequest{
		RouteId:         payload.RouteId,
		ClientId:        payload.ClientId,
		ReservationDate: payload.ReservationDate,
		Seats:           payload.Seats,
		State:           payload.State,
	})

	if err == nil {
		h.eventManager.PublishEvent(contracts.RESERVATION_CREATED, contracts.ReservationCreated{
			ReservationId: reservation.Id,
			Baggage:       payload.Baggage,
		})
	}
}

func NewReservationStartedHandler(service service.ReservationService, eventManager messagebroker.EventManager) *ReservationStartedHandler {
	return &ReservationStartedHandler{service: service, eventManager: eventManager}
}
