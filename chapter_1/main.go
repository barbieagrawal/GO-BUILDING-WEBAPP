package main

import (
	"fmt"
	//"github.com/gorilla/mux"
	"net/http" 
	"time"
)

const (
	Port = ":8084"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The time is now " + time.Now().String()
	fmt.Fprintln(w,response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func main() {
	http.HandleFunc("/static",serveStatic)
	http.HandleFunc("/",serveDynamic)
	http.ListenAndServe(Port,nil)
}
