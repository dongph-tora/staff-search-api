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
	"staff-search-api/pkg/jwt"
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
	if err := db.AutoMigrate(&model.User{}, &model.RefreshToken{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	// JWT
	jwtService := jwt.NewService(cfg.JWTSecret, cfg.JWTAccessExpiry, cfg.JWTRefreshExpiry)

	// Repositories
	userRepo := repository.NewUserRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)

	// Services
	authService := service.NewAuthService(userRepo, refreshTokenRepo, jwtService)
	userService := service.NewUserService(userRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService, userService)
	healthHandler := handler.NewHealthHandler()

	// Fiber app
	app := fiber.New(fiber.Config{
		AppName: "StaffSearch API v1",
	})

	// Routes
	router.Setup(app, jwtService, authHandler, healthHandler)

	// Start
	log.Printf("Starting server on :%s (env=%s)", cfg.AppPort, cfg.AppEnv)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
