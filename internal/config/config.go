package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppEnv     string
	AppPort    string
	AppBaseURL string

	DatabaseURL string
	RedisURL    string

	JWTSecret       string
	JWTAccessExpiry int
	JWTRefreshExpiry int

	GoogleClientID string

	SMTPHost string
	SMTPPort int
	SMTPUser string
	SMTPPass string
	SMTPFrom string

	StorageProvider        string
	StorageBucket          string
	StorageRegion          string
	StorageAccessKeyID     string
	StorageSecretAccessKey string
	StorageEndpoint        string
	StoragePublicURL       string
	StorageURLExpirySeconds int
}

func Load() *Config {
	return &Config{
		AppEnv:     getEnv("APP_ENV", "development"),
		AppPort:    getEnv("APP_PORT", "3000"),
		AppBaseURL: getEnv("APP_BASE_URL", "http://localhost:3000"),

		DatabaseURL: getEnv("DATABASE_URL", "postgres://staffsearch:staffsearch@localhost:5432/staffsearch?sslmode=disable"),
		RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379/0"),

		JWTSecret:        getEnv("JWT_SECRET", "dev-secret-change-me"),
		JWTAccessExpiry:  getEnvInt("JWT_ACCESS_EXPIRY", 3600),
		JWTRefreshExpiry: getEnvInt("JWT_REFRESH_EXPIRY", 2592000),

		GoogleClientID: getEnv("GOOGLE_CLIENT_ID", ""),

		SMTPHost: getEnv("SMTP_HOST", ""),
		SMTPPort: getEnvInt("SMTP_PORT", 587),
		SMTPUser: getEnv("SMTP_USER", ""),
		SMTPPass: getEnv("SMTP_PASS", ""),
		SMTPFrom: getEnv("SMTP_FROM", ""),

		StorageProvider:         getEnv("STORAGE_PROVIDER", "r2"),
		StorageBucket:           getEnv("STORAGE_BUCKET", ""),
		StorageRegion:           getEnv("STORAGE_REGION", "auto"),
		StorageAccessKeyID:      getEnv("STORAGE_ACCESS_KEY_ID", ""),
		StorageSecretAccessKey:  getEnv("STORAGE_SECRET_ACCESS_KEY", ""),
		StorageEndpoint:         getEnv("STORAGE_ENDPOINT", ""),
		StoragePublicURL:        getEnv("STORAGE_PUBLIC_URL", ""),
		StorageURLExpirySeconds: getEnvInt("STORAGE_URL_EXPIRY_SECONDS", 900),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		if n, err := strconv.Atoi(val); err == nil {
			return n
		}
	}
	return fallback
}
