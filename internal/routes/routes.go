package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/manoj-gupta/glance/internal/views"
)

// Init ... Initialize routes
func Init() (*gin.Engine, error) {
	r := gin.Default()

	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours

	// Using DefaultConfig as start point
	// config := cors.DefaultConfig()
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	// r.Use(cors.New(config))

	// Default() allows all origins
	r.Use(cors.Default())

	// Initialize the routes
	InitializeRoutes(r)

	return r, nil
}

// DeInit ... Initialize routes
func DeInit(r *gin.Engine) {
	//TODO
}

// InitializeRoutes ...
func InitializeRoutes(r *gin.Engine) {
	// Handle the todo route
	setupToDoRoutes(r)
}

func setupToDoRoutes(r *gin.Engine) {
	v1 := r.Group("/api")
	{
		v1.GET("todo", views.GetTodos)
		v1.POST("todo", views.CreateTodo)
		v1.GET("todo/:id", views.GetTodo)
		v1.PUT("todo/:id", views.UpdateTodo)
		v1.DELETE("todo/:id", views.DeleteTodo)
		v1.PUT("undoTodo/:id", views.UndoTodo)
	}
}
