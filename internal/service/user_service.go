package service

import (
	"context"
	"fmt"

	"staff-search-api/internal/model"
	"staff-search-api/internal/repository"
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
