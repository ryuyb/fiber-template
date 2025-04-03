//go:build wireinject
// +build wireinject

package main

import (
	"live-pilot/api/handlers"
	"live-pilot/api/routes"
	"live-pilot/pkg/conf"
	"live-pilot/pkg/middleware"
	"live-pilot/pkg/repository"
	"live-pilot/pkg/service"

	"github.com/google/wire"
)

func wireApp(conf.AppConfig) (*FiberApp, func(), error) {
	panic(wire.Build(
		handlers.ProviderSet,
		routes.ProviderSet,
		service.ProviderSet,
		repository.ProviderSet,
		middleware.ProviderSet,
		newFiberApp,
	))
}
