package handler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/internal/service"
	"staff-search-api/pkg/response"
)

type PostHandler struct {
	postService      *service.PostService
	storagePublicURL string
}

func NewPostHandler(postService *service.PostService, storagePublicURL string) *PostHandler {
	return &PostHandler{postService: postService, storagePublicURL: storagePublicURL}
}

func (h *PostHandler) CreatePost(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	var req dto.CreatePostRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "Invalid request body.")
	}

	hasContent := req.Content != nil && strings.TrimSpace(*req.Content) != ""
	hasMedia := req.MediaURL != nil && *req.MediaURL != ""
	if !hasContent && !hasMedia {
		return response.BadRequest(c, "At least a caption or media is required.")
	}

	if hasMedia && req.MediaType == nil {
		return response.BadRequest(c, "media_type is required when media_url is provided.")
	}

	if req.MediaType != nil && *req.MediaType != "image" && *req.MediaType != "video" {
		return response.ValidationError(c, "media_type must be 'image' or 'video'.")
	}

	if req.Content != nil && len(*req.Content) > 500 {
		return response.ValidationError(c, "Caption must be 500 characters or fewer.")
	}

	post, err := h.postService.CreatePost(c.Context(), userID, req)
	if err != nil {
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusCreated, post)
}

func (h *PostHandler) GetFeed(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	cursor := c.Query("cursor", "")
	category := c.Query("category", "")

	limitStr := c.Query("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 20
	}
	if limit > 50 {
		return response.ValidationError(c, "limit must be 50 or fewer.")
	}

	feed, err := h.postService.GetFeed(c.Context(), userID, cursor, limit, category)
	if err != nil {
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, feed)
}

func (h *PostHandler) GetMyPosts(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return response.Unauthorized(c, "Invalid session.")
	}

	cursor := c.Query("cursor", "")
	limitStr := c.Query("limit", "30")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 30
	}
	if limit > 50 {
		limit = 50
	}

	feed, err := h.postService.GetMyPosts(c.Context(), userID, cursor, limit)
	if err != nil {
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, feed)
}

func (h *PostHandler) GetPostByID(c fiber.Ctx) error {
	postID := c.Params("postID")
	if postID == "" {
		return response.BadRequest(c, "Post ID is required.")
	}

	post, err := h.postService.GetByID(c.Context(), postID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return response.NotFound(c, "Post not found.")
		}
		return response.ServerError(c)
	}

	return response.Success(c, fiber.StatusOK, post)
}
