// Package app provides the application layer business logic.
package app

import (
	"encoding/json"
	"fmt"

	"github.com/nicholas-fedor/todo/internal/storage"
	"github.com/nicholas-fedor/todo/internal/types"
)

// ListTodos fetches and decodes all todos from the store.
func ListTodos(store storage.Store) ([]types.TodoItem, error) {
	keys, err := store.List()
	if err != nil {
		return nil, fmt.Errorf("list todos: %w", err)
	}

	todos := make([]types.TodoItem, 0, len(keys))
	for _, key := range keys {
		val, err := store.Get(key)
		if err != nil {
			continue
		}

		var todo types.TodoItem

		err = json.Unmarshal(val, &todo)
		if err != nil {
			continue
		}

		todo.Key = key
		todos = append(todos, todo)
	}

	return todos, nil
}
