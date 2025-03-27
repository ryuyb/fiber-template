package routes

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserRoutes, ProvideRoutes)

func ProvideRoutes(ur *UserRoutes) []Routes {
	return []Routes{ur}
}
