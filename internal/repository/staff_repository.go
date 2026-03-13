package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
)

type StaffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{db: db}
}

func (r *StaffRepository) List(ctx context.Context, category, cursor string, limit int) ([]*model.StaffProfile, string, bool, error) {
	query := r.db.WithContext(ctx).
		Preload("User").
		Preload("PortfolioPhotos", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC")
		})

	if category != "" {
		query = query.Where("job_category = ?", category)
	}
	if cursor != "" {
		query = query.Where("id < ?", cursor)
	}

	var profiles []*model.StaffProfile
	err := query.Order("id DESC").Limit(limit + 1).Find(&profiles).Error
	if err != nil {
		return nil, "", false, err
	}

	hasMore := len(profiles) > limit
	if hasMore {
		profiles = profiles[:limit]
	}

	var nextCursor string
	if len(profiles) > 0 {
		nextCursor = profiles[len(profiles)-1].ID
	}

	return profiles, nextCursor, hasMore, nil
}

func (r *StaffRepository) StaffNumberExists(ctx context.Context, number string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.StaffProfile{}).
		Where("staff_number = ?", number).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *StaffRepository) FindByUserID(ctx context.Context, userID string) (*model.StaffProfile, error) {
	var profile model.StaffProfile
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("PortfolioPhotos", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC")
		}).
		Where("user_id = ?", userID).
		First(&profile).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrNotFound
	}
	return &profile, err
}

func (r *StaffRepository) FindByID(ctx context.Context, id string) (*model.StaffProfile, error) {
	var profile model.StaffProfile
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("PortfolioPhotos", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC")
		}).
		Where("id = ?", id).
		First(&profile).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrNotFound
	}
	return &profile, err
}

func (r *StaffRepository) Create(ctx context.Context, profile *model.StaffProfile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *StaffRepository) Update(ctx context.Context, userID string, fields map[string]any) error {
	return r.db.WithContext(ctx).
		Model(&model.StaffProfile{}).
		Where("user_id = ?", userID).
		Updates(fields).Error
}

func (r *StaffRepository) ExistsByUserID(ctx context.Context, userID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.StaffProfile{}).
		Where("user_id = ?", userID).Count(&count).Error
	return count > 0, err
}

func (r *StaffRepository) CountPortfolioPhotos(ctx context.Context, staffProfileID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.StaffPortfolioPhoto{}).
		Where("staff_profile_id = ?", staffProfileID).
		Count(&count).Error
	return count, err
}

func (r *StaffRepository) AddPortfolioPhoto(ctx context.Context, photo *model.StaffPortfolioPhoto) error {
	return r.db.WithContext(ctx).Create(photo).Error
}

func (r *StaffRepository) GetPortfolioPhotoByID(ctx context.Context, photoID string) (*model.StaffPortfolioPhoto, error) {
	var photo model.StaffPortfolioPhoto
	err := r.db.WithContext(ctx).Where("id = ?", photoID).First(&photo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrNotFound
	}
	return &photo, err
}

func (r *StaffRepository) DeletePortfolioPhoto(ctx context.Context, photoID string) error {
	return r.db.WithContext(ctx).Where("id = ?", photoID).Delete(&model.StaffPortfolioPhoto{}).Error
}

func (r *StaffRepository) MaxDisplayOrder(ctx context.Context, staffProfileID string) (int, error) {
	var maxOrder *int
	r.db.WithContext(ctx).Model(&model.StaffPortfolioPhoto{}).
		Where("staff_profile_id = ?", staffProfileID).
		Select("MAX(display_order)").
		Scan(&maxOrder)
	if maxOrder == nil {
		return 0, nil
	}
	return *maxOrder, nil
}

func (r *StaffRepository) UpdateDisplayOrders(ctx context.Context, staffProfileID string, photoOrders []dto.PhotoOrder) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, po := range photoOrders {
			if err := tx.Model(&model.StaffPortfolioPhoto{}).
				Where("id = ? AND staff_profile_id = ?", po.ID, staffProfileID).
				Update("display_order", po.Order).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
