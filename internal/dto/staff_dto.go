package dto

import (
	"time"

	"staff-search-api/internal/model"
)

type CreateStaffProfileRequest struct {
	JobTitle       string  `json:"job_title"`
	JobCategory    string  `json:"job_category"`
	Location       *string `json:"location"`
	Bio            *string `json:"bio"`
	AcceptBookings *bool   `json:"accept_bookings"`
}

type UpdateStaffProfileRequest struct {
	JobTitle       *string  `json:"job_title"`
	JobCategory    *string  `json:"job_category"`
	Location       *string  `json:"location"`
	Latitude       *float64 `json:"latitude"`
	Longitude      *float64 `json:"longitude"`
	Bio            *string  `json:"bio"`
	AcceptBookings *bool    `json:"accept_bookings"`
	IsAvailable    *bool    `json:"is_available"`
}

type PortfolioPhotoResponse struct {
	ID             string    `json:"id"`
	StaffProfileID string    `json:"staff_profile_id"`
	PhotoURL       string    `json:"photo_url"`
	DisplayOrder   int       `json:"display_order"`
	CreatedAt      time.Time `json:"created_at"`
}

type AddPortfolioPhotoRequest struct {
	PhotoURL     string `json:"photo_url"`
	DisplayOrder *int   `json:"display_order"`
}

type PhotoOrder struct {
	ID    string `json:"id"`
	Order int    `json:"order"`
}

type ReorderPhotosRequest struct {
	PhotoOrders []PhotoOrder `json:"photo_orders"`
}

type StaffProfileResponse struct {
	ID                string                   `json:"id"`
	UserID            string                   `json:"user_id"`
	Name              string                   `json:"name"`
	AvatarURL         *string                  `json:"avatar_url,omitempty"`
	StaffNumber       string                   `json:"staff_number"`
	JobTitle          string                   `json:"job_title"`
	JobCategory       string                   `json:"job_category"`
	Location          *string                  `json:"location,omitempty"`
	Latitude          *float64                 `json:"latitude,omitempty"`
	Longitude         *float64                 `json:"longitude,omitempty"`
	Bio               *string                  `json:"bio,omitempty"`
	IntroVideoURL     *string                  `json:"intro_video_url,omitempty"`
	IsAvailable       bool                     `json:"is_available"`
	AcceptBookings    bool                     `json:"accept_bookings"`
	Rating            float32                  `json:"rating"`
	ReviewCount       int                      `json:"review_count"`
	FollowersCount    int                      `json:"followers_count"`
	TotalTipsReceived int                      `json:"total_tips_received"`
	PortfolioPhotos   []PortfolioPhotoResponse `json:"portfolio_photos"`
	PostMediaURLs     []string                 `json:"post_media_urls"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
}

func ToStaffProfileResponse(profile *model.StaffProfile) StaffProfileResponse {
	photos := make([]PortfolioPhotoResponse, 0, len(profile.PortfolioPhotos))
	for _, p := range profile.PortfolioPhotos {
		photos = append(photos, PortfolioPhotoResponse{
			ID:             p.ID,
			StaffProfileID: p.StaffProfileID,
			PhotoURL:       p.PhotoURL,
			DisplayOrder:   p.DisplayOrder,
			CreatedAt:      p.CreatedAt,
		})
	}
	return StaffProfileResponse{
		ID:                profile.ID,
		UserID:            profile.UserID,
		Name:              profile.User.Name,
		AvatarURL:         profile.User.AvatarURL,
		StaffNumber:       profile.StaffNumber,
		JobTitle:          profile.JobTitle,
		JobCategory:       profile.JobCategory,
		Location:          profile.Location,
		Latitude:          profile.Latitude,
		Longitude:         profile.Longitude,
		Bio:               profile.Bio,
		IntroVideoURL:     profile.IntroVideoURL,
		IsAvailable:       profile.IsAvailable,
		AcceptBookings:    profile.AcceptBookings,
		Rating:            profile.Rating,
		ReviewCount:       profile.ReviewCount,
		FollowersCount:    profile.FollowersCount,
		TotalTipsReceived: profile.TotalTipsReceived,
		PortfolioPhotos:   photos,
		CreatedAt:         profile.CreatedAt,
		UpdatedAt:         profile.UpdatedAt,
	}
}
