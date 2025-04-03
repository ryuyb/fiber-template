package middleware

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddleware(a *fiber.App, logger *fiberzap.LoggerConfig) {
	a.Use(
		fiberzap.New(fiberzap.Config{
			Logger: logger.Logger(),
		}),
		compress.New(compress.Config{
			Level: compress.LevelDefault,
		}),
		cors.New(),
		recover.New(recover.Config{
			EnableStackTrace: false,
		}),
	)
}
