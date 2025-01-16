package main

import (
	"fmt"
	"sync"
	"time"
)

type TimeStruct struct {
	totalChanges int          //tracking no of updates to current time
	currentTime  time.Time    //storing current time
	rwLock       sync.RWMutex //for locking
}

var TimeElement TimeStruct //a global instance of TimeStruct to be modified by goroutines

func updateTime() {
	TimeElement.rwLock.Lock() //prevent other goroutines from reading or writing
	defer TimeElement.rwLock.Unlock()
	TimeElement.currentTime = time.Now() //update current time
	TimeElement.totalChanges++           //increment total changes
}
func main() {
	var wg sync.WaitGroup
	TimeElement.totalChanges = 0
	TimeElement.currentTime = time.Now()
	timer := time.NewTicker(1 * time.Second)       //a ticker that triggers every 1 second to print total changes and current time
	writeTimer := time.NewTicker(10 * time.Second) //a ticker that triggers every 10 seconds to call updateTime and modify TimeElement
	endTimer := make(chan bool)                    //a channel used to signal the goroutine to stop
	wg.Add(1)
	go func() { //anonymous goroutine
		for { //infinite loop waiting for events from select block
			select {
			case <-timer.C: //executes every second
				fmt.Println(TimeElement.totalChanges,
					TimeElement.currentTime.String())
			case <-writeTimer.C: //executes every 10 seconds
				updateTime()
			case <-endTimer: //stops the timer and exits the loop
				timer.Stop()
				return
			}
		}
	}()
	wg.Wait()
	fmt.Println(TimeElement.currentTime.String())
}
