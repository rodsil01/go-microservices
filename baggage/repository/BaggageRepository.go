package repository

import (
	"log"

	ctx "github.com/rodsil01/solution/baggage/domain/context"
	"github.com/rodsil01/solution/baggage/domain/model"
	"github.com/rodsil01/solution/baggage/exceptions"
)

type BaggageRepositoryImpl struct {
	context *ctx.AppDbContext
}

func (r *BaggageRepositoryImpl) FindByReservationId(reservationId string) ([]model.Baggage, *exceptions.AppError) {
	query := "select id, reservation_id, description, weight from baggage where reservation_id = ?"
	rows, err := r.context.Client.Query(query, reservationId)

	if err != nil {
		log.Println("Error while querying baggage table: " + err.Error())
		return nil, exceptions.NewUnexpectedError("Unexpected error")
	}

	baggageArr := make([]model.Baggage, 0)

	for rows.Next() {
		var baggage model.Baggage
		err := rows.Scan(&baggage.Id, &baggage.ReservationId, &baggage.Description, &baggage.Weight)

		if err != nil {
			log.Println("Error while scanning baggage: " + err.Error())
			return nil, exceptions.NewUnexpectedError("Unexpected error")
		}

		baggageArr = append(baggageArr, baggage)
	}

	return baggageArr, nil
}

func (r *BaggageRepositoryImpl) Create(baggage model.Baggage) (*model.Baggage, *exceptions.AppError) {
	query := "insert into baggage (id, reservation_id, description, weight) values (?, ?, ?, ?)"

	_, err := r.context.Client.Exec(query, baggage.Id, baggage.ReservationId, baggage.Description, baggage.Weight)

	if err != nil {
		log.Println("Error while creating new baggage: " + err.Error())
		return nil, exceptions.NewUnexpectedError("Unexpected error")
	}

	return &baggage, nil
}

func (r *BaggageRepositoryImpl) CreateRange(baggage []model.Baggage) ([]string, *exceptions.AppError) {
	query := "insert into baggage (id, reservation_id, description, weight) values (?, ?, ?, ?)"

	// return nil, exceptions.NewUnexpectedError("Error")

	var ids []string

	for _, b := range baggage {
		_, err := r.context.Client.Exec(query, b.Id, b.ReservationId, b.Description, b.Weight)

		if err != nil {
			log.Println("Error while creating new baggage: " + err.Error())
			return nil, exceptions.NewUnexpectedError("Unexpected error")
		}

		ids = append(ids, b.Id)
	}

	return ids, nil
}

func (r *BaggageRepositoryImpl) Delete(baggageId string) *exceptions.AppError {
	query := "delete from baggage where id = ?"

	_, err := r.context.Client.Exec(query, baggageId)

	if err != nil {
		log.Println("Error while deleting baggage: " + err.Error())
		return exceptions.NewUnexpectedError("Unexpected error")
	}

	return nil
}

func (r *BaggageRepositoryImpl) DeleteRange(baggageIds []string) *exceptions.AppError {
	query := "delete from baggage where id = ?"

	for _, id := range baggageIds {
		_, err := r.context.Client.Exec(query, id)

		if err != nil {
			log.Println("Error while deleting baggage: " + err.Error())
			return exceptions.NewUnexpectedError("Unexpected error")
		}
	}

	return nil
}

func NewBaggageRepositoryImpl(context *ctx.AppDbContext) *BaggageRepositoryImpl {
	return &BaggageRepositoryImpl{context: context}
}
