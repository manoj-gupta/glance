package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/manoj-gupta/glance/internal/db"
	"github.com/manoj-gupta/glance/internal/routes"
)

// App .. structure to keep app constructs
type App struct {
	router *gin.Engine
	db     *sql.DB
}

// Initialize ... init app
func (a *App) Initialize() error {
	var err error

	a.db, err = db.Initialize()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Create the router
	a.router = gin.Default()

	// Initialize the routes
	routes.Init(a.router)
	routes.InitializeRoutes(a.router)

	return nil
}

// Run ... run app
func (a *App) Run() {
	// Run router
	a.router.Run()
}

// DeInitialize ... deinit app
func (a *App) DeInitialize() {
	a.db.Close()
}
