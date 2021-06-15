package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()
	channel, err := conn.Channel()

	if err != nil {
		panic(err.Error())
	}
	// Create queue to send message to it
	queue, err := channel.QueueDeclare(
		"test_queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
		)
	if err != nil {
		panic(err.Error())
	}

	payload := "Sending Message"
	err = channel.Publish(
		"",     // exchange
		queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})

	if err != nil {
		panic(err.Error())
	}
	log.Printf(" [x] Congrats, sending message: %s", payload)
}

