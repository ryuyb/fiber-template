package errors

import (
	"github.com/gofiber/fiber/v2"
	"live-pilot/api/presenter"
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

func Error(code int, message string) presenter.ErrorResponse {
	return presenter.ErrorResponse{
		Code:    code,
		Message: message,
	}
}
