package todo

import (
	"encoding/json"

	"github.com/nicholas-fedor/todo/internal/models"
	"github.com/nicholas-fedor/todo/storage"
)

// ListTodos fetches and decodes all todos from the store.
func ListTodos(store storage.Store) ([]models.TodoView, error) {
	keys, err := store.List()
	if err != nil {
		return nil, err
	}

	todos := make([]models.TodoView, 0, len(keys))
	for _, key := range keys {
		val, err := store.Get(key)
		if err != nil {
			continue
		}

		var todo models.TodoView
		if err := json.Unmarshal(val, &todo); err != nil {
			continue
		}

		todo.Key = key
		todos = append(todos, todo)
	}

	return todos, nil
}
