package main

import "fmt"

// thread 2
func main() {
	channel := make(chan string) //empty channel

	//thread 2
	go func() {
		channel <- "Hellow wolrd!" // full channel
	}()

	msg := <- channel //empty channel
	fmt.Println(msg)
}
