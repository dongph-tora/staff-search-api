package handler

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/service"
	"staff-search-api/pkg/response"
)

type AuthHandler struct {
	authService *service.AuthService
	userService *service.UserService
}

func NewAuthHandler(authService *service.AuthService, userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	if req.Email == "" || req.Password == "" {
		return response.BadRequest(c, "Email and password are required.")
	}

	resp, err := h.authService.Login(c.Context(), &req)
	if err != nil {
		return h.mapAuthError(c, err)
	}

	return response.Success(c, fiber.StatusOK, resp)
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	if req.Email == "" || req.Password == "" || req.Name == "" {
		return response.BadRequest(c, "Email, password, and name are required.")
	}

	if len(req.Password) < 6 {
		return response.ValidationError(c, "Password must be at least 6 characters.")
	}

	resp, err := h.authService.Register(c.Context(), &req)
	if err != nil {
		if errors.Is(err, model.ErrConflict) {
			return response.Conflict(c, "An account with this email already exists.")
		}
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusCreated, resp)
}

func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	var req dto.RefreshTokenRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	if req.RefreshToken == "" {
		return response.BadRequest(c, "Refresh token is required.")
	}

	resp, err := h.authService.RefreshToken(c.Context(), req.RefreshToken)
	if err != nil {
		return h.mapAuthError(c, err)
	}

	return response.Success(c, fiber.StatusOK, resp)
}

func (h *AuthHandler) Logout(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	_ = h.authService.Logout(c.Context(), userID)
	return response.Success(c, fiber.StatusOK, fiber.Map{"message": "Logged out successfully."})
}

func (h *AuthHandler) Me(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	user, err := h.userService.GetByID(c.Context(), userID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return response.NotFound(c, "User not found.")
		}
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, dto.UserResponse{
		ID:                    user.ID,
		Email:                 user.Email,
		Name:                  user.Name,
		AvatarURL:             user.AvatarURL,
		Role:                  user.Role,
		IsStaff:               user.IsStaff,
		IsStaffRegistered:     user.IsStaffRegistered,
		Points:                user.Points,
		PrivacyPolicyAccepted: user.PrivacyPolicyAccepted,
		CreatedAt:             user.CreatedAt,
	})
}

func (h *AuthHandler) AcceptPrivacyPolicy(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	var req dto.AcceptPrivacyPolicyRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	if req.Version == "" {
		return response.BadRequest(c, "Policy version is required.")
	}

	if err := h.userService.AcceptPrivacyPolicy(c.Context(), userID, req.Version); err != nil {
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, fiber.Map{"message": "Privacy policy accepted."})
}

func (h *AuthHandler) mapAuthError(c fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, model.ErrInvalidCredentials):
		return response.Unauthorized(c, "Invalid email or password.")
	case errors.Is(err, model.ErrInvalidToken):
		return response.InvalidToken(c, "Token is invalid or expired.")
	case errors.Is(err, model.ErrAccountDisabled):
		return response.AccountDisabled(c)
	default:
		return response.ServerError(c)
	}
}
