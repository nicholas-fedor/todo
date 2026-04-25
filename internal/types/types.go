package types

import "time"

// TodoItem represents a todo item.
type TodoItem struct {
	Key       string    `json:"key"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
