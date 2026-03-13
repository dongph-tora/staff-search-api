package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"staff-search-api/internal/dto"
	"staff-search-api/internal/model"
	"staff-search-api/pkg/ulid"
)

var localAllowedExts = map[string]string{
	".jpg":  "image",
	".jpeg": "image",
	".png":  "image",
	".webp": "image",
	".mp4":  "video",
	".mov":  "video",
}

var localAllowedFolders = map[string]bool{
	"avatars":      true,
	"portfolio":    true,
	"posts":        true,
	"stories":      true,
	"chat":         true,
	"intro_videos": true,
}

type LocalUploadService struct {
	baseDir   string
	publicURL string // e.g. http://localhost:8080
}

func NewLocalUploadService(baseDir, serverPublicURL string) *LocalUploadService {
	return &LocalUploadService{baseDir: baseDir, publicURL: serverPublicURL}
}

func (s *LocalUploadService) Save(_ context.Context, userID, folder string, fh *multipart.FileHeader) (dto.UploadURLResponse, error) {
	if !localAllowedFolders[folder] {
		return dto.UploadURLResponse{}, fmt.Errorf("%w: Invalid folder.", model.ErrValidation)
	}

	ext := strings.ToLower(filepath.Ext(fh.Filename))
	mediaType, ok := localAllowedExts[ext]
	if !ok {
		return dto.UploadURLResponse{}, fmt.Errorf("%w: Unsupported file type.", model.ErrValidation)
	}
	_ = mediaType

	key := fmt.Sprintf("%s/%s/%s%s", folder, userID, ulid.New(), ext)
	destPath := filepath.Join(s.baseDir, key)

	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return dto.UploadURLResponse{}, fmt.Errorf("create dirs: %w", err)
	}

	src, err := fh.Open()
	if err != nil {
		return dto.UploadURLResponse{}, fmt.Errorf("open upload: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(destPath)
	if err != nil {
		return dto.UploadURLResponse{}, fmt.Errorf("create file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return dto.UploadURLResponse{}, fmt.Errorf("write file: %w", err)
	}

	publicURL := fmt.Sprintf("%s/uploads/%s", s.publicURL, key)

	return dto.UploadURLResponse{
		UploadURL: "",
		FileKey:   key,
		PublicURL: publicURL,
		ExpiresIn: 0,
	}, nil
}
