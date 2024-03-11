package repository

import (
	"github.com/rodsil01/solution/baggage/domain/model"
	"github.com/rodsil01/solution/baggage/exceptions"
)

type BaggageRepository interface {
	FindByReservationId(string) ([]model.Baggage, *exceptions.AppError)
	Create(baggage model.Baggage) (*model.Baggage, *exceptions.AppError)
	CreateRange(baggage []model.Baggage) ([]string, *exceptions.AppError)
	Delete(baggageId string) *exceptions.AppError
	DeleteRange(baggageIds []string) *exceptions.AppError
}
