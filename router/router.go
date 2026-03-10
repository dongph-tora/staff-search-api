package router

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"staff-search-api/internal/handler"
	"staff-search-api/internal/middleware"
	"staff-search-api/pkg/jwt"
)

func Setup(
	app *fiber.App,
	jwtService *jwt.Service,
	authHandler *handler.AuthHandler,
	healthHandler *handler.HealthHandler,
) {
	// Global middleware
	app.Use(middleware.CORS())

	// Health check
	app.Get("/health", healthHandler.Check)

	// API v1
	v1 := app.Group("/api/v1")

	// --- Public auth routes (no JWT required) ---
	auth := v1.Group("/auth")
	auth.Use(middleware.RateLimiter(10, 1*time.Minute))

	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)
	auth.Post("/refresh", authHandler.Refresh)

	// --- Protected routes (JWT required) ---
	protected := v1.Group("", middleware.JWTMiddleware(jwtService))

	// Auth — protected
	protectedAuth := protected.Group("/auth")
	protectedAuth.Post("/logout", authHandler.Logout)
	protectedAuth.Get("/me", authHandler.Me)
	protectedAuth.Post("/privacy-policy/accept", authHandler.AcceptPrivacyPolicy)
}
