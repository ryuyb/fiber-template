package routes

import (
	"github.com/gofiber/fiber/v2"
	"live-pilot/api/handlers"
)

type AuthRoutes struct {
	h *handlers.AuthHandler
}

func NewAuthRoutes(h *handlers.AuthHandler) *AuthRoutes {
	return &AuthRoutes{h: h}
}

func (ur *AuthRoutes) ApplyRoutes(r fiber.Router) {
	g := r.Group("/auth")
	g.Post("/login", ur.h.Login)
}
