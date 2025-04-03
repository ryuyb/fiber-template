package main

import (
	"flag"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"live-pilot/api/routes"
	_ "live-pilot/docs"
	"live-pilot/pkg/conf"
	"live-pilot/pkg/middleware"
	"live-pilot/pkg/utils"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "config.yml", "config path, eg: -conf config.yml")
}

type FiberApp struct {
	app    *fiber.App
	logger *fiberzap.LoggerConfig
}

func newFiberApp(logger *fiberzap.LoggerConfig, appRoutes []routes.Routes) *FiberApp {
	app := fiber.New(fiber.Config{
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
		ErrorHandler: middleware.ErrorHandler,
	})

	log.SetLogger(logger)

	middleware.FiberMiddleware(app, logger)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	api := app.Group("/api")
	for _, appRoute := range appRoutes {
		appRoute.ApplyRoutes(api)
	}

	middleware.NotFoundMiddleware(app)

	return &FiberApp{
		app:    app,
		logger: logger,
	}
}

//	@title			LivePilot
//	@version		1.0
//	@description	This is an API for Live

// @host		127.0.0.1:8000
// @BasePath	/api
func main() {
	flag.Parse()
	cfg := conf.New(flagconf)

	fiberApp, cleanup, err := wireApp(cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	defer func(logger *fiberzap.LoggerConfig) {
		err := logger.Sync()
		if err != nil {
			panic(fmt.Errorf("failed to sync logger: %v", err))
		}
	}(fiberApp.logger)

	utils.StartServerWithGracefulShutdown(cfg, fiberApp.app)
}
