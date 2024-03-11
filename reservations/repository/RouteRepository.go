package repository

import (
	"database/sql"
	"log"

	ctx "github.com/rodsil01/solution/reservations/domain/context"
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type RouteRepositoryImpl struct {
	context *ctx.AppDbContext
}

func (r *RouteRepositoryImpl) FindAll() ([]model.Route, *exceptions.AppError) {
	query := "select id, source, destination, departure_date, arrival_date, available_seats, price from routes"
	rows, err := r.context.Client.Query(query)

	if err != nil {
		log.Println("Error while querying routes table: " + err.Error())
		return nil, exceptions.NewUnexpectedError("Unexpected error")
	}

	routes := make([]model.Route, 0)

	for rows.Next() {
		var route model.Route
		err := rows.Scan(&route.Id, &route.Source, &route.Destination, &route.DepartureDate, &route.ArrivalDate, &route.AvailableSeats, &route.Price)

		if err != nil {
			log.Println("Error while scanning routes: " + err.Error())
			return nil, exceptions.NewUnexpectedError("Unexpected error")
		}

		routes = append(routes, route)
	}

	return routes, nil
}

func (r *RouteRepositoryImpl) FindById(id string) (*model.Route, *exceptions.AppError) {
	query := "select id, source, destination, departure_date, arrival_date, available_seats, price from routes where id = ?"

	row := r.context.Client.QueryRow(query, id)

	var route model.Route
	err := row.Scan(&route.Id, &route.Source, &route.Destination, &route.DepartureDate, &route.ArrivalDate, &route.AvailableSeats, &route.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, exceptions.NewUnexpectedError("Unexpected error")
		}
	}

	return &route, nil
}

func NewRouteRepositoryImpl(context *ctx.AppDbContext) *RouteRepositoryImpl {
	return &RouteRepositoryImpl{context: context}
}
