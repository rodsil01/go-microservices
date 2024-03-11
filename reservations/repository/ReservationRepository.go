package repository

import (
	"database/sql"
	"log"

	ctx "github.com/rodsil01/solution/reservations/domain/context"
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type ReservationRepositoryImpl struct {
	context *ctx.AppDbContext
}

func (r *ReservationRepositoryImpl) FindAll() ([]model.Reservation, *exceptions.AppError) {
	query := "select id, route_id, client_id, reservation_date, seats, state from reservations"
	rows, err := r.context.Client.Query(query)

	if err != nil {
		log.Println("Error while querying reservations table: " + err.Error())
		return nil, exceptions.NewUnexpectedError("Unexpected error")
	}

	reservations := make([]model.Reservation, 0)

	for rows.Next() {
		var reservation model.Reservation
		err := rows.Scan(&reservation.Id, &reservation.RouteId, &reservation.ClientId, &reservation.ReservationDate, &reservation.Seats, &reservation.State)

		if err != nil {
			log.Println("Error while scanning reservations: " + err.Error())
			return nil, exceptions.NewUnexpectedError("Unexpected error")
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}

func (r *ReservationRepositoryImpl) FindById(id string) (*model.Reservation, *exceptions.AppError) {
	query := "select id, route_id, client_id, reservation_date, seats, state from reservations where id = ?"

	row := r.context.Client.QueryRow(query, id)

	var reservation model.Reservation
	err := row.Scan(&reservation.Id, &reservation.RouteId, &reservation.ClientId, &reservation.ReservationDate, &reservation.Seats, &reservation.Seats)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, exceptions.NewUnexpectedError("Unexpected error")
		}
	}

	return &reservation, nil
}

func (r *ReservationRepositoryImpl) Create(reservartion model.Reservation) (*model.Reservation, *exceptions.AppError) {
	query := "insert into reservations (id, route_id, client_id, reservation_date, seats, state) values (?, ?, ?, ?, ?, ?)"

	// return nil, exceptions.NewUnexpectedError("Error")

	_, err := r.context.Client.Exec(query, reservartion.Id, reservartion.RouteId, reservartion.ClientId, reservartion.ReservationDate, reservartion.Seats, reservartion.State)

	if err != nil {
		log.Println("Error while creating new reservations: " + err.Error())
		return nil, exceptions.NewUnexpectedError("Unexpected error")
	}

	return &reservartion, nil
}

func (r *ReservationRepositoryImpl) Delete(reservartionId string) *exceptions.AppError {
	query := "delete from reservations where id = ?"

	_, err := r.context.Client.Exec(query, reservartionId)

	if err != nil {
		log.Println("Error while deleting reservation: " + err.Error())
		return exceptions.NewUnexpectedError("Unexpected error")
	}

	return nil
}

func NewReservationRepositoryImpl(context *ctx.AppDbContext) *ReservationRepositoryImpl {
	return &ReservationRepositoryImpl{context: context}
}
