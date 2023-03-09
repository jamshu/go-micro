package main

import (
	"fmt"
	"listner/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	rabbitConn, err := connect()
	if err != nil {
		fmt.Println("connection to rabbitmq failed")
		os.Exit(1)
	}
	defer rabbitConn.Close()

	//start listening
	log.Println("Listening for RabbitMQ messages...")

	//create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	//watch the queue and consume messages

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		panic(err)
	}

}

func connect() (*amqp.Connection, error) {
	fmt.Println("Connecting")
	var count int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			count++
			fmt.Println("RabbitMQ connection not ready: ")
		} else {
			fmt.Println("Connection to rabbitmq established")
			connection = c
			break
		}
		if count > 5 {
			fmt.Println(err)
			return nil, err
		}
		backoff = time.Duration(math.Pow(float64(count), 2)) * time.Second
		fmt.Println("Backoff ................")
		time.Sleep(backoff)
		continue
	}
	return connection, nil
}
