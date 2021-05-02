package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/manoj-gupta/glance/internal/db"
	"github.com/manoj-gupta/glance/internal/models"
)

// GetTodos .. API handler to return all todo items
func GetTodos(c *gin.Context) {
	var todo []models.Todo
	if err := db.DB.Find(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// CreateTodo .. API handler to create a todo item
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	if err := db.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusNotAcceptable,
			gin.H{"error": "todo exists"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// GetTodo .. API handler to return a todo item with specified id
func GetTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateTodo .. API handler to update a todo item with specified id
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "todo not found"})
		return
	}

	// set status to `true` and write back to db
	todo.Status = true
	c.BindJSON(&todo)
	db.DB.Save(todo)

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo .. API handler to delete a todo item with specified id
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	if err := db.DB.Where("id = ?", id).Delete(todo); err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
}

// UndoTodo .. undo a task
func UndoTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "todo not found"})
		return
	}

	// set status to `false` and write back to db
	todo.Status = false
	c.BindJSON(&todo)
	db.DB.Save(todo)

	c.JSON(http.StatusOK, todo)
}
