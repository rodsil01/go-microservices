package service

import (
	"github.com/google/uuid"
	"github.com/rodsil01/solution/contracts"
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/domain/repository"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type ReservationServiceImpl struct {
	repository repository.ReservationRepository
}

func (s *ReservationServiceImpl) GetAllReservations() ([]model.Reservation, *exceptions.AppError) {
	return s.repository.FindAll()
}

func (s *ReservationServiceImpl) CreateReservation(newReservationRequest *contracts.NewReservationRequest) (*model.Reservation, *exceptions.AppError) {
	reservation := model.Reservation{
		Id:              uuid.New().String(),
		RouteId:         newReservationRequest.RouteId,
		ClientId:        newReservationRequest.ClientId,
		ReservationDate: newReservationRequest.ReservationDate,
		Seats:           newReservationRequest.Seats,
		State:           newReservationRequest.State,
	}

	newReservation, err := s.repository.Create(reservation)

	if err != nil {
		return nil, err
	}

	return newReservation, nil
}

func (s *ReservationServiceImpl) DeleteReservation(reservationId string) *exceptions.AppError {
	err := s.repository.Delete(reservationId)

	if err != nil {
		return err
	}

	return nil
}

func (s *ReservationServiceImpl) GetReservationById(id string) (*model.Reservation, *exceptions.AppError) {
	reservation, err := s.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	if reservation == nil {
		return nil, exceptions.NewNotFoundError("Reservation not found")
	}

	return reservation, nil
}

func NewReservationServiceImpl(repository repository.ReservationRepository) *ReservationServiceImpl {
	return &ReservationServiceImpl{repository: repository}
}
