package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"staff-search-api/internal/config"
	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/service"
	"staff-search-api/pkg/response"
)

type StaffHandler struct {
	staffService     *service.StaffService
	portfolioService *service.StaffPortfolioService
}

func NewStaffHandler(staffService *service.StaffService, portfolioService *service.StaffPortfolioService) *StaffHandler {
	return &StaffHandler{staffService: staffService, portfolioService: portfolioService}
}

func (h *StaffHandler) ListStaff(c fiber.Ctx) error {
	category := c.Query("category", "")
	cursor := c.Query("cursor", "")
	limitStr := c.Query("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 20
	}
	if limit > 50 {
		limit = 50
	}

	result, err := h.staffService.ListStaff(c.Context(), category, cursor, limit)
	if err != nil {
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, result)
}

func (h *StaffHandler) GetJobCategories(c fiber.Ctx) error {
	return response.Success(c, fiber.StatusOK, config.JobCategories)
}

func (h *StaffHandler) CreateProfile(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	var req dto.CreateStaffProfileRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	if req.JobTitle == "" {
		return response.ValidationError(c, "Job title is required.")
	}
	if len(req.JobTitle) > 100 {
		return response.ValidationError(c, "Job title must be 100 characters or fewer.")
	}
	if !config.IsValidJobCategory(req.JobCategory) {
		return response.ValidationError(c, "Invalid job category.")
	}

	profile, err := h.staffService.CreateProfile(c.Context(), userID, req)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrConflict):
			return response.Conflict(c, "You already have a staff profile.")
		case errors.Is(err, model.ErrStaffNumberExhausted):
			return response.ServerError(c)
		default:
			return response.ServerError(c)
		}
	}

	return response.Success(c, fiber.StatusCreated, dto.ToStaffProfileResponse(profile))
}

func (h *StaffHandler) UpdateProfile(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	var req dto.UpdateStaffProfileRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	if req.JobTitle == nil && req.JobCategory == nil && req.Location == nil &&
		req.Latitude == nil && req.Longitude == nil &&
		req.Bio == nil && req.AcceptBookings == nil && req.IsAvailable == nil {
		return response.BadRequest(c, "At least one field is required.")
	}

	if req.JobCategory != nil && !config.IsValidJobCategory(*req.JobCategory) {
		return response.ValidationError(c, "Invalid job category.")
	}

	profile, err := h.staffService.UpdateProfile(c.Context(), userID, req)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return response.NotFound(c, "Staff profile not found.")
		}
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, dto.ToStaffProfileResponse(profile))
}

func (h *StaffHandler) GetMyProfile(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	profile, err := h.staffService.GetByUserID(c.Context(), userID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return response.NotFound(c, "Staff profile not found.")
		}
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, dto.ToStaffProfileResponse(profile))
}

func (h *StaffHandler) GetProfile(c fiber.Ctx) error {
	targetUserID := c.Params("userID")
	if targetUserID == "" {
		return response.BadRequest(c, "User ID is required.")
	}

	profile, err := h.staffService.GetByUserID(c.Context(), targetUserID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return response.NotFound(c, "Staff profile not found.")
		}
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, dto.ToStaffProfileResponse(profile))
}

func (h *StaffHandler) AddPortfolioPhoto(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	var req dto.AddPortfolioPhotoRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}
	if req.PhotoURL == "" {
		return response.BadRequest(c, "photo_url is required.")
	}

	photo, err := h.portfolioService.AddPhoto(c.Context(), userID, req)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return response.NotFound(c, "Staff profile not found.")
		case errors.Is(err, model.ErrPhotoLimitReached):
			return response.ValidationError(c, "Portfolio is full. Maximum 12 photos allowed.")
		case errors.Is(err, model.ErrValidation):
			return response.ValidationError(c, "Invalid photo URL.")
		}
		return response.ServerError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.PortfolioPhotoResponse{
		ID:             photo.ID,
		StaffProfileID: photo.StaffProfileID,
		PhotoURL:       photo.PhotoURL,
		DisplayOrder:   photo.DisplayOrder,
		CreatedAt:      photo.CreatedAt,
	})
}

func (h *StaffHandler) DeletePortfolioPhoto(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	photoID := c.Params("photoID")

	err := h.portfolioService.DeletePhoto(c.Context(), userID, photoID)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return response.NotFound(c, "Photo not found.")
		case errors.Is(err, model.ErrForbidden):
			return response.Forbidden(c, "You do not have permission to delete this photo.")
		}
		return response.ServerError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *StaffHandler) ReorderPortfolioPhotos(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	var req dto.ReorderPhotosRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}
	if len(req.PhotoOrders) == 0 {
		return response.BadRequest(c, "photo_orders is required.")
	}

	if err := h.portfolioService.ReorderPhotos(c.Context(), userID, req); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return response.NotFound(c, "Staff profile not found.")
		}
		return response.ServerError(c)
	}

	return c.JSON(fiber.Map{"message": "Order updated."})
}
