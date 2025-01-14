package main

import (
	"chapter_15/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.SetupRoutes(router)

	router.Run(":8080")

}
