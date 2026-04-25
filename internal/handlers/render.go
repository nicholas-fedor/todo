package handlers

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
)

func Render(c fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")

	err := component.Render(c.Context(), c.Response().BodyWriter())
	if err != nil {
		return fmt.Errorf("render component: %w", err)
	}

	return nil
}
