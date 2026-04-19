package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/nicholas-fedor/todo/pkg/pages"
	"github.com/nicholas-fedor/todo/pkg/todo"
	"github.com/nicholas-fedor/todo/storage"
)

// DeleteHandler returns a handler that deletes a todo by id.
func DeleteHandler(store storage.Store) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		key := "todo:" + id

		if err := store.Delete(key); err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("failed to delete")
		}

		todos, err := todo.ListTodos(store)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("failed to list todos")
		}

		return Render(c, pages.TodoList(todos))
	}
}

// generateID returns a UUID-based unique identifier
func generateID() string {
	return uuid.New().String()
}
