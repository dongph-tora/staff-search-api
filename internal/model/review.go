package model

import "time"

type Review struct {
	ID         string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	BookingID  string    `gorm:"type:varchar(26);not null;uniqueIndex" json:"booking_id"`
	Booking    Booking   `gorm:"foreignKey:BookingID" json:"-"`
	ReviewerID string    `gorm:"type:varchar(26);not null;index" json:"reviewer_id"`
	Reviewer   User      `gorm:"foreignKey:ReviewerID" json:"-"`
	RevieweeID string    `gorm:"type:varchar(26);not null;index" json:"reviewee_id"`
	Reviewee   User      `gorm:"foreignKey:RevieweeID" json:"-"`
	Rating     int16     `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Comment    *string   `gorm:"type:text" json:"comment,omitempty"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}
