package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"github.com/manoj-gupta/glance/internal/db"
	"github.com/manoj-gupta/glance/internal/models"
	"github.com/manoj-gupta/glance/internal/routes"
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
	err := app.Init()
	if err == nil {
		// Start app
		app.Run()
	}

	// Clean up app
	defer app.DeInit()
}

// App .. structure to keep app constructs
type App struct {
	router *gin.Engine
	db     *gorm.DB
}

// Init ... init app
func (a *App) Init() error {
	var err error

	// Initialize Database
	a.db, err = db.Init()
	if err != nil {
		fmt.Println("Database Initialization failed")
		return err
	}

	// DB migrate
	a.db.AutoMigrate(&models.Todo{})

	// Initialize router
	a.router, err = routes.Init()
	if err != nil {
		fmt.Println("Routes Initialization failed")
		return err
	}

	return nil
}

// Run ... run app
func (a *App) Run() {
	// Run router
	a.router.Run()
}

// DeInit ... deinit app
func (a *App) DeInit() {
	if a.db != nil {
		db.DeInit(a.db)
	}

	if a.router != nil {
		routes.DeInit(a.router)
	}
}
