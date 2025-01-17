package main

import (
	"fmt"
	"sync"
	"time"
)

var ready = false                      //shared condition variable
var cond = sync.NewCond(&sync.Mutex{}) //condition variable with mutex

func main() {
	// Consumer Goroutine
	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()

		for !ready {
			fmt.Println("Waiting for signal...")
			cond.Wait() //Wait for signal
		}
		fmt.Println("Got the signal - Ready is true!")
	}()
	// Producer Goroutine
	go func() {
		time.Sleep(2 * time.Second)
		cond.L.Lock()
		ready = true
		fmt.Println("Ready is true - Sending signal")
		cond.L.Unlock()
		cond.Signal()
	}()

	time.Sleep(3 * time.Second) // lazy wait for goroutines to finish
}
