package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manoj-gupta/glance/internal/model"
)

func handleGetTasks(c *gin.Context) {
	events := model.GetDefaultEvents()
	c.JSON(http.StatusOK, gin.H{"events": events})
}

// Init ...
func Init(r *gin.Engine) {
	// load html templates
	r.LoadHTMLGlob("templates/*")
}

// InitializeRoutes ...
func InitializeRoutes(r *gin.Engine) {

	// Handle the index route
	r.GET("/", showIndexPage)
}

func showIndexPage(c *gin.Context) {
	events := model.GetDefaultEvents()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": events,
		},
	)

}
