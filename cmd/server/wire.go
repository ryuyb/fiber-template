//go:build wireinject
// +build wireinject

package main

import (
	"live-poilot/api/handlers"
	"live-poilot/api/routes"
	"live-poilot/pkg/conf"
	"live-poilot/pkg/repository"
	"live-poilot/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func wireApp(conf.AppConfig, *zap.Logger) (*fiber.App, func(), error) {
	panic(wire.Build(handlers.ProviderSet, routes.ProviderSet, service.ProviderSet, repository.ProviderSet, newFiberApp))
}
