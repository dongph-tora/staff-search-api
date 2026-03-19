package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"staff-search-api/internal/model"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *PostRepository) GetByID(ctx context.Context, postID string) (*model.Post, error) {
	var post model.Post
	err := r.db.WithContext(ctx).
		Preload("Author").
		Where("id = ? AND is_active = true", postID).
		First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetFeed(ctx context.Context, cursor string, limit int, category string) ([]*model.Post, string, bool, error) {
	query := r.db.WithContext(ctx).
		Preload("Author").
		Where("posts.is_active = true")

	if category != "" {
		query = query.
			Joins("JOIN staff_profiles ON staff_profiles.user_id = posts.author_id").
			Where("staff_profiles.job_category = ?", category)
	}

	if cursor != "" {
		t, err := time.Parse(time.RFC3339Nano, cursor)
		if err == nil {
			query = query.Where("posts.created_at < ?", t)
		}
	}

	var posts []*model.Post
	err := query.Order("posts.created_at DESC, posts.id DESC").Limit(limit + 1).Find(&posts).Error
	if err != nil {
		return nil, "", false, err
	}

	hasMore := len(posts) > limit
	if hasMore {
		posts = posts[:limit]
	}

	var nextCursor string
	if len(posts) > 0 {
		nextCursor = posts[len(posts)-1].CreatedAt.UTC().Format(time.RFC3339Nano)
	}

	return posts, nextCursor, hasMore, nil
}

func (r *PostRepository) GetByAuthor(ctx context.Context, authorID, cursor string, limit int) ([]*model.Post, string, bool, error) {
	query := r.db.WithContext(ctx).
		Preload("Author").
		Where("author_id = ? AND is_active = true", authorID)

	if cursor != "" {
		t, err := time.Parse(time.RFC3339Nano, cursor)
		if err == nil {
			query = query.Where("created_at < ?", t)
		}
	}

	var posts []*model.Post
	err := query.Order("created_at DESC, id DESC").Limit(limit + 1).Find(&posts).Error
	if err != nil {
		return nil, "", false, err
	}

	hasMore := len(posts) > limit
	if hasMore {
		posts = posts[:limit]
	}

	var nextCursor string
	if len(posts) > 0 {
		nextCursor = posts[len(posts)-1].CreatedAt.UTC().Format(time.RFC3339Nano)
	}

	return posts, nextCursor, hasMore, nil
}

// GetLatestMediaURLsByAuthorIDs returns up to `perAuthor` media URLs per author, ordered newest-first.
func (r *PostRepository) GetLatestMediaURLsByAuthorIDs(ctx context.Context, authorIDs []string, perAuthor int) (map[string][]string, error) {
	if len(authorIDs) == 0 {
		return map[string][]string{}, nil
	}

	type row struct {
		AuthorID string
		MediaURL string
	}

	// Rank posts per author and take top N
	var rows []row
	err := r.db.WithContext(ctx).Raw(`
		SELECT author_id, media_url
		FROM (
			SELECT author_id, media_url,
				   ROW_NUMBER() OVER (PARTITION BY author_id ORDER BY created_at DESC) AS rn
			FROM posts
			WHERE author_id IN ? AND media_url IS NOT NULL AND media_url <> '' AND is_active = true
		) ranked
		WHERE rn <= ?
	`, authorIDs, perAuthor).Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string, len(authorIDs))
	for _, r := range rows {
		result[r.AuthorID] = append(result[r.AuthorID], r.MediaURL)
	}
	return result, nil
}

func (r *PostRepository) GetLikedPostIDs(ctx context.Context, userID string, postIDs []string) (map[string]bool, error) {
	if len(postIDs) == 0 {
		return map[string]bool{}, nil
	}
	var likes []model.Like
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND post_id IN ?", userID, postIDs).
		Find(&likes).Error
	if err != nil {
		return nil, err
	}
	result := make(map[string]bool, len(likes))
	for _, l := range likes {
		result[l.PostID] = true
	}
	return result, nil
}
