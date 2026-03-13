package model

import "time"

type Service struct {
	ID              string       `gorm:"primaryKey;type:varchar(26)" json:"id"`
	StaffProfileID  string       `gorm:"type:varchar(26);not null;index" json:"staff_profile_id"`
	StaffProfile    StaffProfile `gorm:"foreignKey:StaffProfileID;constraint:OnDelete:CASCADE" json:"-"`
	Name            string       `gorm:"type:varchar(255);not null" json:"name"`
	Description     *string      `gorm:"type:text" json:"description,omitempty"`
	Price           float64      `gorm:"type:decimal(10,2);not null" json:"price"`
	DurationMinutes int          `gorm:"not null" json:"duration_minutes"`
	IsActive        bool         `gorm:"not null;default:true" json:"is_active"`
	CreatedAt       time.Time    `gorm:"not null;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time    `gorm:"not null;autoUpdateTime" json:"updated_at"`
}

type Booking struct {
	ID             string       `gorm:"primaryKey;type:varchar(26)" json:"id"`
	UserID         string       `gorm:"type:varchar(26);not null;index" json:"user_id"`
	User           User         `gorm:"foreignKey:UserID" json:"-"`
	StaffProfileID string       `gorm:"type:varchar(26);not null;index" json:"staff_profile_id"`
	StaffProfile   StaffProfile `gorm:"foreignKey:StaffProfileID" json:"-"`
	ServiceID      *string      `gorm:"type:varchar(26)" json:"service_id,omitempty"`
	Service        *Service     `gorm:"foreignKey:ServiceID" json:"-"`
	Status         string       `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	Note           *string      `gorm:"type:text" json:"note,omitempty"`
	ScheduledAt    *time.Time   `json:"scheduled_at,omitempty"`
	CreatedAt      time.Time    `gorm:"not null;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time    `gorm:"not null;autoUpdateTime" json:"updated_at"`
}
