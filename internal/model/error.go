package model

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrConflict        = errors.New("conflict")
	ErrInvalidToken    = errors.New("invalid token")
	ErrAccountDisabled = errors.New("account disabled")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrValidation           = errors.New("validation error")
	ErrStaffNumberExhausted = errors.New("staff_number_exhausted")
	ErrPhotoLimitReached    = errors.New("photo_limit_reached")
)
