package contracts

import "time"

type NewReservationRequest struct {
	RouteId         string    `json:"routeId"`
	ClientId        string    `json:"clientId"`
	ReservationDate time.Time `json:"reservationDate"`
	Seats           int       `json:"seats"`
	State           int       `json:"state"`
}

type ReservationStarted struct {
	RouteId         string                            `json:"routeId"`
	ClientId        string                            `json:"clientId"`
	ReservationDate time.Time                         `json:"reservationDate"`
	Seats           int                               `json:"seats"`
	State           int                               `json:"state"`
	Baggage         []NewBaggageRequestForReservation `json:"baggage"`
}

type ReservationCreated struct {
	ReservationId string                            `json:"reservationId"`
	Baggage       []NewBaggageRequestForReservation `json:"baggage"`
}
