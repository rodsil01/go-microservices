package messagebroker

type EventManager interface {
	PublishEvent(eventType string, event interface{}) error
	ConsumeEvents() error
}
