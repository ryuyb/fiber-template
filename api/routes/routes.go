package routes

import "github.com/gofiber/fiber/v2"

type Routes interface {
	ApplyRoutes(fiber.Router)
}
