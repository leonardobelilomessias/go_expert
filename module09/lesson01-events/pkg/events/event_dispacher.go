package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already regitered")

type EventDispacher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispacher() *EventDispacher {
	return &EventDispacher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ev *EventDispacher) Dispatch(event EventInterface) error {
	if _, ok := ev.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range ev.handlers[event.GetName()] {
			wg.Add(1)
			go handler.Handler(event,wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispacher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
func (ed *EventDispacher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispacher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (ed *EventDispacher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}
