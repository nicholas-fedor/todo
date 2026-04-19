package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nicholas-fedor/todo/internal/handler"
	"github.com/nicholas-fedor/todo/pkg/pages"
)

func NotFoundMiddleware(c fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return handler.Render(c, pages.NotFound())
}
