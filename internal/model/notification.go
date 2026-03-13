package model

import "time"

type Notification struct {
	ID        string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	UserID    string    `gorm:"type:varchar(26);not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Body      string    `gorm:"type:text;not null" json:"body"`
	Data      *string   `gorm:"type:jsonb" json:"data,omitempty"`
	IsRead    bool      `gorm:"not null;default:false" json:"is_read"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}
