package routes

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserRoutes, NewAuthRoutes, ProvideRoutes)

func ProvideRoutes(user *UserRoutes, auth *AuthRoutes) []Routes {
	return []Routes{user, auth}
}
