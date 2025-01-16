package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var (
	writer chan bool
	rwLock sync.RWMutex
	wg     sync.WaitGroup
)

func writeFile(num int) {
	defer wg.Done()

	rwLock.Lock()
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		rwLock.Unlock()
		return
	}

	_, err = f.WriteString(strconv.FormatInt(int64(num), 10) + "\n")
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}

	f.Close()
	rwLock.Unlock()

	writer <- true
}

func main() {
	writer = make(chan bool, 10)

	// Create empty file
	os.WriteFile("test.txt", []byte(""), 0777)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeFile(i)
		<-writer // Wait for write to complete
	}

	wg.Wait()
	fmt.Println("Done!")
}
