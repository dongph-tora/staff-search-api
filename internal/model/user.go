package model

import "time"

type User struct {
	ID                    string     `gorm:"primaryKey;type:varchar(26)" json:"id"`
	Email                 string     `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
	PasswordHash          *string    `gorm:"type:varchar(255)" json:"-"`
	Name                  string     `gorm:"type:varchar(255);not null" json:"name"`
	PhoneNumber           *string    `gorm:"type:varchar(20)" json:"phone_number,omitempty"`
	AvatarURL             *string    `gorm:"type:text" json:"avatar_url,omitempty"`
	Bio                   *string    `gorm:"type:text" json:"bio,omitempty"`
	Role                  string     `gorm:"type:varchar(20);not null;default:'user'" json:"role"`
	IsStaff               bool       `gorm:"not null;default:false" json:"is_staff"`
	IsStaffRegistered     bool       `gorm:"not null;default:false" json:"is_staff_registered"`
	IsVerified            bool       `gorm:"not null;default:false" json:"is_verified"`
	GoogleID              *string    `gorm:"uniqueIndex;type:varchar(255)" json:"-"`
	AppleID               *string    `gorm:"uniqueIndex;type:varchar(255)" json:"-"`
	AuthProvider          string     `gorm:"type:varchar(50);not null;default:'email'" json:"auth_provider"`
	Status                string     `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
	Points                int        `gorm:"not null;default:0" json:"points"`
	PrivacyPolicyAccepted bool       `gorm:"not null;default:false" json:"privacy_policy_accepted"`
	PrivacyPolicyVersion  *string    `gorm:"type:varchar(20)" json:"privacy_policy_version,omitempty"`
	LastLoginAt           *time.Time `json:"last_login_at,omitempty"`
	CreatedAt             time.Time  `gorm:"not null;autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time  `gorm:"not null;autoUpdateTime" json:"updated_at"`
}

type RefreshToken struct {
	ID        string    `gorm:"primaryKey;type:varchar(26)" json:"id"`
	UserID    string    `gorm:"type:varchar(26);not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	TokenHash string    `gorm:"uniqueIndex;type:varchar(64);not null" json:"-"`
	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}
