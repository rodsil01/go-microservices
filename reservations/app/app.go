package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rodsil01/solution/messagebroker"
	eventconfig "github.com/rodsil01/solution/reservations/app/event-config"
	handlerconfig "github.com/rodsil01/solution/reservations/app/handler-config"
)

func Start() {
	sanityCheck()

	eventsConfig := eventconfig.EventConfig{}

	rabbitMQUsername := os.Getenv("RABBIT_MQ_USERNAME")
	rabbitMQPassword := os.Getenv("RABBIT_MQ_PASSWORD")
	rabbitMQHost := os.Getenv("RABBIT_MQ_HOST")
	rabbitMQPort := os.Getenv("RABBIT_MQ_PORT")
	rabbitMQVHost := os.Getenv("RABBIT_MQ_VHOST")

	eventsConfig.AddRabbitMQ(fmt.Sprintf("amqp://%s:%s@%s:%s/%s", rabbitMQUsername, rabbitMQPassword, rabbitMQHost, rabbitMQPort, rabbitMQVHost))
	defer messagebroker.GetEventManager().Close()

	handlerConfig := handlerconfig.HandlerConfig{}

	router := mux.NewRouter()

	handlerConfig.AddClientHandlers(router)
	handlerConfig.AddRouteHandlers(router)
	handlerConfig.AddReservationHandlers(router)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func sanityCheck() {
	message := ""

	if os.Getenv("SERVER_ADDRESS") == "" {
		message = message + "\n" + "Environment variable 'SERVER_ADDRESS' must be specified"
	}
	if os.Getenv("SERVER_PORT") == "" {
		message = message + "\n" + "Environment variable 'SERVER_PORT' must be specified"
	}
	if os.Getenv("DB_USER") == "" {
		message = message + "\n" + "Environment variable 'DB_USER' must be specified"
	}
	if os.Getenv("DB_PASSWORD") == "" {
		message = message + "\n" + "Environment variable 'DB_PASSWORD' must be specified"
	}
	if os.Getenv("DB_ADDRESS") == "" {
		message = message + "\n" + "Environment variable 'DB_ADDRESS' must be specified"
	}
	if os.Getenv("DB_PORT") == "" {
		message = message + "\n" + "Environment variable 'DB_PORT' must be specified"
	}
	if os.Getenv("DB_NAME") == "" {
		message = message + "\n" + "Environment variable 'DB_NAME' must be specified"
	}
	if os.Getenv("RABBIT_MQ_USERNAME") == "" {
		message = message + "\n" + "Environment variable 'RABBIT_MQ_USERNAME' must be specified"
	}
	if os.Getenv("RABBIT_MQ_PASSWORD") == "" {
		message = message + "\n" + "Environment variable 'RABBIT_MQ_PASSWORD' must be specified"
	}
	if os.Getenv("RABBIT_MQ_HOST") == "" {
		message = message + "\n" + "Environment variable 'RABBIT_MQ_HOST' must be specified"
	}
	if os.Getenv("RABBIT_MQ_PORT") == "" {
		message = message + "\n" + "Environment variable 'RABBIT_MQ_PORT' must be specified"
	}

	if message != "" {
		log.Fatal(message)
	}
}
