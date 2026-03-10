package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"staff-search-api/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, r.mapError(err)
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, r.mapError(err)
	}
	return &user, nil
}

func (r *UserRepository) FindByGoogleID(ctx context.Context, googleID string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("google_id = ?", googleID).First(&user).Error
	if err != nil {
		return nil, r.mapError(err)
	}
	return &user, nil
}

func (r *UserRepository) UpdateLastLogin(ctx context.Context, userID string) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", userID).
		Updates(map[string]any{
			"last_login_at": now,
			"updated_at":    now,
		}).Error
}

func (r *UserRepository) LinkGoogleProvider(ctx context.Context, userID, googleID string, avatarURL *string) error {
	updates := map[string]any{
		"google_id": googleID,
	}
	if avatarURL != nil {
		updates["avatar_url"] = *avatarURL
	}

	return r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", userID).
		Updates(updates).
		UpdateColumn("auth_provider", gorm.Expr(
			"CASE WHEN auth_provider = 'email' THEN 'google' ELSE 'multiple' END",
		)).Error
}

func (r *UserRepository) AcceptPrivacyPolicy(ctx context.Context, userID, version string) error {
	return r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", userID).
		Updates(map[string]any{
			"privacy_policy_accepted": true,
			"privacy_policy_version":  version,
		}).Error
}

func (r *UserRepository) mapError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.ErrNotFound
	}
	return err
}
