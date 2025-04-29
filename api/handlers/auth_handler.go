package handlers

import (
	"live-pilot/api/presenter"
	"live-pilot/api/presenter/errors"
	"live-pilot/pkg/service"
	"live-pilot/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	us *service.UserService
}

func NewAuthHandler(us *service.UserService) *AuthHandler {
	return &AuthHandler{us: us}
}

// Login user login
//
//	@Summary		user login
//	@Description	User login with username and password returns an access token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			loginReq	body		presenter.UserLoginReq	true	"LoginInfo"
//	@Success		200			{object}	presenter.UserLoginResp
//	@Failure		400			{object}	presenter.ErrorResponse
//	@Router			/auth/login [post]
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
	return c.JSON(presenter.UserLoginResp{AccessToken: accessToken})
}
