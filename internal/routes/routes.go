package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manoj-gupta/glance/internal/models"
	"github.com/manoj-gupta/glance/internal/views"
)

func handleGetTasks(c *gin.Context) {
	events := models.GetDefaultEvents()
	c.JSON(http.StatusOK, gin.H{"events": events})
}

// Init ... Initialize routes
func Init() (*gin.Engine, error) {
	r := gin.Default()

	// load html templates
	r.LoadHTMLGlob("templates/*")

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

	// Handle the index route
	r.GET("/", showIndexPage)

	// Handle the todo route
	setupToDoRoutes(r)
	//r.GET("/todos", showTodosPage)

	// Handle GET requests at /event/view/event_id
	r.GET("/event/view/:event_id", showEvent)
}

func setupToDoRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("todo", views.GetTodos)
		v1.POST("todo", views.CreateTodo)
		v1.GET("todo/:id", views.GetTodo)
		v1.PUT("todo/:id", views.UpdateTodo)
		v1.DELETE("todo/:id", views.DeleteTodo)
	}
}

func showIndexPage(c *gin.Context) {
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title": "Home Page",
		},
	)
}

func showTodosPage(c *gin.Context) {
	events := models.GetDefaultEvents()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"events.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": events,
		},
	)
}

func showEvent(c *gin.Context) {
	eID, err := strconv.Atoi(c.Param("event_id"))

	if err != nil {
		// invalid event ID specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	event, err := models.GetEventsByID(eID)

	if err != nil {
		// event not found, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Call  HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the event.html template
		"event.html",
		// Pass the data that the page uses
		gin.H{
			"title":   event.Title,
			"payload": event,
		},
	)

	return
}
