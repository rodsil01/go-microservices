package contracts

type NewBaggageRequest struct {
	ReservationId string `json:"reservationId"`
	Description   string `json:"description"`
	Weight        string `json:"weight"`
}

type NewBaggageRequestForReservation struct {
	Description string `json:"description"`
	Weight      string `json:"weight"`
}

type BaggageDto struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Weight      string `json:"weight"`
}

type BaggageReservationFailed struct {
	Ids           []string `json:"ids"`
	ReservationId string   `json:"reservationId"`
}
