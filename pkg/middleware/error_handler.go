package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"live-pilot/api/presenter"
	"live-pilot/api/presenter/errors"
)

const (
	contentTypeHeader = fiber.HeaderContentType
	jsonContentType   = fiber.MIMEApplicationJSONCharsetUTF8
	defaultErrorMsg   = "Internal Server Error"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var errorResponse presenter.ErrorResponse

	switch e := err.(type) {
	case *fiber.Error:
		errorResponse = presenter.ErrorResponse{
			Code:    e.Code,
			Message: e.Message,
		}
	case presenter.ErrorResponse:
		errorResponse = e
	default:
		errorResponse = errors.InternalServerError(defaultErrorMsg)
		log.Errorf("ErrorHandler: %s | Path: %s | Method: %s | Error: %v",
			defaultErrorMsg,
			ctx.Path(),
			ctx.Method(),
			err,
		)
	}

	ctx.Set(contentTypeHeader, jsonContentType)
	return ctx.Status(errorResponse.Code).JSON(errorResponse)
}
