// Package middleware provides HTTP middleware components for the todo application.
// It includes middleware for handling common HTTP concerns like 404 page rendering.
package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"github.com/nicholas-fedor/todo/internal/handlers"
	"github.com/nicholas-fedor/todo/internal/web/pages"
)

func NotFoundMiddleware(c fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)

	err := handlers.Render(c, pages.NotFound())
	if err != nil {
		return fmt.Errorf("not found middleware: %w", err)
	}

	return nil
}
