package service

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"staff-search-api/internal/config"
	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/pkg/storage"
	"staff-search-api/pkg/ulid"
)

var allowedContentTypes = map[string]bool{
	"image/jpeg":      true,
	"image/png":       true,
	"image/webp":      true,
	"video/mp4":       true,
	"video/quicktime": true,
}

var allowedFolders = map[string]bool{
	"avatars":      true,
	"portfolio":    true,
	"posts":        true,
	"stories":      true,
	"chat":         true,
	"intro_videos": true,
}

var extToContentType = map[string]string{
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".webp": "image/webp",
	".mp4":  "video/mp4",
	".mov":  "video/quicktime",
}

type MediaService struct {
	storage storage.StorageClient
	cfg     *config.Config
}

func NewMediaService(storageClient storage.StorageClient, cfg *config.Config) *MediaService {
	return &MediaService{storage: storageClient, cfg: cfg}
}

func (s *MediaService) GenerateUploadURL(ctx context.Context, userID string, req dto.UploadURLRequest) (dto.UploadURLResponse, error) {
	if !allowedContentTypes[req.ContentType] {
		return dto.UploadURLResponse{}, fmt.Errorf("%w: Unsupported content type.", model.ErrValidation)
	}

	if !allowedFolders[req.Folder] {
		return dto.UploadURLResponse{}, fmt.Errorf("%w: Invalid folder.", model.ErrValidation)
	}

	ext := filepath.Ext(req.FileName)
	if expectedType, ok := extToContentType[ext]; ok {
		if expectedType != req.ContentType {
			return dto.UploadURLResponse{}, fmt.Errorf("%w: File extension does not match content type.", model.ErrValidation)
		}
	} else {
		return dto.UploadURLResponse{}, fmt.Errorf("%w: Unsupported file extension.", model.ErrValidation)
	}

	key := fmt.Sprintf("%s/%s/%s%s", req.Folder, userID, ulid.New(), ext)

	expiry := time.Duration(s.cfg.StorageURLExpirySeconds) * time.Second
	presignedURL, err := s.storage.GeneratePresignedPutURL(ctx, key, req.ContentType, expiry)
	if err != nil {
		return dto.UploadURLResponse{}, fmt.Errorf("generate presigned url: %w", err)
	}

	publicURL := s.cfg.StoragePublicURL + "/" + key

	return dto.UploadURLResponse{
		UploadURL: presignedURL,
		FileKey:   key,
		PublicURL: publicURL,
		ExpiresIn: s.cfg.StorageURLExpirySeconds,
	}, nil
}

func (s *MediaService) DeleteFile(ctx context.Context, fileKey string) error {
	return s.storage.DeleteObject(ctx, fileKey)
}
