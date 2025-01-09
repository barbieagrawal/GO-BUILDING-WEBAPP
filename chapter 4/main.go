package main

import (
    "database/sql"
    "fmt"
	"html/template"
    _"github.com/lib/pq"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

const (
	dsn = "user=postgres password=noob101 dbname=cms sslmode=disable"
)

// Page struct to hold page data
type Page struct {
	Title    	string
	RawContent 	string
	Content    	template.HTML
	Date    	string
	GUID 		string
}

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

	// Create a Page instance to store data
	thisPage := Page{}

	// Query the database for page data using the provided ID
	err = db.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE id = $1", id).Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date) 
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		log.Println(err)
		return
	}

	// Convert raw content to template.HTML to render safely
	thisPage.Content = template.HTML(thisPage.RawContent)

	// Parse the template
	t, err := template.ParseFiles("templates/blog.html")
	if err != nil {
		http.Error(w, "Template parsing failed", http.StatusInternalServerError)
		log.Println("Template error:", err)
		return
	}

	// Render the template with the page data
	err = t.Execute(w, thisPage)
	if err != nil {
		http.Error(w, "Template execution failed", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
		return
	}

	log.Println("Page served successfully:", thisPage)
}

func RedirIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

//Index Page Handler
func pageIndex(w http.ResponseWriter, r *http.Request) {
	//Connect to the database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer db.Close()

	var Pages = []Page{}

	rows, err := db.Query("SELECT page_title, page_content, page_date FROM pages ORDER BY page_date DESC")
	if err != nil {
		http.Error(w, "Failed to fetch pages", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		thisPage := Page{}
		err = rows.Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date)
		if err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		thisPage.Content = template.HTML(thisPage.RawContent)
		Pages = append(Pages, thisPage)
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	t.Execute(w, Pages)
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define the route for pages with dynamic ID
	r.HandleFunc("/page/{id:[0-9]+}", pageHandler)

	r.HandleFunc("/", RedirIndex) 
	r.HandleFunc("/home", pageIndex)

	// Start the server
	fmt.Println("Server started at :8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}