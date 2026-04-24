package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"

	"github.com/nicholas-fedor/todo/internal/storage"
	"github.com/nicholas-fedor/todo/internal/types"
)

// CreateHandler returns a handler that creates a new todo from JSON body.
func CreateHandler(store storage.Store) fiber.Handler {
	return func(c fiber.Ctx) error {
		title := c.FormValue("title")
		if title == "" {
			return c.Status(
				fiber.StatusBadRequest).
				SendString("title is required")
		}

		id := generateID()
		key := "todo:" + id

		todo := types.TodoItem{Key: key, Title: title}

		data, err := json.Marshal(todo)
		if err != nil {
			return c.Status(
				fiber.StatusInternalServerError).
				SendString("failed to marshal")
		}

		if err := store.Set(key, data); err != nil {
			return c.Status(
				fiber.StatusInternalServerError).
				SendString("failed to save")
		}

		return HomeHandler(store)(c)
	}
}
