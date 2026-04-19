package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nicholas-fedor/todo/pkg/pages"
	"github.com/nicholas-fedor/todo/pkg/todo"
	"github.com/nicholas-fedor/todo/storage"
)

// HomeHandler returns a handler that renders the home page with all stored todos.
func HomeHandler(store storage.Store) fiber.Handler {
	return func(c fiber.Ctx) error {
		todos, err := todo.ListTodos(store)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("failed to list todos")
		}

		return Render(c, pages.Home(todos))
	}
}
