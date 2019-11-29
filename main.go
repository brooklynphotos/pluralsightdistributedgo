package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// server()
	client()
}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(q.Name, "goconsumer", true, false, false, false, nil)

	failOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		log.Printf("Received message: %s", msg.Body)
	}
}

/**
 * serves messages
 */
func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello RabbitMQ"),
	}

	// publish to a blank exchange with this message we just created
	ch.Publish("", q.Name, true, false, msg)
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	// make the connection
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	// create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// create a queue and it must not already exist
	// we are using Direct exchange so it will look like things went directly to the queue
	q, err := ch.QueueDeclare("hello",
		false,
		false,
		false, // for now ok to be exclusive
		false,
		nil)
	failOnError(err, "Failed to open a queue")
	return conn, ch, &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
