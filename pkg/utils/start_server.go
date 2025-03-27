package utils

import (
	"live-poilot/pkg/conf"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func StartServer(cfg conf.AppConfig, a *fiber.App) {
	if err := a.Listen(cfg.Server.Addr); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}
}

func StartServerWithGracefulShutdown(cfg conf.AppConfig, a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Fatalf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	StartServer(cfg, a)

	<-idleConnsClosed
}
