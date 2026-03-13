package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"staff-search-api/internal/config"
	"staff-search-api/internal/handler"
	"staff-search-api/internal/model"
	"staff-search-api/internal/repository"
	"staff-search-api/internal/service"
	"staff-search-api/pkg/database"
	"staff-search-api/pkg/email"
	"staff-search-api/pkg/jwt"
	"staff-search-api/pkg/storage"
	"staff-search-api/router"
)

func main() {
	// Load .env in development
	if os.Getenv("APP_ENV") == "" {
		_ = godotenv.Load()
	}

	cfg := config.Load()

	// Database (GORM)
	db, err := database.NewPostgres(cfg.DatabaseURL, cfg.AppEnv)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate
	if err := db.AutoMigrate(
		&model.User{},
		&model.RefreshToken{},
		&model.PasswordResetToken{},
		&model.StaffProfile{},
		&model.StaffPortfolioPhoto{},
		&model.StaffSocialLink{},
		&model.Post{},
		&model.Like{},
		&model.Comment{},
		&model.Follow{},
		&model.Story{},
		&model.Service{},
		&model.Booking{},
		&model.Tip{},
		&model.PointTransaction{},
		&model.LiveStream{},
		&model.Review{},
		&model.Notification{},
		&model.HeadhuntOffer{},
	); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	// JWT
	jwtService := jwt.NewService(cfg.JWTSecret, cfg.JWTAccessExpiry, cfg.JWTRefreshExpiry)

	// Repositories
	userRepo := repository.NewUserRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	passwordResetRepo := repository.NewPasswordResetRepository(db)

	// Staff
	staffRepo := repository.NewStaffRepository(db)
	staffNumberService := service.NewStaffNumberService(staffRepo)

	// Posts
	postRepo := repository.NewPostRepository(db)

	staffService := service.NewStaffService(staffRepo, staffNumberService, userRepo, postRepo, db)
	postService := service.NewPostService(postRepo)

	// Email service
	var emailSvc email.EmailSender
	if cfg.SMTPHost != "" {
		emailSvc = email.NewSMTPEmailSender(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPFrom)
	} else {
		emailSvc = &email.NoOpEmailSender{}
	}

	// Media
	const uploadsDir = "./uploads"
	var storageClient storage.StorageClient
	var localUploadSvc *service.LocalUploadService

	if cfg.StorageProvider == "local" {
		localClient := storage.NewLocalStorageClient(uploadsDir)
		storageClient = localClient
		localUploadSvc = service.NewLocalUploadService(uploadsDir, cfg.AppBaseURL)
		log.Println("⚠️  Storage: local filesystem (dev only)")
	} else {
		s3Client, err := storage.NewS3StorageClient(cfg)
		if err != nil {
			log.Fatalf("Failed to initialize storage client: %v", err)
		}
		storageClient = s3Client
	}
	mediaService := service.NewMediaService(storageClient, cfg)
	portfolioService := service.NewStaffPortfolioService(staffRepo, storageClient, cfg)

	// Services
	authService := service.NewAuthService(userRepo, refreshTokenRepo, jwtService, passwordResetRepo, emailSvc, cfg.AppBaseURL, cfg.GoogleClientID)
	userService := service.NewUserService(userRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService, userService)
	healthHandler := handler.NewHealthHandler()
	userHandler := handler.NewUserHandler(userService)
	staffHandler := handler.NewStaffHandler(staffService, portfolioService)
	mediaHandler := handler.NewMediaHandler(mediaService, localUploadSvc)
	publicURL := cfg.StoragePublicURL
	if cfg.StorageProvider == "local" {
		publicURL = cfg.AppBaseURL
	}
	postHandler := handler.NewPostHandler(postService, publicURL)

	// Fiber app
	app := fiber.New(fiber.Config{
		AppName: "StaffSearch API v1",
	})

	// Routes
	router.Setup(app, jwtService, authHandler, healthHandler, userHandler, staffHandler, mediaHandler, postHandler)

	// Start
	log.Printf("Starting server on :%s (env=%s)", cfg.AppPort, cfg.AppEnv)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
