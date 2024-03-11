package service

import (
	"github.com/rodsil01/solution/baggage/domain/model"
	"github.com/rodsil01/solution/baggage/exceptions"
	"github.com/rodsil01/solution/contracts"
)

type BaggageService interface {
	GetBaggageByReservationId(string) ([]contracts.BaggageDto, *exceptions.AppError)
	CreateBaggage(newBaggageRequest *contracts.NewBaggageRequest) (*model.Baggage, *exceptions.AppError)
	CreateBaggageRange(newBaggageRequest []*contracts.NewBaggageRequest) ([]string, *exceptions.AppError)
	DeleteBaggage(baggageId string) *exceptions.AppError
	DeleteBaggageRange(baggageIds []string) *exceptions.AppError
}
