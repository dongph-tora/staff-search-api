package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/repository"
	"staff-search-api/pkg/jwt"
	"staff-search-api/pkg/ulid"
)

type AuthService struct {
	userRepo         *repository.UserRepository
	refreshTokenRepo *repository.RefreshTokenRepository
	jwtService       *jwt.Service
}

func NewAuthService(
	userRepo *repository.UserRepository,
	refreshTokenRepo *repository.RefreshTokenRepository,
	jwtService *jwt.Service,
) *AuthService {
	return &AuthService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtService:       jwtService,
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

	return s.buildAuthResponse(ctx, user, false)
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

	return s.buildAuthResponse(ctx, user, true)
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

	return s.buildAuthResponse(ctx, user, false)
}

func (s *AuthService) Logout(ctx context.Context, userID string) error {
	return s.refreshTokenRepo.DeleteByUserID(ctx, userID)
}

func (s *AuthService) buildAuthResponse(ctx context.Context, user *model.User, isNewUser bool) (*dto.AuthResponse, error) {
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
