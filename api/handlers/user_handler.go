package handlers

import (
	"live-pilot/api/presenter/errors"
	"live-pilot/pkg/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{us: us}
}

// GetUser handles GET /user/:id
// @Summary Get user by ID
// @Description Get user details by user ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user/{id} [get]
func (u *UserHandler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return errors.BadRequest("invalid user id: %s", idStr)
	}
	return c.JSON(u.us.GetUser(uint32(id)))
}
