package service

import (
	"context"
	"fmt"
	"time"

	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/repository"
	"staff-search-api/pkg/ulid"
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) CreatePost(ctx context.Context, userID string, req dto.CreatePostRequest) (*dto.PostResponse, error) {
	now := time.Now()
	post := &model.Post{
		ID:        ulid.New(),
		AuthorID:  userID,
		Content:   req.Content,
		MediaURL:  req.MediaURL,
		MediaType: req.MediaType,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.postRepo.Create(ctx, post); err != nil {
		return nil, fmt.Errorf("create post: %w", err)
	}

	full, err := s.postRepo.GetByID(ctx, post.ID)
	if err != nil {
		return nil, fmt.Errorf("get created post: %w", err)
	}

	resp := toPostResponse(full, false)
	return &resp, nil
}

func (s *PostService) GetFeed(ctx context.Context, userID, cursor string, limit int, category string) (dto.FeedResponse, error) {
	posts, nextCursor, hasMore, err := s.postRepo.GetFeed(ctx, cursor, limit, category)
	if err != nil {
		return dto.FeedResponse{}, fmt.Errorf("get feed: %w", err)
	}

	postIDs := make([]string, 0, len(posts))
	for _, p := range posts {
		postIDs = append(postIDs, p.ID)
	}

	likedMap, err := s.postRepo.GetLikedPostIDs(ctx, userID, postIDs)
	if err != nil {
		likedMap = map[string]bool{}
	}

	responses := make([]dto.PostResponse, 0, len(posts))
	for _, p := range posts {
		responses = append(responses, toPostResponse(p, likedMap[p.ID]))
	}

	var nc *string
	if nextCursor != "" {
		nc = &nextCursor
	}

	return dto.FeedResponse{
		Posts:      responses,
		NextCursor: nc,
		HasMore:    hasMore,
	}, nil
}

func (s *PostService) GetMyPosts(ctx context.Context, userID, cursor string, limit int) (dto.FeedResponse, error) {
	posts, nextCursor, hasMore, err := s.postRepo.GetByAuthor(ctx, userID, cursor, limit)
	if err != nil {
		return dto.FeedResponse{}, fmt.Errorf("get my posts: %w", err)
	}

	postIDs := make([]string, 0, len(posts))
	for _, p := range posts {
		postIDs = append(postIDs, p.ID)
	}
	likedMap, _ := s.postRepo.GetLikedPostIDs(ctx, userID, postIDs)

	responses := make([]dto.PostResponse, 0, len(posts))
	for _, p := range posts {
		responses = append(responses, toPostResponse(p, likedMap[p.ID]))
	}

	var nc *string
	if nextCursor != "" {
		nc = &nextCursor
	}

	return dto.FeedResponse{Posts: responses, NextCursor: nc, HasMore: hasMore}, nil
}

func (s *PostService) GetByID(ctx context.Context, postID string) (*dto.PostResponse, error) {
	post, err := s.postRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	resp := toPostResponse(post, false)
	return &resp, nil
}

func toPostResponse(p *model.Post, isLiked bool) dto.PostResponse {
	return dto.PostResponse{
		ID:       p.ID,
		AuthorID: p.AuthorID,
		Author: dto.AuthorInfo{
			ID:        p.Author.ID,
			Name:      p.Author.Name,
			AvatarURL: p.Author.AvatarURL,
		},
		Content:       p.Content,
		MediaURL:      p.MediaURL,
		MediaType:     p.MediaType,
		LikesCount:    p.LikesCount,
		CommentsCount: p.CommentsCount,
		IsLiked:       isLiked,
		CreatedAt:     p.CreatedAt,
	}
}
