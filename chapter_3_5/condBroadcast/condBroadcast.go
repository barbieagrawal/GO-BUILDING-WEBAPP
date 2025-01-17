package main

import (
	"fmt"
	"sync"
	"time"
)

var ready = false
var cond = sync.NewCond(&sync.Mutex{})

func main() {
	var wg sync.WaitGroup

	//Multiple consumers
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cond.L.Lock()
			for !ready {
				fmt.Printf("Consumer %d waiting for signal\n", id)
				cond.Wait()
			}
			fmt.Printf("Consumer %d received the signal\n", id)
			cond.L.Unlock()
		}(i)
	}

	//Producer
	go func() {
		time.Sleep(2 * time.Second)
		cond.L.Lock()
		ready = true
		fmt.Println("Producer is set and ready to go")
		cond.L.Unlock()
		cond.Broadcast()
	}()

	wg.Wait()
	fmt.Println("All consumers are finished")
}
