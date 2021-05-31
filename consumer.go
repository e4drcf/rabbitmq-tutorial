package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

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

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
		}
	}()
	fmt.Println("Successfully Connected To our RabbitMQ Instance")
	fmt.Println(" [*] - waiting for messages ")
	<-forever
}
