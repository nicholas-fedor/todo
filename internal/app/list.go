package app

import (
	"encoding/json"

	"github.com/nicholas-fedor/todo/internal/storage"
	"github.com/nicholas-fedor/todo/internal/types"
)

// ListTodos fetches and decodes all todos from the store.
func ListTodos(store storage.Store) ([]types.TodoItem, error) {
	keys, err := store.List()
	if err != nil {
		return nil, err
	}

	todos := make([]types.TodoItem, 0, len(keys))
	for _, key := range keys {
		val, err := store.Get(key)
		if err != nil {
			continue
		}

		var todo types.TodoItem
		if err := json.Unmarshal(val, &todo); err != nil {
			continue
		}

		todo.Key = key
		todos = append(todos, todo)
	}

	return todos, nil
}
