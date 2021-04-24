package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/manoj-gupta/glance/internal/models"
)

// GetTodos .. API handler to return all todo items
func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := models.GetTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// CreateTodo .. API handler to create a todo item
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// GetTodo .. API handler to return a todo item with specified id
func GetTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := models.GetTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// UpdateTodo .. API handler to update a todo item with specified id
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.GetTodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
		return
	}
	todo.Status = true
	c.BindJSON(&todo)
	err = models.UpdateTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// DeleteTodo .. API handler to delete a todo item with specified id
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.DeleteTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}

// UndoTodo .. undo a task
func UndoTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.GetTodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
		return
	}
	todo.Status = false
	c.BindJSON(&todo)
	err = models.UpdateTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
