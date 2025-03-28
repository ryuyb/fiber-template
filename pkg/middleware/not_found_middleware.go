package middleware

import (
	"live-pilot/api/presenter/errors"

	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(errors.NotFound("endpoint is not found"))
		},
	)
}
