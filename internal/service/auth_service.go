package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"

	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/repository"
	"staff-search-api/pkg/email"
	"staff-search-api/pkg/jwt"
	"staff-search-api/pkg/ulid"
)

type AuthService struct {
	userRepo          *repository.UserRepository
	refreshTokenRepo  *repository.RefreshTokenRepository
	jwtService        *jwt.Service
	passwordResetRepo *repository.PasswordResetRepository
	emailService      email.EmailSender
	appBaseURL        string
	googleClientID    string
}

func NewAuthService(
	userRepo *repository.UserRepository,
	refreshTokenRepo *repository.RefreshTokenRepository,
	jwtService *jwt.Service,
	passwordResetRepo *repository.PasswordResetRepository,
	emailService email.EmailSender,
	appBaseURL string,
	googleClientID string,
) *AuthService {
	return &AuthService{
		userRepo:          userRepo,
		refreshTokenRepo:  refreshTokenRepo,
		jwtService:        jwtService,
		passwordResetRepo: passwordResetRepo,
		emailService:      emailService,
		appBaseURL:        appBaseURL,
		googleClientID:    googleClientID,
	}
}

func (s *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, model.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("find user: %w", err)
	}

	if user.Status == "disabled" {
		return nil, model.ErrAccountDisabled
	}

	if user.PasswordHash == nil {
		return nil, model.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, model.ErrInvalidCredentials
	}

	if err := s.userRepo.UpdateLastLogin(ctx, user.ID); err != nil {
		return nil, fmt.Errorf("update last login: %w", err)
	}

	return s.BuildAuthResponse(ctx, user, false)
}

func (s *AuthService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	existing, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, fmt.Errorf("find user: %w", err)
	}
	if existing != nil {
		return nil, model.ErrConflict
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}
	hashStr := string(hash)

	now := time.Now()
	user := &model.User{
		ID:                ulid.New(),
		Email:             req.Email,
		PasswordHash:      &hashStr,
		Name:              req.Name,
		Role:              "user",
		IsStaff:           false,
		IsStaffRegistered: false,
		IsVerified:        false,
		AuthProvider:      "email",
		Status:            "active",
		Points:            0,
		PrivacyPolicyAccepted: false,
		LastLoginAt:       &now,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return s.BuildAuthResponse(ctx, user, true)
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshTokenStr string) (*dto.AuthResponse, error) {
	userID, err := s.jwtService.ValidateRefreshToken(refreshTokenStr)
	if err != nil {
		return nil, model.ErrInvalidToken
	}

	tokenHash := jwt.HashToken(refreshTokenStr)
	storedToken, err := s.refreshTokenRepo.FindByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, model.ErrInvalidToken
		}
		return nil, fmt.Errorf("find refresh token: %w", err)
	}

	// Rotate: delete old token
	if err := s.refreshTokenRepo.DeleteByID(ctx, storedToken.ID); err != nil {
		return nil, fmt.Errorf("delete old refresh token: %w", err)
	}

	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}

	if user.Status == "disabled" {
		return nil, model.ErrAccountDisabled
	}

	return s.BuildAuthResponse(ctx, user, false)
}

func (s *AuthService) Logout(ctx context.Context, userID string) error {
	return s.refreshTokenRepo.DeleteByUserID(ctx, userID)
}

func (s *AuthService) RequestPasswordReset(ctx context.Context, emailAddr string) error {
	user, err := s.userRepo.FindByEmail(ctx, emailAddr)
	if err != nil {
		// Do not leak information about whether email exists
		return nil
	}

	plaintextToken := ulid.New()
	hash := sha256.Sum256([]byte(plaintextToken))
	tokenHash := fmt.Sprintf("%x", hash)

	now := time.Now()
	token := &model.PasswordResetToken{
		ID:        ulid.New(),
		UserID:    user.ID,
		TokenHash: tokenHash,
		ExpiresAt: now.Add(1 * time.Hour),
		CreatedAt: now,
	}

	if err := s.passwordResetRepo.Create(ctx, token); err != nil {
		return fmt.Errorf("create password reset token: %w", err)
	}

	resetLink := s.appBaseURL + "/reset-password?token=" + plaintextToken
	_ = s.emailService.SendPasswordReset(ctx, emailAddr, resetLink)

	return nil
}

func (s *AuthService) ConfirmPasswordReset(ctx context.Context, plaintextToken, newPassword string) error {
	hash := sha256.Sum256([]byte(plaintextToken))
	tokenHash := fmt.Sprintf("%x", hash)

	token, err := s.passwordResetRepo.FindByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return model.ErrInvalidToken
		}
		return fmt.Errorf("find reset token: %w", err)
	}

	if time.Now().After(token.ExpiresAt) {
		return model.ErrTokenExpired
	}

	if token.UsedAt != nil {
		return model.ErrTokenUsed
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	hashedStr := string(hashed)

	if err := s.userRepo.UpdatePasswordHash(ctx, token.UserID, hashedStr); err != nil {
		return fmt.Errorf("update password: %w", err)
	}

	return s.passwordResetRepo.MarkUsed(ctx, token.ID)
}

func (s *AuthService) ChangePassword(ctx context.Context, userID, currentPassword, newPassword string) error {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return model.ErrInvalidCredentials
		}
		return fmt.Errorf("find user: %w", err)
	}

	if user.PasswordHash == nil {
		return model.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(currentPassword)); err != nil {
		return model.ErrInvalidCredentials
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	return s.userRepo.UpdatePasswordHash(ctx, userID, string(hashed))
}

func (s *AuthService) BuildAuthResponse(ctx context.Context, user *model.User, isNewUser bool) (*dto.AuthResponse, error) {
	pair, err := s.jwtService.GenerateTokenPair(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("generate token pair: %w", err)
	}

	// Store refresh token hash
	now := time.Now()
	rt := &model.RefreshToken{
		ID:        ulid.New(),
		UserID:    user.ID,
		TokenHash: jwt.HashToken(pair.RefreshToken),
		ExpiresAt: now.Add(30 * 24 * time.Hour),
		CreatedAt: now,
	}
	if err := s.refreshTokenRepo.Create(ctx, rt); err != nil {
		return nil, fmt.Errorf("store refresh token: %w", err)
	}

	return &dto.AuthResponse{
		AccessToken:  pair.AccessToken,
		RefreshToken: pair.RefreshToken,
		TokenType:    pair.TokenType,
		ExpiresIn:    pair.ExpiresIn,
		User: dto.UserResponse{
			ID:                    user.ID,
			Email:                 user.Email,
			Name:                  user.Name,
			PhoneNumber:           user.PhoneNumber,
			AvatarURL:             user.AvatarURL,
			Role:                  user.Role,
			IsStaff:               user.IsStaff,
			IsStaffRegistered:     user.IsStaffRegistered,
			IsNewUser:             isNewUser,
			Points:                user.Points,
			PrivacyPolicyAccepted: user.PrivacyPolicyAccepted,
			CreatedAt:             user.CreatedAt,
		},
	}, nil
}

type GoogleUserInfo struct {
	GoogleID  string
	Email     string
	Name      string
	AvatarURL string
	Verified  bool
}

func (s *AuthService) VerifyGoogleToken(ctx context.Context, idToken string) (*GoogleUserInfo, error) {
	payload, err := idtoken.Validate(ctx, idToken, s.googleClientID)
	if err != nil {
		return nil, model.ErrInvalidToken
	}

	sub, _ := payload.Claims["sub"].(string)
	if sub == "" {
		sub = payload.Subject
	}
	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	picture, _ := payload.Claims["picture"].(string)
	verified, _ := payload.Claims["email_verified"].(bool)

	return &GoogleUserInfo{
		GoogleID:  sub,
		Email:     email,
		Name:      name,
		AvatarURL: picture,
		Verified:  verified,
	}, nil
}
