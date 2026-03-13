package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"staff-search-api/internal/config"
	"staff-search-api/internal/model"
)

type StorageClient interface {
	GeneratePresignedPutURL(ctx context.Context, key string, contentType string, expiry time.Duration) (string, error)
	DeleteObject(ctx context.Context, key string) error
}

type S3StorageClient struct {
	s3Client  *s3.Client
	presigner *s3.PresignClient
	bucket    string
}

func NewS3StorageClient(cfg *config.Config) (*S3StorageClient, error) {
	creds := credentials.NewStaticCredentialsProvider(
		cfg.StorageAccessKeyID,
		cfg.StorageSecretAccessKey,
		"",
	)

	awsCfg := aws.Config{
		Region:      cfg.StorageRegion,
		Credentials: creds,
	}

	if cfg.StorageEndpoint != "" {
		awsCfg.BaseEndpoint = aws.String(cfg.StorageEndpoint)
	}

	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return &S3StorageClient{
		s3Client:  s3Client,
		presigner: s3.NewPresignClient(s3Client),
		bucket:    cfg.StorageBucket,
	}, nil
}

func (c *S3StorageClient) GeneratePresignedPutURL(ctx context.Context, key string, contentType string, expiry time.Duration) (string, error) {
	req, err := c.presigner.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(c.bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}, s3.WithPresignExpires(expiry))
	if err != nil {
		return "", fmt.Errorf("presign put object: %w", err)
	}
	return req.URL, nil
}

func (c *S3StorageClient) DeleteObject(ctx context.Context, key string) error {
	_, err := c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		var noSuchKey *types.NoSuchKey
		if errors.As(err, &noSuchKey) {
			return model.ErrNotFound
		}
		return fmt.Errorf("delete object: %w", err)
	}
	return nil
}
