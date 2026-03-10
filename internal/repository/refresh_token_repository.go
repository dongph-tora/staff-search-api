package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"staff-search-api/internal/model"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Create(ctx context.Context, token *model.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *RefreshTokenRepository) FindByTokenHash(ctx context.Context, hash string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	err := r.db.WithContext(ctx).
		Where("token_hash = ? AND expires_at > NOW()", hash).
		First(&token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return &token, nil
}

func (r *RefreshTokenRepository) DeleteByID(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.RefreshToken{}, "id = ?", id).Error
}

func (r *RefreshTokenRepository) DeleteByUserID(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).Delete(&model.RefreshToken{}, "user_id = ?", userID).Error
}
