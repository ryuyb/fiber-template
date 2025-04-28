package handlers

import (
	"github.com/gofiber/fiber/v2"
	"live-pilot/api/presenter"
	"live-pilot/api/presenter/errors"
	"live-pilot/pkg/service"
	"live-pilot/pkg/utils"
)

type AuthHandler struct {
	us *service.UserService
}

func NewAuthHandler(us *service.UserService) *AuthHandler {
	return &AuthHandler{us: us}
}

func (u *AuthHandler) Login(c *fiber.Ctx) error {
	loginReq := presenter.UserLoginReq{}
	err := c.BodyParser(&loginReq)
	if err != nil {
		return errors.BadRequest("Invalid login information")
	}
	user, err := u.us.CheckPassword(loginReq)
	if user == nil || err != nil {
		return errors.BadRequest("Incorrect username or password")
	}
	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return errors.BadRequest("Generate access token failed")
	}
	return c.JSON(fiber.Map{
		"accessToken": accessToken,
	})
}
