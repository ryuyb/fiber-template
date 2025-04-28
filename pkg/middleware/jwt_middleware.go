package middleware

import (
	"context"
	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"live-pilot/api/presenter/errors"
	"live-pilot/pkg/conf"
	"live-pilot/pkg/repository"
	"live-pilot/pkg/utils"
	"strconv"
)

const (
	JwtContextKey  = "jwt"
	UserContextKey = "user"
)

func JwtProtected() fiber.Handler {
	config := jwtMiddleware.Config{
		SigningKey:     jwtMiddleware.SigningKey{Key: []byte(conf.GetConfig().Jwt.SigningKey)},
		ContextKey:     JwtContextKey,
		Claims:         &utils.JwtClaims{},
		SuccessHandler: jwtSuccessHandler,
		ErrorHandler:   jwtErrorHandler,
	}
	return jwtMiddleware.New(config)
}

func jwtSuccessHandler(c *fiber.Ctx) error {
	token, ok := c.Context().UserValue(JwtContextKey).(*jwt.Token)
	if !ok {
		return errors.InternalServerError("Failed to parse token")
	}
	if !token.Valid {
		return errors.Unauthorized("Invalid token")
	}
	subject, err := token.Claims.GetSubject()
	if err != nil {
		log.Errorf("Failed to parse token subject: %v", err)
		return errors.InternalServerError("Failed to parse token")
	}
	userId, err := strconv.Atoi(subject)
	if err != nil {
		log.Errorf("Failed to parse token subject: %v", err)
	}
	repo := repository.GetRepository()
	user, err := repo.GetClient().User.Get(context.Background(), uint32(userId))
	if err != nil {
		return errors.BadRequest("User not found")
	}
	c.Context().SetUserValue(UserContextKey, user)
	return c.Next()
}

func jwtErrorHandler(_ *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return errors.BadRequest(err.Error())
	}
	return errors.Unauthorized(err.Error())
}
