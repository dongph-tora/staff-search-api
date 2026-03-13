package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"staff-search-api/internal/model"
)

// LocalStorageClient stores files on the local filesystem under ./uploads/.
// For development only — replace with S3/R2 in production.
type LocalStorageClient struct {
	baseDir string
}

func NewLocalStorageClient(baseDir string) *LocalStorageClient {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		panic(fmt.Sprintf("local storage: cannot create base dir %s: %v", baseDir, err))
	}
	return &LocalStorageClient{baseDir: baseDir}
}

// GeneratePresignedPutURL is a no-op for local storage.
// Direct upload is handled by the /api/v1/media/upload endpoint instead.
func (c *LocalStorageClient) GeneratePresignedPutURL(_ context.Context, _ string, _ string, _ time.Duration) (string, error) {
	return "", fmt.Errorf("local storage does not support presigned URLs — use POST /api/v1/media/upload")
}

func (c *LocalStorageClient) DeleteObject(_ context.Context, key string) error {
	path := c.baseDir + "/" + key
	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return model.ErrNotFound
		}
		return fmt.Errorf("local storage delete: %w", err)
	}
	return nil
}

func (c *LocalStorageClient) BaseDir() string {
	return c.baseDir
}
