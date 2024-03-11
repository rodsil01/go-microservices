package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodsil01/solution/reservations/domain/service"
)

type RouteHandler struct {
	service service.RouteService
}

func (rh *RouteHandler) GetAllRoutes(w http.ResponseWriter, r *http.Request) {
	routes, err := rh.service.GetAllRoutes()
	createResponse(w, routes, err)
}

func (ch *RouteHandler) GetRouteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["route_id"]

	route, err := ch.service.GetRouteById(id)

	createResponse(w, route, err)
}

func NewRouteHandler(service service.RouteService) *RouteHandler {
	return &RouteHandler{service: service}
}
