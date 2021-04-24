package models

const (
	oneshot = 0
	monthly = 24 * 60 * 60
)

// Todo .. json binding to models.Todo
type Todo struct {
	ID     uint   `json:"_id"`
	Task   string `json:"task"`
	Status bool   `json:"status,omitempty"`
}

// TableName .. todo table name
func (b *Todo) TableName() string {
	return "todo"
}
