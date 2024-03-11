package model

type Baggage struct {
	Id            string `json:"id"`
	ReservationId string `json:"reservationId"`
	Description   string `json:"description"`
	Weight        string `json:"weight"`
}
