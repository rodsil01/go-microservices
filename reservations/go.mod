module github.com/rodsil01/solution/reservations

go 1.22.1

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/rodsil01/solution/contracts v1.0.0
	github.com/rodsil01/solution/messagebroker v1.0.0
)

require github.com/streadway/amqp v1.1.0 // indirect

replace (
	github.com/rodsil01/solution/contracts => ../contracts
	github.com/rodsil01/solution/messagebroker => ../message-broker
)
