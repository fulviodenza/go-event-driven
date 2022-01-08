package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// This is the mediator's module, mediator dequeues
// the message and handles it based on user defined message
// handling. For example 1 = topic1, 2 = topic2
// here goes the events definition and their topic assignment and
// publish them, then the correct queue listen for the topic it is interested.
// If no queue is interested in that topic, the message is discarded
// as input takes

func (em *EventMediator) ConsumeMessage() {
	// consume message
}

func (em *EventMediator) SendMessage(m Message) {

	// check if the message's topic is inside the map and publish to the queues
	// listening for that topic
	if t, ok := em.RecognizedEvents[m]; ok {

		for _, ch := range em.EventChannels.Conn {

			err := ch.ExchangeDeclare(
				t,        //name
				"direct", // type
				true,     // durable
				false,    // auto-deleted
				false,    // internal
				false,    // no-wait
				nil,      // arguments
			)
			failOnError(err, "Failed to create exchange")
			// err = ch.QueueBind(
			// 	q.Name,
			// 	"message",
			// 	t,
			// 	false,
			// 	nil,
			// )
			// failOnError(err, "Failed to bind a queue")

			body, _ := GetBytes(m)

			err = ch.Publish(
				t,
				"",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				},
			)
			failOnError(err, "Failed to publish a message")
		}
	}
}
