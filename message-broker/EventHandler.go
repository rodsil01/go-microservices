package messagebroker

type EventHandler interface {
	HandleEvent(eventType string, event interface{})
}
