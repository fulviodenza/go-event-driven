package main

import (
	"bytes"
	"encoding/gob"

	"github.com/streadway/amqp"
)

type Message struct {
	n interface{}
}

//This operation should be done from the user

// q, err := ch.QueueDeclare(
// 	"hello", // name
// 	false,   // durable
// 	false,   // delete when unused
// 	false,   // exclusive
// 	false,   // no-wait
// 	nil,     // arguments
//   )
//   failOnError(err, "Failed to declare a queue")

// Enqueue is performed by the Event which triggers the action,
// After the trigger the event should be enqueued to the single EventQueue
// And then should be Dequeued by the EventMediator which will dispatch
// the event inside the correct EventChannel

func (e Event) Enqueue(m Message, conn *amqp.Connection) error {

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	b, err := GetBytes(m.n)
	failOnError(err, "Failed to cast the message into bytes")

	err = ch.Publish(
		"",                      // exchange
		e.EventQueue.Queue.Name, // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(b),
		})
	failOnError(err, "Failed to publish a message")

	return nil
}

func GetBytes(i interface{}) ([]byte, error) {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(i)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
