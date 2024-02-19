package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	go func() {
		for {
			atomic.AddInt64(&i, 1)

			msg := Message{i, "hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)

			msg := Message{i, "hello from Apacke Kafka"}
			c2 <- msg
		}
	}()
	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from RabbitMQ :ID %d -  %s\n", msg.Id, msg.Msg)

		case msg := <-c2:
			fmt.Printf("Received from Apacke Kafka :ID %d -  %s\n", msg.Id, msg.Msg)

		case <-time.After(time.Second * 3):
			println("timeout")
			// default:
			// 	println("default")
		}
	}

}
