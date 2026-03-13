package handler

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/service"
	"staff-search-api/pkg/response"
)

type MediaHandler struct {
	mediaService    *service.MediaService
	localUploadService *service.LocalUploadService
}

func NewMediaHandler(mediaService *service.MediaService, localUploadService *service.LocalUploadService) *MediaHandler {
	return &MediaHandler{mediaService: mediaService, localUploadService: localUploadService}
}

func (h *MediaHandler) UploadFile(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	folder := c.Query("folder", "posts")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return response.BadRequest(c, "file field is required.")
	}

	if h.localUploadService == nil {
		return response.BadRequest(c, "Direct upload is only available in local/dev mode.")
	}

	result, err := h.localUploadService.Save(c.Context(), userID, folder, fileHeader)
	if err != nil {
		if errors.Is(err, model.ErrValidation) {
			msg := strings.TrimPrefix(err.Error(), model.ErrValidation.Error()+": ")
			return response.ValidationError(c, msg)
		}
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, result)
}

func (h *MediaHandler) GenerateUploadURL(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	var req dto.UploadURLRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	if req.FileName == "" || req.ContentType == "" || req.Folder == "" {
		return response.BadRequest(c, "file_name, content_type, and folder are required.")
	}

	if len(req.FileName) > 255 {
		return response.ValidationError(c, "file_name must be 255 characters or fewer.")
	}

	resp, err := h.mediaService.GenerateUploadURL(c.Context(), userID, req)
	if err != nil {
		if errors.Is(err, model.ErrValidation) {
			msg := strings.TrimPrefix(err.Error(), model.ErrValidation.Error()+": ")
			return response.ValidationError(c, msg)
		}
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, resp)
}

func (h *MediaHandler) DeleteFile(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	var req dto.DeleteMediaRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	if req.FileKey == "" {
		return response.BadRequest(c, "file_key is required.")
	}

	// Verify ownership: file key must contain /userID/
	if !strings.Contains(req.FileKey, "/"+userID+"/") {
		return response.Forbidden(c, "You do not have permission to delete this file.")
	}

	err := h.mediaService.DeleteFile(c.Context(), req.FileKey)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return response.NotFound(c, "File not found.")
		}
		return response.ServerError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
