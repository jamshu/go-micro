package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var webPort = "80"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	rabbitConn, err := connect()
	if err != nil {
		fmt.Println("connection to rabbitmq failed")
		os.Exit(1)
	}
	defer rabbitConn.Close()
	app := Config{
		Rabbit: rabbitConn,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	fmt.Printf("running the broker service on the port %s\n", webPort)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
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
