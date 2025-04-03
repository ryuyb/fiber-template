package main

import (
	"flag"
	"live-pilot/api/routes"
	"live-pilot/pkg/conf"
	"live-pilot/pkg/middleware"
	"live-pilot/pkg/utils"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "live-pilot/docs"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "config.yml", "config path, eg: -conf config.yml")
}

func newFiberApp(logger *zap.Logger, appRoutes []routes.Routes) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
		ErrorHandler: middleware.ErrorHandler,
	})

	middleware.FiberMiddleware(app, logger)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	api := app.Group("/api")
	for _, appRoute := range appRoutes {
		appRoute.ApplyRoutes(api)
	}

	middleware.NotFoundMiddleware(app)

	return app
}

//	@title			LivePilot
//	@version		1.0
//	@description	This is an API for Live

// @host		127.0.0.1:8000
// @BasePath	/api
func main() {
	cfg := conf.New(flagconf)

	logger, _ := zap.NewDevelopment(
		zap.AddCallerSkip(3),
		zap.AddStacktrace(zapcore.PanicLevel),
	)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)

	app, cleanup, err := wireApp(cfg, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	utils.StartServerWithGracefulShutdown(cfg, app)
}
