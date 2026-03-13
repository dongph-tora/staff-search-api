package model

import "time"

type HeadhuntOffer struct {
	ID             string       `gorm:"primaryKey;type:varchar(26)" json:"id"`
	CompanyID      string       `gorm:"type:varchar(26);not null;index" json:"company_id"`
	Company        User         `gorm:"foreignKey:CompanyID" json:"-"`
	StaffProfileID string       `gorm:"type:varchar(26);not null;index" json:"staff_profile_id"`
	StaffProfile   StaffProfile `gorm:"foreignKey:StaffProfileID" json:"-"`
	Title          string       `gorm:"type:varchar(255);not null" json:"title"`
	Message        string       `gorm:"type:text;not null" json:"message"`
	SalaryMin      *float64     `gorm:"type:decimal(10,2)" json:"salary_min,omitempty"`
	SalaryMax      *float64     `gorm:"type:decimal(10,2)" json:"salary_max,omitempty"`
	Status         string       `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	ExpiresAt      *time.Time   `json:"expires_at,omitempty"`
	CreatedAt      time.Time    `gorm:"not null;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time    `gorm:"not null;autoUpdateTime" json:"updated_at"`
}
