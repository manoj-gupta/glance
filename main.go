package main

import (
	"log"

	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Initialize app
	app := App{}
	err := app.Initialize()
	if err != nil {
		log.Fatal("App Initialization failed")
		return
	}

	// Clean up app
	defer app.DeInitialize()

	// Start app
	app.Run()
}
