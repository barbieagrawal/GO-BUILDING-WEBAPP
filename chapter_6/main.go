package main

import (
	"github.com/barbieagrawal/chapter6/controllers"
	"github.com/barbieagrawal/chapter6/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.Run()
}
