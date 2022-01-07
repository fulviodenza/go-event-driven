package main

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/streadway/amqp"
)

type Message struct {
	n interface{}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
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

func (e Event) Enqueue(m Message) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	b, err := getBytes(m.n)
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

func getBytes(i interface{}) ([]byte, error) {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(i)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
