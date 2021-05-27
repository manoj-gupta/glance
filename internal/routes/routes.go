package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/manoj-gupta/glance/internal/controllers"
)

// Init ... Initialize routes
func Init() (*gin.Engine, error) {
	r := gin.Default()

	config := cors.DefaultConfig()

	// origin
	config.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}

	// To be able to send tokens to the server.
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	// OPTIONS method for ReactJS
	config.AddAllowMethods("OPTIONS")

	// Register the middleware
	r.Use(cors.New(config))

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
		// login routes
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
		v1.GET("/user", controllers.User)
		v1.POST("/logout", controllers.Logout)

		// todo routes
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.CreateTodo)
		v1.GET("todo/:id", controllers.GetTodo)
		v1.PUT("todo/:id", controllers.UpdateTodo)
		v1.DELETE("todo/:id", controllers.DeleteTodo)
		v1.PUT("undoTodo/:id", controllers.UndoTodo)
	}
}
