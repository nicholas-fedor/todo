package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/nicholas-fedor/todo/internal/models"
	"github.com/nicholas-fedor/todo/pkg/pages"
	"github.com/nicholas-fedor/todo/pkg/todo"
	"github.com/nicholas-fedor/todo/storage"
)

// ToggleHandler returns a handler that toggles a todo's completion status.
func ToggleHandler(store storage.Store) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		key := "todo:" + id

		val, err := store.Get(key)
		if err != nil {
			if err == storage.ErrNotFound {
				return c.Status(
					fiber.StatusNotFound).
					SendString("todo not found")
			}

			return c.Status(
				fiber.StatusInternalServerError).
				SendString("failed to get todo")
		}

		var todoItem models.TodoView
		if err := json.Unmarshal(val, &todoItem); err != nil {
			return c.Status(
				fiber.StatusInternalServerError).
				SendString("invalid data")
		}

		todoItem.Completed = !todoItem.Completed

		data, err := json.Marshal(todoItem)
		if err != nil {
			return c.Status(
				fiber.StatusInternalServerError).
				SendString("failed to marshal")
		}

		if err := store.Set(key, data); err != nil {
			return c.Status(
				fiber.StatusInternalServerError).
				SendString("failed to update")
		}

		todos, err := todo.ListTodos(store)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("failed to list todos")
		}

		return Render(c, pages.TodoList(todos))
	}
}
