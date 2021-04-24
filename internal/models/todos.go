package models

import (

	// needed for "postgres" driver
	_ "github.com/lib/pq"

	"github.com/manoj-gupta/glance/internal/db"
)

// GetTodos .. DB interface to return all todo items
func GetTodos(todo *[]Todo) (err error) {
	if err = db.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// CreateTodo .. DB interface to create a todo item
func CreateTodo(todo *Todo) (err error) {
	if err = db.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// GetTodo .. DB interface to return a todo item with specified id
func GetTodo(todo *Todo, id string) (err error) {
	if err := db.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTodo .. DB interface to update a todo item with specified id
func UpdateTodo(todo *Todo, id string) (err error) {
	db.DB.Save(todo)
	return nil
}

// DeleteTodo .. DB interface to delete a todo item with specified id
func DeleteTodo(todo *Todo, id string) (err error) {
	db.DB.Where("id = ?", id).Delete(todo)
	return nil
}
