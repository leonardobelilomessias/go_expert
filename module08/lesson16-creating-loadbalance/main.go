package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d received %d \n ", workerId, x)
		time.Sleep(time.Second)
	}
}
func main() {
	data := make(chan int)
	amounworkers := 1000000
	//start workers
	for i := 0; i < amounworkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 10000000; i++ {
		data <- i
	}
}
