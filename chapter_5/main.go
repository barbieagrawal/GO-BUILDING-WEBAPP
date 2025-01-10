package main

import (
	"errors"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type todoList struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var initialtodoList = []todoList{
	{ID: 1, Item: "Clean Room", Completed: false},
	{ID: 2, Item: "Read Book", Completed: false},
	{ID: 3, Item: "Record Video", Completed: false},
}

func initDatabase() {
	dsn := "host=localhost user=postgres password=noob101 dbname=todo_db port=5432 sslmode=disable"
	var err error

	// Connect to the database
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Automatically migrate the schema
	err = db.AutoMigrate(&todoList{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Insert initial todoList into the database if not already present
	for _, t := range initialtodoList {
		var existingTodo todoList
		// Check if the todo already exists to avoid duplicates
		if err := db.First(&existingTodo, "id = ?", t.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Insert new todo if it doesn't exist
				if err := db.Create(&t).Error; err != nil {
					log.Printf("Failed to insert todo %v: %v", t, err)
				}
			} else {
				log.Printf("Error checking existing todo %v: %v", t, err)
			}
		}
	}

	log.Println("Database connected and initial data inserted successfully!")
}

func gettodoList(context *gin.Context) { // Get request
	var todoListList []todoList
	if err := db.Find(&todoListList).Error; err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch todoList"})
		return
	}
	context.IndentedJSON(http.StatusOK, todoListList)
}

func addTodo(context *gin.Context) { // Post request
	var newTodo todoList
	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	if err := db.Create(&newTodo).Error; err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create todo"})
		return
	}
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todoList, error) { 
	var todo todoList
	if err := db.First(&todo, "id = ?", id).Error; err != nil {
		return nil, errors.New("todo not found")
	}
	return &todo, nil
}

func getTodo(context *gin.Context) { // Get request
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggletodoListtatus(context *gin.Context) { //Patch request
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	if err := db.Save(todo).Error; err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update todo"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	initDatabase()
	router := gin.Default()
	router.GET("/todoList", gettodoList)           // Get todoList list
	router.POST("/todoList", addTodo)           // Add a todo
	router.GET("/todoList/:id", getTodo)        // Get a todo by ID
	router.PATCH("/todoList/:id", toggletodoListtatus) // Update todo status
	router.Run("localhost:8084")
}
