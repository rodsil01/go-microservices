package service

import (
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/domain/repository"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type ClientServiceImpl struct {
	repository repository.ClientRepository
}

func (s *ClientServiceImpl) GetAllClients() ([]model.Client, *exceptions.AppError) {
	return s.repository.FindAll()
}

func (s *ClientServiceImpl) GetClientById(id string) (*model.Client, *exceptions.AppError) {
	client, err := s.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, exceptions.NewNotFoundError("Client not found")
	}

	return client, nil
}

func NewClientServiceImpl(repository repository.ClientRepository) *ClientServiceImpl {
	return &ClientServiceImpl{repository: repository}
}
