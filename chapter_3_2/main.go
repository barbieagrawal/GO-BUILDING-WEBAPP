package main

import (
	"fmt"
	"sync"
)

// HW Question - increment counter functionality with calling 10 goroutines using WaitGroup

// using WaitGroups
// func main() {
// 	counter := 0
// 	iterations := 10
// 	wg := new(sync.WaitGroup)
// 	wg.Add(iterations)

// 	for i := 0; i < iterations; i++ {
// 		go func() {
// 			counter++
// 			fmt.Println(counter)
// 			defer wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(counter)
// }

// using mutex lock -> avoids race conditions
func main() {
	counter := 0
	iterations := 10
	var mu sync.Mutex
	wg := new(sync.WaitGroup)
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
		fmt.Println(counter)
	}
	wg.Wait()
	fmt.Println("Final value:", counter)
}
