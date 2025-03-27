package handlers

import (
	"live-poilot/pkg/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{us: us}
}

func (u *UserHandler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "id is invalid"}
	}
	return c.JSON(u.us.GetUser(uint32(id)))
}
