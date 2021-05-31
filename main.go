package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")

	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	fmt.Println("Successfully Connected To our RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer func() {
		err = ch.Close()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello Consumer"),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
