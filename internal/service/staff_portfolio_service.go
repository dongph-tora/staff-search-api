package service

import (
	"context"
	"strings"

	"staff-search-api/internal/config"
	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/repository"
	pkg_storage "staff-search-api/pkg/storage"
	"staff-search-api/pkg/ulid"
)

type StaffPortfolioService struct {
	staffRepo     *repository.StaffRepository
	storageClient pkg_storage.StorageClient
	cfg           *config.Config
}

func NewStaffPortfolioService(staffRepo *repository.StaffRepository, storageClient pkg_storage.StorageClient, cfg *config.Config) *StaffPortfolioService {
	return &StaffPortfolioService{staffRepo: staffRepo, storageClient: storageClient, cfg: cfg}
}

func (s *StaffPortfolioService) AddPhoto(ctx context.Context, userID string, req dto.AddPortfolioPhotoRequest) (*model.StaffPortfolioPhoto, error) {
	profile, err := s.staffRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err // ErrNotFound propagates
	}

	count, err := s.staffRepo.CountPortfolioPhotos(ctx, profile.ID)
	if err != nil {
		return nil, err
	}
	if count >= 12 {
		return nil, model.ErrPhotoLimitReached
	}

	if !strings.HasPrefix(req.PhotoURL, s.cfg.StoragePublicURL) {
		return nil, model.ErrValidation
	}

	displayOrder := int(count) + 1
	if req.DisplayOrder != nil {
		displayOrder = *req.DisplayOrder
	}

	photo := &model.StaffPortfolioPhoto{
		ID:             ulid.New(),
		StaffProfileID: profile.ID,
		PhotoURL:       req.PhotoURL,
		DisplayOrder:   displayOrder,
	}

	if err := s.staffRepo.AddPortfolioPhoto(ctx, photo); err != nil {
		return nil, err
	}
	return photo, nil
}

func (s *StaffPortfolioService) DeletePhoto(ctx context.Context, userID, photoID string) error {
	photo, err := s.staffRepo.GetPortfolioPhotoByID(ctx, photoID)
	if err != nil {
		return err
	}

	profile, err := s.staffRepo.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if photo.StaffProfileID != profile.ID {
		return model.ErrForbidden
	}

	// Extract file key (remove public URL prefix)
	fileKey := strings.TrimPrefix(photo.PhotoURL, s.cfg.StoragePublicURL)
	fileKey = strings.TrimPrefix(fileKey, "/")

	// Delete from storage (best effort — don't fail if already gone)
	_ = s.storageClient.DeleteObject(ctx, fileKey)

	return s.staffRepo.DeletePortfolioPhoto(ctx, photoID)
}

func (s *StaffPortfolioService) ReorderPhotos(ctx context.Context, userID string, req dto.ReorderPhotosRequest) error {
	profile, err := s.staffRepo.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}
	return s.staffRepo.UpdateDisplayOrders(ctx, profile.ID, req.PhotoOrders)
}
