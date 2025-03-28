//go:build wireinject
// +build wireinject

package main

import (
	"live-pilot/api/handlers"
	"live-pilot/api/routes"
	"live-pilot/pkg/conf"
	"live-pilot/pkg/repository"
	"live-pilot/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func wireApp(conf.AppConfig, *zap.Logger) (*fiber.App, func(), error) {
	panic(wire.Build(handlers.ProviderSet, routes.ProviderSet, service.ProviderSet, repository.ProviderSet, newFiberApp))
}
