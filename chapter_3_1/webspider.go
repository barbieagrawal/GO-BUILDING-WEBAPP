package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var applicationStatus bool
var urls []string
var urlsProcessed int
var foundUrls []string
var fullText string
var totalURLCount int
var wg sync.WaitGroup
var v1 int

// Fetches URLs and sends the content to textChannel. It also signals completion on statusChannel.
func readURLs(statusChannel chan int, textChannel chan string) {
	fmt.Println("Grabbing", len(urls), "urls")
	for i := 0; i < totalURLCount; i++ {
		resp, err := http.Get(urls[i])
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			continue
		}
		text, _ := io.ReadAll(resp.Body)
		textChannel <- string(text)
		resp.Body.Close()
		statusChannel <- 0
	}
}

// Reads content from textChannel and appends it to fullText. Monitors processChannel for a termination signal.
func addToScrapedText(textChannel chan string, processChannel chan bool) {
	for {
		select {
		case text := <-textChannel:
			fullText += text
		case stop := <-processChannel:
			if stop {
				close(textChannel)
				close(processChannel)
				return
			}
		}
	}
}

// Tracks the number of processed URLs using statusChannel. Sends a signal to terminate the process when all URLs are done.
func evaluateStatus(statusChannel chan int, processChannel chan bool) {
	for {
		select {
		case <-statusChannel:
			urlsProcessed++
			if urlsProcessed == totalURLCount {
				fmt.Println("All URLs processed")
				processChannel <- true
				return
			}
		}
	}
}

// func main() {
// 	//initialize channels and global variables
// 	applicationStatus = true
// 	statusChannel := make(chan int)
// 	textChannel := make(chan string)
// 	processChannel := make(chan bool)

// 	urls = []string{
// 		"http://www.mastergoco.com/index1.html",
// 		"http://www.mastergoco.com/index2.html",
// 		"http://www.mastergoco.com/index3.html",
// 		"http://www.mastergoco.com/index4.html",
// 		"http://www.mastergoco.com/index5.html",
// 	}

// 	totalURLCount = len(urls)

// 	go readURLs(statusChannel, textChannel)
// 	go addToScrapedText(textChannel, processChannel)
// 	go evaluateStatus(statusChannel, processChannel)

// 	for {
// 		if !applicationStatus {
// 			fmt.Println("Scraped Content:\n", fullText)
// 			fmt.Println("Done!")
// 			break
// 		}
// 	}
// }
