package model

import "time"

type Tip struct {
	ID             string       `gorm:"primaryKey;type:varchar(26)" json:"id"`
	SenderID       string       `gorm:"type:varchar(26);not null;index" json:"sender_id"`
	Sender         User         `gorm:"foreignKey:SenderID" json:"-"`
	RecipientID    string       `gorm:"type:varchar(26);not null;index" json:"recipient_id"`
	Recipient      User         `gorm:"foreignKey:RecipientID" json:"-"`
	StaffProfileID string       `gorm:"type:varchar(26);not null;index" json:"staff_profile_id"`
	StaffProfile   StaffProfile `gorm:"foreignKey:StaffProfileID" json:"-"`
	Amount         int          `gorm:"not null" json:"amount"`
	Message        *string      `gorm:"type:text" json:"message,omitempty"`
	CreatedAt      time.Time    `gorm:"not null;autoCreateTime" json:"created_at"`
}

type PointTransaction struct {
	ID           string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	UserID       string    `gorm:"type:varchar(26);not null;index" json:"user_id"`
	User         User      `gorm:"foreignKey:UserID" json:"-"`
	Type         string    `gorm:"type:varchar(50);not null" json:"type"`
	Amount       int       `gorm:"not null" json:"amount"`
	BalanceAfter int       `gorm:"not null" json:"balance_after"`
	ReferenceID  *string   `gorm:"type:varchar(26)" json:"reference_id,omitempty"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}
