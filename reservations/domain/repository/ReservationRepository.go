package repository

import (
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type ReservationRepository interface {
	FindAll() ([]model.Reservation, *exceptions.AppError)
	FindById(string) (*model.Reservation, *exceptions.AppError)
	Create(reservation model.Reservation) (*model.Reservation, *exceptions.AppError)
	Delete(reservationId string) *exceptions.AppError
}
