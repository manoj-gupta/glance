package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/manoj-gupta/glance/internal/routes"
)

func main() {
	fmt.Println("Hello, Glance")

	// Create the router
	router := gin.Default()

	// Initialize the routes
	routes.InitializeRoutes(router)

	// Run router
	router.Run()
}
