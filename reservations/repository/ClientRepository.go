package repository

import (
	"database/sql"
	"log"

	ctx "github.com/rodsil01/solution/reservations/domain/context"
	"github.com/rodsil01/solution/reservations/domain/model"
	"github.com/rodsil01/solution/reservations/exceptions"
)

type ClientRepositoryImpl struct {
	context *ctx.AppDbContext
}

func (r *ClientRepositoryImpl) FindAll() ([]model.Client, *exceptions.AppError) {
	query := "select id, name, last_name, email from clients"
	rows, err := r.context.Client.Query(query)

	if err != nil {
		log.Println("Error while querying clients table: " + err.Error())
		return nil, exceptions.NewUnexpectedError("Unexpected error")
	}

	clients := make([]model.Client, 0)

	for rows.Next() {
		var client model.Client
		err := rows.Scan(&client.Id, &client.Name, &client.LastName, &client.Email)

		if err != nil {
			log.Println("Error while scanning clients: " + err.Error())
			return nil, exceptions.NewUnexpectedError("Unexpected error")
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func (r *ClientRepositoryImpl) FindById(id string) (*model.Client, *exceptions.AppError) {
	query := "select id, name, last_name, email from clients where id = ?"

	row := r.context.Client.QueryRow(query, id)

	var client model.Client
	err := row.Scan(&client.Id, &client.Name, &client.LastName, &client.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, exceptions.NewUnexpectedError("Unexpected error")
		}
	}

	return &client, nil
}

func NewClientRepositoryImpl(context *ctx.AppDbContext) *ClientRepositoryImpl {
	return &ClientRepositoryImpl{context: context}
}
