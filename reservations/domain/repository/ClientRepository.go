package repository

import (
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type ClientRepository interface {
	FindAll() ([]model.Client, *exceptions.AppError)
	FindById(string) (*model.Client, *exceptions.AppError)
}
