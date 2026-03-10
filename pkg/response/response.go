package response

import (
	"github.com/gofiber/fiber/v3"
)

// Standard error envelope — every non-2xx response uses this shape.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func Success(c fiber.Ctx, status int, data any) error {
	return c.Status(status).JSON(data)
}

func Error(c fiber.Ctx, status int, code string, message string) error {
	return c.Status(status).JSON(ErrorResponse{
		Error:   code,
		Message: message,
	})
}

func BadRequest(c fiber.Ctx, message string) error {
	return Error(c, fiber.StatusBadRequest, "bad_request", message)
}

func InvalidToken(c fiber.Ctx, message string) error {
	return Error(c, fiber.StatusBadRequest, "invalid_token", message)
}

func Unauthorized(c fiber.Ctx, message string) error {
	return Error(c, fiber.StatusUnauthorized, "unauthorized", message)
}

func Forbidden(c fiber.Ctx, message string) error {
	return Error(c, fiber.StatusForbidden, "forbidden", message)
}

func AccountDisabled(c fiber.Ctx) error {
	return Error(c, fiber.StatusForbidden, "account_disabled", "This account has been suspended.")
}

func NotFound(c fiber.Ctx, message string) error {
	return Error(c, fiber.StatusNotFound, "not_found", message)
}

func Conflict(c fiber.Ctx, message string) error {
	return Error(c, fiber.StatusConflict, "conflict", message)
}

func ValidationError(c fiber.Ctx, message string) error {
	return Error(c, fiber.StatusUnprocessableEntity, "validation_error", message)
}

func ServerError(c fiber.Ctx) error {
	return Error(c, fiber.StatusInternalServerError, "server_error", "An internal error occurred.")
}
