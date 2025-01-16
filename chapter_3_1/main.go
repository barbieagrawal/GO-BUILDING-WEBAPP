package main

import (
	"strings"
	"sync"
)

var initialString string //holds initial string
var finalString string   //will hold final capitalized string
var stringLength int     //length of initial string

func capitalize(letterChannel chan string, currentLetter string, wg *sync.WaitGroup) {
	thisLetter := strings.ToUpper(currentLetter) //capitalize
	wg.Done()                                    //tell waitgroup this goroutine is done
	letterChannel <- thisLetter                  //send capitalized letter to the channel
}

func addToFinalStack(letterChannel chan string, wg *sync.WaitGroup) {
	letter := <-letterChannel //Receive letter from channel
	finalString += letter     //Add it to the final string
	wg.Done()                 //signal that goroutine is done
}

// func main() {
// 	runtime.GOMAXPROCS(2)
// 	var wg sync.WaitGroup //waitgroup is created to sync goroutines
// 	initialString = "Four score and seven years ago our fathers brought forth on this continent, a new nation, conceived in Liberty, and dedicated to the proposition that all men are created equal."
// 	initialBytes := []byte(initialString)             //convert string to byte slice to process it letter by letter
// 	var letterChannel chan string = make(chan string) //channel is created
// 	stringLength = len(initialBytes)
// 	for i := 0; i < stringLength; i++ {
// 		wg.Add(2)                                                  //add 2 tasks to waitgroup
// 		go capitalize(letterChannel, string(initialBytes[i]), &wg) //capitalize
// 		go addToFinalStack(letterChannel, &wg)                     //add letter to final string
// 		wg.Wait()                                                  //wait for both tasks to finish before continuing
// 	}
// 	fmt.Println(finalString)
// }
