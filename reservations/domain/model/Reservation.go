package model

import "time"

type Reservation struct {
	Id              string    `json:"id"`
	RouteId         string    `json:"routeId"`
	ClientId        string    `json:"clientId"`
	ReservationDate time.Time `json:"reservationDate"`
	Seats           int       `json:"seats"`
	State           int       `json:"state"`
}
