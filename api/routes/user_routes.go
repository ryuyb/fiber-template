package routes

import (
	"live-pilot/api/handlers"
	"live-pilot/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	h *handlers.UserHandler
}

func NewUserRoutes(h *handlers.UserHandler) *UserRoutes {
	return &UserRoutes{h: h}
}

func (ur *UserRoutes) ApplyRoutes(r fiber.Router) {
	g := r.Group("/user")
	g.Get("/:id", middleware.JwtProtected(), ur.h.GetUser)
	g.Post("/", ur.h.SaveUser)
}
