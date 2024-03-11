package service

import (
	"github.com/google/uuid"
	"github.com/rodsil01/solution/baggage/domain/model"
	"github.com/rodsil01/solution/baggage/domain/repository"
	"github.com/rodsil01/solution/baggage/exceptions"
	"github.com/rodsil01/solution/contracts"
)

type BaggageServiceImpl struct {
	repository repository.BaggageRepository
}

func (s *BaggageServiceImpl) GetBaggageByReservationId(reservationId string) ([]contracts.BaggageDto, *exceptions.AppError) {
	result, err := s.repository.FindByReservationId(reservationId)

	if err != nil {
		return nil, err
	}

	baggageArr := make([]contracts.BaggageDto, 0)

	for _, obj := range result {
		baggage := convertToDto(obj)
		baggageArr = append(baggageArr, *baggage)
	}

	return baggageArr, nil
}

func (s *BaggageServiceImpl) CreateBaggage(newBaggageRequest *contracts.NewBaggageRequest) (*model.Baggage, *exceptions.AppError) {
	baggage := model.Baggage{
		Id:            uuid.New().String(),
		ReservationId: newBaggageRequest.ReservationId,
		Description:   newBaggageRequest.Description,
		Weight:        newBaggageRequest.Weight,
	}

	newReservation, err := s.repository.Create(baggage)

	if err != nil {
		return nil, err
	}

	return newReservation, nil
}

func (s *BaggageServiceImpl) CreateBaggageRange(newBaggageRequest []*contracts.NewBaggageRequest) ([]string, *exceptions.AppError) {
	var baggage []model.Baggage

	for _, r := range newBaggageRequest {
		b := model.Baggage{
			Id:            uuid.New().String(),
			ReservationId: r.ReservationId,
			Description:   r.Description,
			Weight:        r.Weight,
		}
		baggage = append(baggage, b)
	}

	ids, err := s.repository.CreateRange(baggage)

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (s *BaggageServiceImpl) DeleteBaggage(baggageId string) *exceptions.AppError {
	err := s.repository.Delete(baggageId)

	if err != nil {
		return err
	}

	return nil
}

func (s *BaggageServiceImpl) DeleteBaggageRange(baggageIds []string) *exceptions.AppError {
	err := s.repository.DeleteRange(baggageIds)

	if err != nil {
		return err
	}

	return nil
}

func convertToDto(baggage model.Baggage) *contracts.BaggageDto {
	return &contracts.BaggageDto{
		Id:          baggage.Id,
		Description: baggage.Description,
		Weight:      baggage.Weight,
	}
}

func NewBaggageServiceImpl(repository repository.BaggageRepository) *BaggageServiceImpl {
	return &BaggageServiceImpl{repository: repository}
}
