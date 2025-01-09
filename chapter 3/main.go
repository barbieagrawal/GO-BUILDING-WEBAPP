package main

import (
    "database/sql"
    "fmt"
    _"github.com/lib/pq"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

const (
	dsn = "user=postgres password=noob101 dbname=cms sslmode=disable"
)

// Connect to the PostgreSQL database
func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Page handler that serves the page data from the database
func pageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Connect to the database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer db.Close()

	// Variables to store data from the query
	var pageTitle, pageContent, pageDate string

	// Query the database for page data using the provided ID
	err = db.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE id = $1", id).Scan(&pageTitle, &pageContent, &pageDate)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		log.Println(err)
		return
	}

	// Send the response back to the user
	fmt.Fprintf(w, "Page Title: %s\nContent: %s\nDate: %s", pageTitle, pageContent, pageDate)
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define the route for pages with dynamic ID
	r.HandleFunc("/page/{id:[0-9]+}", pageHandler)

	// Start the server
	fmt.Println("Server started at :8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}