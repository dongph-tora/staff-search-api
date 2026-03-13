package service

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/repository"
	"staff-search-api/pkg/ulid"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}
	return user, nil
}

func (s *UserService) AcceptPrivacyPolicy(ctx context.Context, userID, version string) error {
	return s.userRepo.AcceptPrivacyPolicy(ctx, userID, version)
}

func (s *UserService) UpdateProfile(ctx context.Context, userID string, req dto.UpdateUserRequest) (*model.User, error) {
	fields := map[string]any{}

	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			return nil, fmt.Errorf("%w: Name cannot be empty.", model.ErrValidation)
		}
		if len(name) > 100 {
			return nil, fmt.Errorf("%w: Name must be 100 characters or fewer.", model.ErrValidation)
		}
		fields["name"] = name
	}

	if req.Bio != nil {
		bio := strings.TrimSpace(*req.Bio)
		if len(bio) > 500 {
			return nil, fmt.Errorf("%w: Bio must be 500 characters or fewer.", model.ErrValidation)
		}
		fields["bio"] = bio
	}

	if req.Phone != nil {
		phone := strings.TrimSpace(*req.Phone)
		if phone != "" {
			matched, _ := regexp.MatchString(`^\+?[0-9]{7,15}$`, phone)
			if !matched {
				return nil, fmt.Errorf("%w: Enter a valid phone number.", model.ErrValidation)
			}
		}
		fields["phone_number"] = phone
	}

	if len(fields) > 0 {
		fields["updated_at"] = time.Now()
		if err := s.userRepo.Update(ctx, userID, fields); err != nil {
			return nil, fmt.Errorf("update user: %w", err)
		}
	}

	return s.userRepo.FindByID(ctx, userID)
}

func (s *UserService) UpsertGoogleUser(ctx context.Context, googleID, email, name, avatarURL string, verified bool) (*model.User, bool, error) {
	// Try find by google_id
	user, err := s.userRepo.FindByGoogleID(ctx, googleID)
	if err == nil {
		_ = s.userRepo.UpdateLastLogin(ctx, user.ID)
		return user, false, nil
	}

	// Try find by email
	user, err = s.userRepo.FindByEmail(ctx, email)
	if err == nil {
		if user.Status == "disabled" {
			return nil, false, model.ErrAccountDisabled
		}
		_ = s.userRepo.LinkGoogleProvider(ctx, user.ID, googleID, &avatarURL)
		_ = s.userRepo.UpdateLastLogin(ctx, user.ID)
		return user, false, nil
	}

	// Create new user
	now := time.Now()
	newUser := &model.User{
		ID:           ulid.New(),
		Email:        email,
		Name:         name,
		AvatarURL:    &avatarURL,
		GoogleID:     &googleID,
		AuthProvider: "google",
		Role:         "user",
		Status:       "active",
		IsVerified:   verified,
		Points:       0,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if err := s.userRepo.Create(ctx, newUser); err != nil {
		return nil, false, fmt.Errorf("create google user: %w", err)
	}
	return newUser, true, nil
}
