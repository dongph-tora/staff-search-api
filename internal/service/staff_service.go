package service

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/repository"
	"staff-search-api/pkg/ulid"
)

type StaffService struct {
	staffRepo          *repository.StaffRepository
	staffNumberService *StaffNumberService
	userRepo           *repository.UserRepository
	postRepo           *repository.PostRepository
	db                 *gorm.DB
}

func NewStaffService(
	staffRepo *repository.StaffRepository,
	staffNumberService *StaffNumberService,
	userRepo *repository.UserRepository,
	postRepo *repository.PostRepository,
	db *gorm.DB,
) *StaffService {
	return &StaffService{
		staffRepo:          staffRepo,
		staffNumberService: staffNumberService,
		userRepo:           userRepo,
		postRepo:           postRepo,
		db:                 db,
	}
}

func (s *StaffService) CreateProfile(ctx context.Context, userID string, req dto.CreateStaffProfileRequest) (*model.StaffProfile, error) {
	exists, err := s.staffRepo.ExistsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("check existing profile: %w", err)
	}
	if exists {
		return nil, model.ErrConflict
	}

	staffNumber, err := s.staffNumberService.Generate(ctx)
	if err != nil {
		return nil, fmt.Errorf("generate staff number: %w", err)
	}

	acceptBookings := true
	if req.AcceptBookings != nil {
		acceptBookings = *req.AcceptBookings
	}

	now := time.Now()
	profile := &model.StaffProfile{
		ID:             ulid.New(),
		UserID:         userID,
		StaffNumber:    staffNumber,
		JobTitle:       req.JobTitle,
		JobCategory:    req.JobCategory,
		Location:       req.Location,
		Bio:            req.Bio,
		IsAvailable:    false,
		AcceptBookings: acceptBookings,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(profile).Error; err != nil {
			return err
		}
		return tx.Model(&model.User{}).
			Where("id = ?", userID).
			Updates(map[string]any{
				"is_staff_registered": true,
				"role":                "staff",
				"updated_at":          now,
			}).Error
	})
	if err != nil {
		return nil, fmt.Errorf("create staff profile: %w", err)
	}

	return s.staffRepo.FindByUserID(ctx, userID)
}

func (s *StaffService) UpdateProfile(ctx context.Context, userID string, req dto.UpdateStaffProfileRequest) (*model.StaffProfile, error) {
	_, err := s.staffRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	fields := map[string]any{"updated_at": time.Now()}
	if req.JobTitle != nil {
		fields["job_title"] = *req.JobTitle
	}
	if req.JobCategory != nil {
		fields["job_category"] = *req.JobCategory
	}
	if req.Location != nil {
		fields["location"] = *req.Location
	}
	if req.Latitude != nil {
		fields["latitude"] = *req.Latitude
	}
	if req.Longitude != nil {
		fields["longitude"] = *req.Longitude
	}
	if req.Bio != nil {
		fields["bio"] = *req.Bio
	}
	if req.AcceptBookings != nil {
		fields["accept_bookings"] = *req.AcceptBookings
	}
	if req.IsAvailable != nil {
		fields["is_available"] = *req.IsAvailable
	}

	if err := s.staffRepo.Update(ctx, userID, fields); err != nil {
		return nil, fmt.Errorf("update staff profile: %w", err)
	}

	return s.staffRepo.FindByUserID(ctx, userID)
}

func (s *StaffService) GetByUserID(ctx context.Context, userID string) (*model.StaffProfile, error) {
	return s.staffRepo.FindByUserID(ctx, userID)
}

type StaffListResponse struct {
	Staff      []dto.StaffProfileResponse `json:"staff"`
	NextCursor *string                    `json:"next_cursor"`
	HasMore    bool                       `json:"has_more"`
}

func (s *StaffService) ListStaff(ctx context.Context, category, cursor string, limit int) (StaffListResponse, error) {
	profiles, nextCursor, hasMore, err := s.staffRepo.List(ctx, category, cursor, limit)
	if err != nil {
		return StaffListResponse{}, fmt.Errorf("list staff: %w", err)
	}

	// Collect user IDs to fetch latest post media in one query
	userIDs := make([]string, 0, len(profiles))
	for _, p := range profiles {
		userIDs = append(userIDs, p.UserID)
	}
	postMediaMap, _ := s.postRepo.GetLatestMediaURLsByAuthorIDs(ctx, userIDs, 5)

	items := make([]dto.StaffProfileResponse, 0, len(profiles))
	for _, p := range profiles {
		r := dto.ToStaffProfileResponse(p)
		r.PostMediaURLs = postMediaMap[p.UserID]
		if r.PostMediaURLs == nil {
			r.PostMediaURLs = []string{}
		}
		items = append(items, r)
	}

	var nc *string
	if nextCursor != "" {
		nc = &nextCursor
	}

	return StaffListResponse{Staff: items, NextCursor: nc, HasMore: hasMore}, nil
}
