package main

import "github.com/streadway/amqp"

// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }

//EventMediator
func (em *EventMediator) AttachEventChannel(ec EventChannel) {
	em.EventChannels = append(em.EventChannels, ec)
}

func (em *EventMediator) AddRecognizedEvent(s string) {
	em.RecognizedEvents = append(em.RecognizedEvents, s)
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

//EventChannel
func (ec *EventChannel) AddQueue(q amqp.Queue) {
	ec.Queue = append(ec.Queue, q)
}

func (ec *EventChannel) AddEventProcessor(ep EventProcessor) {
	ec.EventProcessors = append(ec.EventProcessors, ep)
}
