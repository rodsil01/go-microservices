package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodsil01/solution/reservations/domain/service"
)

type ClientHandler struct {
	service service.ClientService
}

func (ch *ClientHandler) GetAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := ch.service.GetAllClients()
	createResponse(w, clients, err)
}

func (ch *ClientHandler) GetClientById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["client_id"]

	client, err := ch.service.GetClientById(id)

	createResponse(w, client, err)
}

func NewClientHandler(service service.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}
