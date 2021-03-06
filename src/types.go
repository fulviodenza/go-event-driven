package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Initiating event is the event that starts the whole event process, it is sent to
// an initiating event queue which is accepted by the event mediator.
//created

type Event struct {
	EventQueue EventQueue
}

// An Event Mediator manages and controls the workflow for initiating events that
// require the coordination of multiple event processors. The event mediator only knows
// the steps involved in processing the event and therefore generates corresponding
// point-to-point messaging fashion
//created

type EventMediator struct {
	RecognizedEvents map[Message]string
	// Topics           []string
	EventChannels EventChannel
}

// Event Processors listen to dedicated event channels, process the event, and usually
// respond back to the mediator that they have completed their work.
type EventProcessor struct {
	Components []func()
}

// Event Queue is a queue from which the events starts, Initiating Event component
// send the event to the Event Queue which sends that event to Event(s) Mediator
//created
type EventQueue struct {
	Queue          amqp.Queue
	EventMediators []EventMediator
}

type EventChannel struct {
	Conn            map[amqp.Queue]amqp.Channel
	EventProcessors []EventProcessor
}
