package main

import (
	"PMD/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	const functionName = "main.main"
	fmt.Println(functionName)
	router = gin.Default()
	handlers.SetupRoutes(router)
	router.Run(":4000")
}
