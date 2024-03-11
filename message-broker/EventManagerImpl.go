package messagebroker

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

type EventManagerImpl struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	handlers []EventHandler
}

var instance *EventManagerImpl
var once sync.Once

func GetEventManager() *EventManagerImpl {
	once.Do(func() {
		instance = &EventManagerImpl{}
	})
	return instance
}

func (em *EventManagerImpl) Initialize(url string) error {
	conn, err := amqp.Dial(url)

	if err != nil {
		return err
	}

	em.conn = conn

	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	em.channel = ch

	fmt.Println("Connection to " + url + " stablished successfuly")

	return nil
}

func (em *EventManagerImpl) InjectHandlers(handlers ...EventHandler) {
	em.handlers = append(em.handlers, handlers...)
}

func (em *EventManagerImpl) PublishEvent(eventType string, payload interface{}) error {
	event := Event{
		Type: eventType,
		Data: payload,
	}

	body, err := json.Marshal(event)

	if err != nil {
		return err
	}

	err = em.channel.Publish(
		"",
		"events_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (em *EventManagerImpl) ConsumeEvents() error {
	q, err := em.channel.QueueDeclare(
		"events_queue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	msgs, err := em.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			var event Event

			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Printf("Failed to unmarshal event: %v", err)
				continue
			}

			for _, handler := range em.handlers {
				go handler.HandleEvent(event.Type, event.Data)
			}
		}
	}()

	return nil
}

func (em *EventManagerImpl) Close() {
	if em.channel != nil {
		em.channel.Close()
	}
	if em.conn != nil {
		em.conn.Close()
	}
}
