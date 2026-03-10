package dto

import "time"

// --- Request DTOs ---

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type GoogleSignInRequest struct {
	IDToken string `json:"id_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type PasswordResetRequest struct {
	Email string `json:"email"`
}

type PasswordResetConfirmRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type AcceptPrivacyPolicyRequest struct {
	Version string `json:"version"`
}

// --- Response DTOs ---

type AuthResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	TokenType    string       `json:"token_type"`
	ExpiresIn    int          `json:"expires_in"`
	User         UserResponse `json:"user"`
}

type UserResponse struct {
	ID                    string     `json:"id"`
	Email                 string     `json:"email"`
	Name                  string     `json:"name"`
	AvatarURL             *string    `json:"avatar_url"`
	Role                  string     `json:"role"`
	IsStaff               bool       `json:"is_staff"`
	IsStaffRegistered     bool       `json:"is_staff_registered"`
	IsNewUser             bool       `json:"is_new_user,omitempty"`
	Points                int        `json:"points"`
	PrivacyPolicyAccepted bool       `json:"privacy_policy_accepted"`
	CreatedAt             time.Time  `json:"created_at"`
}
