package models

import (
	// needed for "postgres" driver
	_ "github.com/lib/pq"
)

// Todo .. todo list structure
type Todo struct {
	ID     uint   `json:"_id"`
	Task   string `json:"task"`
	Status bool   `json:"status,omitempty"`
}

// TableName .. todo table name
func (b *Todo) TableName() string {
	return "todo"
}
