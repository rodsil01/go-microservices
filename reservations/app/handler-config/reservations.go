package handlerconfig

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodsil01/solution/messagebroker"
	ctx "github.com/rodsil01/solution/reservations/domain/context"
	eventhandlers "github.com/rodsil01/solution/reservations/event-handlers"
	handlers "github.com/rodsil01/solution/reservations/http-handlers"
	"github.com/rodsil01/solution/reservations/repository"
	"github.com/rodsil01/solution/reservations/service"
)

func (config *HandlerConfig) AddReservationHandlers(router *mux.Router) {
	context := ctx.GetDbContext()
	eventManager := messagebroker.GetEventManager()

	reservationService := service.NewReservationServiceImpl(repository.NewReservationRepositoryImpl(context))

	clientHandler := handlers.NewReservationHandler(reservationService, eventManager)

	router.HandleFunc("/reservations", clientHandler.GetAllClients).Methods(http.MethodGet)
	router.HandleFunc("/reservations/{reservation_id}", clientHandler.GetReservationById).Methods(http.MethodGet)

	router.HandleFunc("/reservations", clientHandler.CreateReservation).Methods(http.MethodPost)

	eventManager.InjectHandlers(
		eventhandlers.NewReservationStartedHandler(reservationService, eventManager),
		eventhandlers.NewBaggageReservationFailedHandler(reservationService, eventManager),
	)
}
