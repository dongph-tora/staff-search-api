package dto

import "time"

type CreatePostRequest struct {
	Content   *string `json:"content"`
	MediaURL  *string `json:"media_url"`
	MediaType *string `json:"media_type"`
}

type AuthorInfo struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

type PostResponse struct {
	ID            string     `json:"id"`
	AuthorID      string     `json:"author_id"`
	Author        AuthorInfo `json:"author"`
	Content       *string    `json:"content,omitempty"`
	MediaURL      *string    `json:"media_url,omitempty"`
	MediaType     *string    `json:"media_type,omitempty"`
	LikesCount    int        `json:"likes_count"`
	CommentsCount int        `json:"comments_count"`
	IsLiked       bool       `json:"is_liked"`
	CreatedAt     time.Time  `json:"created_at"`
}

type FeedResponse struct {
	Posts      []PostResponse `json:"posts"`
	NextCursor *string        `json:"next_cursor"`
	HasMore    bool           `json:"has_more"`
}
