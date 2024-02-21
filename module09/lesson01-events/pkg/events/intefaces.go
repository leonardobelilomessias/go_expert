package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handler(event EventInterface)
}

type EventDispacherInterface interface {
	Register(event string, handle EventHandlerInterface) error
	Dispach(event EventInterface) error
	Remove(eventName string, handle EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}

