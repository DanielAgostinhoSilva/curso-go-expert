package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID      uint64
	Message string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var idCount uint64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddUint64(&idCount, 1)
			time.Sleep(time.Second)
			msg := Message{ID: idCount, Message: "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddUint64(&idCount, 1)
			time.Sleep(time.Second * 2)
			msg := Message{ID: idCount, Message: "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("received from RabbitMQ: ID: %d - %s\n", msg1.ID, msg1.Message)
		case msg2 := <-c2:
			fmt.Printf("received from Kafka: ID: %d - %s\n", msg2.ID, msg2.Message)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}
}
