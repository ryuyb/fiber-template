package errors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"live-pilot/api/presenter"
)

func NotFound(format string, v ...string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    fiber.StatusNotFound,
		Message: fmt.Sprintf(format, v),
	}
}

func InternalServerError(format string, v ...string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    fiber.StatusInternalServerError,
		Message: fmt.Sprintf(format, v),
	}
}

func BadRequest(format string, v ...string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    fiber.StatusBadRequest,
		Message: fmt.Sprintf(format, v),
	}
}
