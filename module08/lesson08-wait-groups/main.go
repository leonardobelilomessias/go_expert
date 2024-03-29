package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string , wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d:Task %s is runnning\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup:= sync.WaitGroup{}
	waitGroup.Add(25)
	///	Thread 2
	go task("tarefa A",&waitGroup)
	///	Thread 3
	go task("tarefa B",&waitGroup)
	///	Thread 4
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d:Task %s is runnning\n", i, "anoniymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
