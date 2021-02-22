package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	oneshot = 0
	monthly = 24 * 60 * 60
)

// Event - Model of a basic event
type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Freq  int32  `json:"freq"`
	Desc  string `json:"desc"`
}

func handleGetTasks(c *gin.Context) {
	var events = []Event{
		{
			ID:    1,
			Title: "Pay Electricity Bill",
			Freq:  oneshot,
			Desc:  "One time bill",
		},
		{
			ID:    2,
			Title: "Pay Gas Bill",
			Freq:  monthly,
			Desc:  "One time bill",
		},
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

func main() {
	fmt.Println("Hello, Glance")

	// Create the router
	router := gin.Default()

	// load html templates
	router.LoadHTMLGlob("templates/*")

	// router handlers
	router.GET("/", func(c *gin.Context) {
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

	router.GET("/events/", handleGetTasks)

	// Run router
	router.Run()
}
