package handlerconfig

import (
	"net/http"

	"github.com/gorilla/mux"
	ctx "github.com/rodsil01/solution/reservations/domain/context"
	handlers "github.com/rodsil01/solution/reservations/http-handlers"
	"github.com/rodsil01/solution/reservations/repository"
	"github.com/rodsil01/solution/reservations/service"
)

func (config *HandlerConfig) AddRouteHandlers(router *mux.Router) {
	context := ctx.GetDbContext()

	clientHandler := handlers.NewRouteHandler(service.NewRouteServiceImpl(repository.NewRouteRepositoryImpl(context)))

	router.HandleFunc("/routes", clientHandler.GetAllRoutes).Methods(http.MethodGet)
	router.HandleFunc("/routes/{route_id}", clientHandler.GetRouteById).Methods(http.MethodGet)
}
