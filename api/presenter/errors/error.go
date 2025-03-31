package errors

import (
	"live-pilot/api/presenter"

	"github.com/gofiber/fiber/v2"
)

func NotFound(message string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    fiber.StatusNotFound,
		Message: message,
	}
}

func InternalServerError(message string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    fiber.StatusInternalServerError,
		Message: message,
	}
}

func BadRequest(message string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    fiber.StatusBadRequest,
		Message: message,
	}
}
