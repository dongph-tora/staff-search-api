package router

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"staff-search-api/internal/handler"
	"staff-search-api/internal/middleware"
	"staff-search-api/pkg/jwt"
)

func Setup(
	app *fiber.App,
	jwtService *jwt.Service,
	authHandler *handler.AuthHandler,
	healthHandler *handler.HealthHandler,
	userHandler *handler.UserHandler,
	staffHandler *handler.StaffHandler,
	mediaHandler *handler.MediaHandler,
	postHandler *handler.PostHandler,
) {
	app.Use(middleware.CORS())
	app.Get("/health", healthHandler.Check)
	app.Use("/uploads", static.New("./uploads"))

	v1 := app.Group("/api/v1")

	// Public auth routes
	auth := v1.Group("/auth")
	auth.Use(middleware.RateLimiter(10, 1*time.Minute))
	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)
	auth.Post("/refresh", authHandler.Refresh)
	auth.Post("/google", authHandler.GoogleSignIn)
	auth.Post("/password-reset/request", authHandler.RequestPasswordReset)
	auth.Post("/password-reset/confirm", authHandler.ConfirmPasswordReset)

	// Public staff routes
	v1.Get("/staff/job-categories", staffHandler.GetJobCategories)
	v1.Get("/staff", staffHandler.ListStaff)

	// Protected routes
	protected := v1.Group("", middleware.JWTMiddleware(jwtService))

	// Auth — protected
	protectedAuth := protected.Group("/auth")
	protectedAuth.Post("/logout", authHandler.Logout)
	protectedAuth.Get("/me", authHandler.Me)
	protectedAuth.Post("/privacy-policy/accept", authHandler.AcceptPrivacyPolicy)
	protectedAuth.Post("/change-password", authHandler.ChangePassword)

	// Users
	users := protected.Group("/users")
	users.Patch("/me", userHandler.UpdateProfile)

	// Staff
	staff := protected.Group("/staff")
	staff.Post("/profile", staffHandler.CreateProfile)
	staff.Patch("/profile", staffHandler.UpdateProfile)
	staff.Get("/me", staffHandler.GetMyProfile)
	staff.Get("/:userID", staffHandler.GetProfile)

	// Portfolio
	portfolio := staff.Group("/portfolio")
	portfolio.Post("/photos", staffHandler.AddPortfolioPhoto)
	portfolio.Delete("/photos/:photoID", staffHandler.DeletePortfolioPhoto)
	portfolio.Patch("/photos/reorder", staffHandler.ReorderPortfolioPhotos)

	// Posts
	posts := protected.Group("/posts")
	posts.Post("", postHandler.CreatePost)
	posts.Get("/feed", postHandler.GetFeed)
	posts.Get("/mine", postHandler.GetMyPosts)
	posts.Get("/:postID", postHandler.GetPostByID)

	// Media
	media := protected.Group("/media")
	media.Post("/upload-url", mediaHandler.GenerateUploadURL)
	media.Post("/upload", mediaHandler.UploadFile)
	media.Delete("", mediaHandler.DeleteFile)
}
