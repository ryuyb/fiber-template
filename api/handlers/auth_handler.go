package handlers

import (
	"github.com/gofiber/fiber/v2"
	"live-pilot/api/presenter"
	"live-pilot/api/presenter/errors"
	"live-pilot/pkg/service"
)

type AuthHandler struct {
	us *service.UserService
}

func NewAuthHandler(us *service.UserService) *AuthHandler {
	return &AuthHandler{us: us}
}

func (u *AuthHandler) Login(c *fiber.Ctx) error {
	loginReq := presenter.UserLoginReq{}
	err := c.BodyParser(loginReq)
	if err != nil {
		return errors.BadRequest("Invalid login information")
	}
	passed, err := u.us.CheckPassword(loginReq)
	if !passed || err != nil {
		return errors.BadRequest("Incorrect username or password")
	}
	return c.JSON(fiber.Map{})
}
