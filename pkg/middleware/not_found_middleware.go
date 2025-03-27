package middleware

import (
	"live-poilot/api/presenter"

	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(presenter.ErrorResponse{Code: fiber.StatusNotFound, Message: "endpoint is not found"})
		},
	)
}
