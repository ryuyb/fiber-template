package handlers

import (
	"fmt"
	"live-pilot/api/presenter"
	"live-pilot/api/presenter/errors"
	"live-pilot/pkg/ent"
	"live-pilot/pkg/service"
	"strconv"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{us: us}
}

// GetUser handles GET /user/:id
//
//	@Summary		Get user by ID
//	@Description	Get user details by user ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int	true	"User ID"
//	@Success		200			{object}	ent.User
//	@Failure		400,404,500	{object}	presenter.ErrorResponse
//	@Router			/user/{id} [get]
func (u *UserHandler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return errors.BadRequest(fmt.Sprintf("invalid user id: %s", idStr))
	}
	user, err := u.us.GetUser(uint32(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.NotFound(fmt.Sprintf("user not found with id: %d", id))
		}
		log.Errorf("error getting user: %#v", err)
		return errors.InternalServerError(fmt.Sprintf("error getting user with id: %d", id))
	}
	return c.JSON(user)
}

// SaveUser Create or update users
//
//	@Summary		Create or update users
//	@Description	Determine whether the user is a new or updated user based on the ID in the request body
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			saveUserReq	body		presenter.SaveUserReq	true	"UserInfo"
//	@Success		200			{object}	ent.User
//	@Failure		400,404,500	{object}	presenter.ErrorResponse
//	@Router			/users [post]
func (u *UserHandler) SaveUser(c *fiber.Ctx) error {
	var saveUserReq presenter.SaveUserReq
	if err := c.BodyParser(&saveUserReq); err != nil {
		log.Errorf("error parsing request body: %#v", err)
		return errors.BadRequest("invalid user request")
	}
	user, err := u.us.SaveUser(saveUserReq)
	if err != nil {
		if ent.IsConstraintError(err) {
			return errors.BadRequest("user is already exists")
		}
		if ent.IsNotFound(err) {
			return errors.NotFound("user not found")
		}
		return errors.InternalServerError(fmt.Sprintf("occurred database error: %v", err))
	}
	return c.JSON(user)
}
