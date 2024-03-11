module github.com/rodsil01/solution/baggage

go 1.22.1

require (
	github.com/rodsil01/solution/contracts v1.0.0
	github.com/rodsil01/solution/messagebroker v1.0.0
)

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/streadway/amqp v1.1.0 // indirect
)

replace (
	github.com/rodsil01/solution/contracts => ../contracts
	github.com/rodsil01/solution/messagebroker => ../message-broker
)
