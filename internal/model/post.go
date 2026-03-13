package model

import "time"

type Post struct {
	ID            string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	AuthorID      string    `gorm:"type:varchar(26);not null;index" json:"author_id"`
	Author        User      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE" json:"-"`
	Content       *string   `gorm:"type:text" json:"content,omitempty"`
	MediaURL      *string   `gorm:"type:text" json:"media_url,omitempty"`
	MediaType     *string   `gorm:"type:varchar(10)" json:"media_type,omitempty"`
	LikesCount    int       `gorm:"not null;default:0" json:"likes_count"`
	CommentsCount int       `gorm:"not null;default:0" json:"comments_count"`
	IsActive      bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt     time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"not null;autoUpdateTime" json:"updated_at"`
}

type Comment struct {
	ID        string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	PostID    string    `gorm:"type:varchar(26);not null;index" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"-"`
	AuthorID  string    `gorm:"type:varchar(26);not null;index" json:"author_id"`
	Author    User      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE" json:"-"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	IsActive  bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updated_at"`
}

type Like struct {
	ID        string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	UserID    string    `gorm:"type:varchar(26);not null;index" json:"user_id"`
	PostID    string    `gorm:"type:varchar(26);not null;index" json:"post_id"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}

func (Like) TableName() string { return "likes" }

type Follow struct {
	ID         string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	FollowerID string    `gorm:"type:varchar(26);not null;index" json:"follower_id"`
	FollowedID string    `gorm:"type:varchar(26);not null;index" json:"followed_id"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}

type Story struct {
	ID         string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	AuthorID   string    `gorm:"type:varchar(26);not null;index" json:"author_id"`
	Author     User      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE" json:"-"`
	MediaURL   string    `gorm:"type:text;not null" json:"media_url"`
	MediaType  string    `gorm:"type:varchar(10);not null" json:"media_type"`
	ViewsCount int       `gorm:"not null;default:0" json:"views_count"`
	ExpiresAt  time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}
