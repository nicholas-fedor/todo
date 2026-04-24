package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	"github.com/nicholas-fedor/todo/internal/app"
	"github.com/nicholas-fedor/todo/internal/storage"
	"github.com/nicholas-fedor/todo/internal/web/pages"
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

		todos, err := app.ListTodos(store)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("failed to list todos")
		}

		return Render(c, pages.TodoList(todos))
	}
}

// generateID returns a UUID-based unique identifier.
func generateID() string {
	return uuid.New().String()
}
