package model

import "time"

type LiveStream struct {
	ID               string       `gorm:"primaryKey;type:varchar(26)" json:"id"`
	StaffProfileID   string       `gorm:"type:varchar(26);not null;index" json:"staff_profile_id"`
	StaffProfile     StaffProfile `gorm:"foreignKey:StaffProfileID" json:"-"`
	AgoraChannelName string       `gorm:"type:varchar(255);not null" json:"agora_channel_name"`
	Status           string       `gorm:"type:varchar(20);not null;default:'live'" json:"status"`
	ViewerCount      int          `gorm:"not null;default:0" json:"viewer_count"`
	TipTotal         int          `gorm:"not null;default:0" json:"tip_total"`
	StartedAt        time.Time    `gorm:"not null;autoCreateTime" json:"started_at"`
	EndedAt          *time.Time   `json:"ended_at,omitempty"`
	CreatedAt        time.Time    `gorm:"not null;autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time    `gorm:"not null;autoUpdateTime" json:"updated_at"`
}
