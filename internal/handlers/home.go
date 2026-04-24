package handlers

import (
	"github.com/gofiber/fiber/v3"

	"github.com/nicholas-fedor/todo/internal/app"
	"github.com/nicholas-fedor/todo/internal/storage"
	"github.com/nicholas-fedor/todo/internal/web/pages"
)

// HomeHandler returns a handler that renders the home page with all stored todos.
func HomeHandler(store storage.Store) fiber.Handler {
	return func(c fiber.Ctx) error {
		todos, err := app.ListTodos(store)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("failed to list todos")
		}

		return Render(c, pages.Home(todos))
	}
}
