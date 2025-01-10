package main

import (
	"log"
	"os"
)

var (
	Warn  *log.Logger
	Error *log.Logger
)

func main() {
	warnFile, err := os.OpenFile("warnings.log", os.O_RDWR|os.O_APPEND, 0660)
	defer warnFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	errorFile, err := os.OpenFile("error.log", os.O_RDWR|os.O_APPEND, 0660)
	defer errorFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	Warn = log.New(warnFile, "WARNING barb: ", log.LstdFlags)
	Error = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime)

	numone := 5
	numtwo := 10
	if numone < numtwo {
		Warn.Println("numone is less than numtwo")
	} else {
		Error.Println("numone is not less than numtwo")
	}

	Warn.Println("This is a warning message")
	Error.Println("This is an error message")
}
