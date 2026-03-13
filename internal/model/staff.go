package model

import "time"

type StaffProfile struct {
	ID                string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	UserID            string    `gorm:"type:varchar(26);not null;uniqueIndex" json:"user_id"`
	User              User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	StaffNumber       string    `gorm:"type:varchar(6);not null;uniqueIndex" json:"staff_number"`
	JobTitle          string    `gorm:"type:varchar(100);not null" json:"job_title"`
	JobCategory       string    `gorm:"type:varchar(50);not null" json:"job_category"`
	Location          *string   `gorm:"type:varchar(255)" json:"location,omitempty"`
	Latitude          *float64  `gorm:"type:decimal(10,7)" json:"latitude,omitempty"`
	Longitude         *float64  `gorm:"type:decimal(10,7)" json:"longitude,omitempty"`
	Bio               *string   `gorm:"type:text" json:"bio,omitempty"`
	IntroVideoURL     *string   `gorm:"type:text" json:"intro_video_url,omitempty"`
	IsAvailable       bool      `gorm:"not null;default:false" json:"is_available"`
	AcceptBookings    bool      `gorm:"not null;default:true" json:"accept_bookings"`
	Rating            float32   `gorm:"type:decimal(3,2);not null;default:0.00" json:"rating"`
	ReviewCount       int       `gorm:"not null;default:0" json:"review_count"`
	FollowersCount    int       `gorm:"not null;default:0" json:"followers_count"`
	TotalTipsReceived int                   `gorm:"not null;default:0" json:"total_tips_received"`
	PortfolioPhotos   []StaffPortfolioPhoto `gorm:"foreignKey:StaffProfileID;constraint:OnDelete:CASCADE" json:"-"`
	CreatedAt         time.Time             `gorm:"not null;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time             `gorm:"not null;autoUpdateTime" json:"updated_at"`
}

type StaffPortfolioPhoto struct {
	ID             string       `gorm:"primaryKey;type:varchar(26)" json:"id"`
	StaffProfileID string       `gorm:"type:varchar(26);not null;index" json:"staff_profile_id"`
	StaffProfile   StaffProfile `gorm:"foreignKey:StaffProfileID;constraint:OnDelete:CASCADE" json:"-"`
	PhotoURL       string       `gorm:"type:text;not null" json:"photo_url"`
	DisplayOrder   int          `gorm:"not null;default:0" json:"display_order"`
	CreatedAt      time.Time    `gorm:"not null;autoCreateTime" json:"created_at"`
}

type StaffSocialLink struct {
	ID             string       `gorm:"primaryKey;type:varchar(26)" json:"id"`
	StaffProfileID string       `gorm:"type:varchar(26);not null;index" json:"staff_profile_id"`
	StaffProfile   StaffProfile `gorm:"foreignKey:StaffProfileID;constraint:OnDelete:CASCADE" json:"-"`
	Platform       string       `gorm:"type:varchar(50);not null" json:"platform"`
	URL            string       `gorm:"type:text;not null" json:"url"`
	CreatedAt      time.Time    `gorm:"not null;autoCreateTime" json:"created_at"`
}
