package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"os"
)

const (
	PORT = ":8084"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //look for query string variables from http.Request and parse them into a map
	pageID := vars["id"] //values will be accessible by ID
	fileName := "files/" + pageID + ".html"

	_, err := os.Stat(fileName)
    if err != nil {
        fileName = "files/404.html"
    }

	http.ServeFile(w, r, fileName)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler) //regex for page ID
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}

