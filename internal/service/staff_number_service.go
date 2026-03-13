package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"staff-search-api/internal/model"
	"staff-search-api/internal/repository"
)

type StaffNumberService struct {
	staffRepo *repository.StaffRepository
}

func NewStaffNumberService(staffRepo *repository.StaffRepository) *StaffNumberService {
	return &StaffNumberService{staffRepo: staffRepo}
}

func (s *StaffNumberService) Generate(ctx context.Context) (string, error) {
	const maxAttempts = 10
	for attempt := 0; attempt < maxAttempts; attempt++ {
		n, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			return "", fmt.Errorf("generate random number: %w", err)
		}
		number := fmt.Sprintf("%06d", n.Int64())
		exists, err := s.staffRepo.StaffNumberExists(ctx, number)
		if err != nil {
			return "", fmt.Errorf("check staff number: %w", err)
		}
		if !exists {
			return number, nil
		}
	}
	return "", model.ErrStaffNumberExhausted
}
