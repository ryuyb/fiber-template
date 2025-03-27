package middleware

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func FiberMiddleware(a *fiber.App, logger *zap.Logger) {
	a.Use(
		fiberzap.New(fiberzap.Config{
			Logger: logger,
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
