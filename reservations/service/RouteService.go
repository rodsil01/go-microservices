package service

import (
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/domain/repository"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type RouteServiceImpl struct {
	repository repository.RouteRepository
}

func (s *RouteServiceImpl) GetAllRoutes() ([]model.Route, *exceptions.AppError) {
	return s.repository.FindAll()
}

func (s *RouteServiceImpl) GetRouteById(id string) (*model.Route, *exceptions.AppError) {
	route, err := s.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	if route == nil {
		return nil, exceptions.NewNotFoundError("Route not found")
	}

	return route, nil
}

func NewRouteServiceImpl(repository repository.RouteRepository) *RouteServiceImpl {
	return &RouteServiceImpl{repository: repository}
}
