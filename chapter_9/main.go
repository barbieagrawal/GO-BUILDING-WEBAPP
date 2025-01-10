package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

//Reflected XSS implementation

// Handlers for HTML input
func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	fmt.Fprintln(w, input)
}

func HTMLHandlerSafe(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	escapedInput := template.HTMLEscapeString(input)
	fmt.Fprintln(w, escapedInput)
}

// Handlers for JavaScript input
func JSHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	fmt.Fprintln(w, input)
}

func JSHandlerSafe(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	escapedInput := template.JSEscapeString(input)
	fmt.Fprintln(w, escapedInput)
}

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/html", HTMLHandler)
	router.HandleFunc("/html_safe", HTMLHandlerSafe)
	router.HandleFunc("/js", JSHandler)
	router.HandleFunc("/js_safe", JSHandlerSafe)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
