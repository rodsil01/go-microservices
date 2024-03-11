package model

import "time"

type Route struct {
	Id             string    `json:"id"`
	Source         string    `json:"source"`
	Destination    string    `json:"destination"`
	DepartureDate  time.Time `json:"departureDate"`
	ArrivalDate    time.Time `json:"arrivalDate"`
	AvailableSeats int       `json:"availableSeats"`
	Price          float32   `json:"price"`
}
