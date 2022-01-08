package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }

//EventMediator
func (em *EventMediator) AttachEventChannel(ec EventChannel) {
	em.EventChannels = ec
}

//EventProcessor
func (ep *EventProcessor) AddComponent(f func()) {
	ep.Components = append(ep.Components, f)
}

//EventQueue
// func (eq *EventQueue) AddQueue(q amqp.Queue) {
// 	eq.Queue = append(eq.Queue, q)
// }

func (eq *EventQueue) AddEventMediator(em EventMediator) {
	eq.EventMediators = append(eq.EventMediators, em)
}

// //EventChannel
// func (ec *EventChannel) AddQueue(ch amqp.Channel) {
// 	ec.Channel = append(ec.Channel, ch)
// }

func (ec *EventChannel) AddEventProcessor(ep EventProcessor) {
	ec.EventProcessors = append(ec.EventProcessors, ep)
}

func CreateEventMediator(associations map[Message]string, ec EventChannel) EventMediator {
	return EventMediator{
		RecognizedEvents: associations,
		EventChannels:    ec,
	}
}

func CreateEvent(eq EventQueue) Event {
	return Event{
		eq,
	}
}

func CreateEventChannel(conn map[amqp.Queue]amqp.Channel, ep []EventProcessor) EventChannel {
	return EventChannel{
		Conn:            conn,
		EventProcessors: ep,
	}
}

func CreateEventQueue(q amqp.Queue, em []EventMediator) EventQueue {
	return EventQueue{
		q,
		em,
	}
}
