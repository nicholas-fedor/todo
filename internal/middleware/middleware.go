package middleware

import (
	"github.com/gofiber/fiber/v3"

	"github.com/nicholas-fedor/todo/internal/handlers"
	"github.com/nicholas-fedor/todo/internal/web/pages"
)

func NotFoundMiddleware(c fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)

	return handlers.Render(c, pages.NotFound())
}
