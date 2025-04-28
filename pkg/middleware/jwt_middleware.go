package middleware

import (
	"github.com/gofiber/fiber/v2"
	"live-pilot/api/presenter/errors"
	"live-pilot/pkg/conf"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
)

func JwtProtected() fiber.Handler {
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(conf.GetConfig().Jwt.SigningKey)},
		ContextKey:   "user",
		ErrorHandler: jwtErrorHandler,
	}

	return jwtMiddleware.New(config)
}

func jwtErrorHandler(_ *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return errors.BadRequest(err.Error())
	}
	return errors.Unauthorized(err.Error())
}
