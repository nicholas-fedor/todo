package models

import "time"

// TodoView represents a todo item for Templ rendering.
type TodoView struct {
	Key       string    `json:"key"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
