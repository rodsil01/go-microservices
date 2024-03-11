package service

import (
	"github.com/rodsil01/solution/contracts"
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type ReservationService interface {
	GetAllReservations() ([]model.Reservation, *exceptions.AppError)
	GetReservationById(string) (*model.Reservation, *exceptions.AppError)
	CreateReservation(newReservationRequest *contracts.NewReservationRequest) (*model.Reservation, *exceptions.AppError)
	DeleteReservation(reservationId string) *exceptions.AppError
}
