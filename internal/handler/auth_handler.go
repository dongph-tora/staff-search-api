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

func (h *AuthHandler) RequestPasswordReset(c fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
	}
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}
	if body.Email == "" {
		return response.BadRequest(c, "Email is required.")
	}

	_ = h.authService.RequestPasswordReset(c.Context(), body.Email)
	return response.Success(c, fiber.StatusOK, fiber.Map{
		"message": "If that email is registered, you will receive a reset link shortly.",
	})
}

func (h *AuthHandler) ConfirmPasswordReset(c fiber.Ctx) error {
	var body struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}
	if body.Token == "" || body.NewPassword == "" {
		return response.BadRequest(c, "Token and new_password are required.")
	}
	if len(body.NewPassword) < 6 {
		return response.ValidationError(c, "Password must be at least 6 characters.")
	}

	err := h.authService.ConfirmPasswordReset(c.Context(), body.Token, body.NewPassword)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrInvalidToken):
			return response.InvalidToken(c, "This reset link is invalid.")
		case errors.Is(err, model.ErrTokenExpired):
			return response.InvalidToken(c, "This reset link has expired. Please request a new one.")
		case errors.Is(err, model.ErrTokenUsed):
			return response.InvalidToken(c, "This reset link has already been used.")
		default:
			return response.ServerError(c)
		}
	}

	return response.Success(c, fiber.StatusOK, fiber.Map{
		"message": "Password updated successfully.",
	})
}

func (h *AuthHandler) GoogleSignIn(c fiber.Ctx) error {
	var body struct {
		IDToken string `json:"id_token"`
	}
	if err := c.Bind().JSON(&body); err != nil || body.IDToken == "" {
		return response.BadRequest(c, "id_token is required.")
	}

	info, err := h.authService.VerifyGoogleToken(c.Context(), body.IDToken)
	if err != nil {
		return response.InvalidToken(c, "Google authentication failed. Please try again.")
	}

	user, isNew, err := h.userService.UpsertGoogleUser(c.Context(), info.GoogleID, info.Email, info.Name, info.AvatarURL, info.Verified)
	if err != nil {
		if errors.Is(err, model.ErrAccountDisabled) {
			return response.AccountDisabled(c)
		}
		return response.ServerError(c)
	}

	if user.Status == "disabled" {
		return response.AccountDisabled(c)
	}

	resp, err := h.authService.BuildAuthResponse(c.Context(), user, isNew)
	if err != nil {
		return response.ServerError(c)
	}

	statusCode := fiber.StatusOK
	if isNew {
		statusCode = fiber.StatusCreated
	}
	return response.Success(c, statusCode, resp)
}

func (h *AuthHandler) ChangePassword(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	var req struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}
	if req.CurrentPassword == "" || req.NewPassword == "" {
		return response.BadRequest(c, "current_password and new_password are required.")
	}
	if len(req.NewPassword) < 8 {
		return response.BadRequest(c, "New password must be at least 8 characters.")
	}

	if err := h.authService.ChangePassword(c.Context(), userID, req.CurrentPassword, req.NewPassword); err != nil {
		return h.mapAuthError(c, err)
	}

	return response.Success(c, fiber.StatusOK, fiber.Map{"message": "Password changed successfully."})
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
