package handler

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/service"
	"staff-search-api/pkg/response"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) UpdateProfile(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	var req dto.UpdateUserRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	user, err := h.userService.UpdateProfile(c.Context(), userID, req)
	if err != nil {
		if errors.Is(err, model.ErrValidation) {
			msg := strings.TrimPrefix(err.Error(), model.ErrValidation.Error()+": ")
			return response.ValidationError(c, msg)
		}
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, dto.UserResponse{
		ID:                    user.ID,
		Email:                 user.Email,
		Name:                  user.Name,
		PhoneNumber:           user.PhoneNumber,
		AvatarURL:             user.AvatarURL,
		Role:                  user.Role,
		IsStaff:               user.IsStaff,
		IsStaffRegistered:     user.IsStaffRegistered,
		Points:                user.Points,
		PrivacyPolicyAccepted: user.PrivacyPolicyAccepted,
		CreatedAt:             user.CreatedAt,
	})
}
