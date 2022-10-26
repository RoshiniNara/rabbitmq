package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {

	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	// create a channel to block the execution of main function until a new message is received
	forever := make(chan bool)
	//use go routines to send values to channel
	go func() {
		for d := range msgs {
			fmt.Printf("Received messages: %s\n", d.Body)
		}
	}()
	fmt.Println("Successuly connected to our RabbitMQ Instance")
	fmt.Println(" [*] - waiting for messages")
	//receives the value from the channel.
	<-forever

}
