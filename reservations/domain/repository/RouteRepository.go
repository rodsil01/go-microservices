package repository

import (
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type RouteRepository interface {
	FindAll() ([]model.Route, *exceptions.AppError)
	FindById(string) (*model.Route, *exceptions.AppError)
}
