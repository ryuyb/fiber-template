package errors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"live-pilot/api/presenter"
)

func NotFound(message string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    fiber.StatusNotFound,
		Message: fmt.Sprintf(message),
	}
}

func InternalServerError(message string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    fiber.StatusInternalServerError,
		Message: fmt.Sprintf(message),
	}
}

func BadRequest(message string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    fiber.StatusBadRequest,
		Message: fmt.Sprintf(message),
	}
}
