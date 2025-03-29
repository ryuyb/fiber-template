package middleware

import (
	"live-pilot/api/presenter"
	"live-pilot/api/presenter/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
)

const (
	defaultErrorMsg = "Internal Server Error"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var errorResponse presenter.ErrorResponse

	if e, ok := lo.ErrorsAs[*fiber.Error](err); ok {
		errorResponse = presenter.ErrorResponse{
			Code:    e.Code,
			Message: e.Message,
		}
	} else if e, ok := lo.ErrorsAs[presenter.ErrorResponse](err); ok {
		errorResponse = e
	} else {
		errorResponse = errors.InternalServerError(defaultErrorMsg)
		log.Panicf("ErrorHandler: %s | Path: %s | Method: %s | Error: %v",
			defaultErrorMsg,
			ctx.Path(),
			ctx.Method(),
			err,
		)
	}

	return ctx.Status(errorResponse.Code).JSON(errorResponse)
}
