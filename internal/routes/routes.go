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

// InitializeRoutes ...
func InitializeRoutes(r *gin.Engine) {
	// load html templates
	r.LoadHTMLGlob("templates/*")

	// router handlers
	r.GET("/", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"Welcome": "to Glance"})
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})

	r.GET("/events/", handleGetTasks)
}
