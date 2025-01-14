package main

import (
	"net/http"

	"chapter_17/meander"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router instance
	router := gin.Default()

	// Define a route for /journeys
	router.GET("/journeys", func(c *gin.Context) {
		c.JSON(http.StatusOK, meander.Journeys) // Respond with JSON
	})

	// Start the server on port 8080
	router.Run(":8080")
}
