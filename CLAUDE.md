# CLAUDE.md — staff-search-api

This file provides guidance to Claude Code (claude.ai/code) when working with the Go backend API.

## Stack

- **Go 1.26.1** — use `~/go/bin/go1.26.1` (system Go is 1.25.x, do not use it)
- **Fiber v3.1.0** — `github.com/gofiber/fiber/v3` (v3 API, not v2)
- **Fiber CLI 0.14.1** — `~/go/bin/fiber`
- **GORM** — ORM (gorm.io/gorm + gorm.io/driver/postgres)
- **PostgreSQL** — primary database via GORM
- **Redis** — sessions, cache, pub/sub (go-redis/v9)
- **JWT** — golang-jwt/v5, HS256
- **ULID** — primary keys, oklog/ulid/v2

## Common Commands

```bash
# Run dev server (hot reload)
cd staff-search-api && ~/go/bin/fiber dev

# Start infra (Postgres 5432, Redis 6380)
docker compose up -d

# Build
~/go/bin/go1.26.1 build ./...

# Run
~/go/bin/go1.26.1 run main.go

# Add dependency
~/go/bin/go1.26.1 get <package>@<version>

# Tidy modules
~/go/bin/go1.26.1 mod tidy

# Run tests
~/go/bin/go1.26.1 test ./...

# Run single test
~/go/bin/go1.26.1 test ./internal/handler/... -run TestFunctionName -v

# Database migrations
migrate -path migrations -database $DATABASE_URL up
migrate -path migrations -database $DATABASE_URL down

# Seed demo data
psql $DATABASE_URL -f seeds/demo_data.sql
```

## Fiber v3 API Differences from v2

Handler signature changed: `func(c fiber.Ctx) error` (not `*fiber.Ctx`).

```go
// v3 — correct
app.Get("/", func(c fiber.Ctx) error {
    return c.JSON(fiber.Map{"ok": true})
})
```

## Architecture — Clean Architecture (3-Layer)

```
HTTP Request
    ↓
[Router]  router/router.go — registers all routes + middleware
    ↓
[Handler] internal/handler/ — parse HTTP, delegate to service, write response
    ↓
[Service] internal/service/ — business logic, no HTTP knowledge
    ↓
[Repository] internal/repository/ — SQL queries, no business rules
    ↓
PostgreSQL / Redis
```

**Strict rules:**
- Handler NEVER queries DB directly
- Repository NEVER contains business logic
- Service NEVER imports fiber or HTTP packages

## Directory Structure

```
staff-search-api/
├── main.go                              # Entry point — wires config, DB, services, router
├── router/router.go                     # All route definitions + middleware groups
├── internal/
│   ├── config/config.go                 # Typed env config struct
│   ├── handler/                         # One file per domain
│   │   ├── auth_handler.go
│   │   ├── health_handler.go
│   │   ├── media_handler.go
│   │   ├── post_handler.go
│   │   ├── staff_handler.go
│   │   └── user_handler.go
│   ├── service/                         # One file per domain
│   │   ├── auth_service.go
│   │   ├── local_upload_service.go
│   │   ├── media_service.go
│   │   ├── post_service.go
│   │   ├── staff_number_service.go
│   │   ├── staff_portfolio_service.go
│   │   ├── staff_service.go
│   │   └── user_service.go
│   ├── repository/                      # One file per domain
│   │   ├── booking_repository.go
│   │   ├── follow_repository.go
│   │   ├── notification_repository.go
│   │   ├── password_reset_repository.go
│   │   ├── post_repository.go
│   │   ├── refresh_token_repository.go
│   │   ├── review_repository.go
│   │   ├── staff_repository.go
│   │   └── user_repository.go
│   ├── model/                           # DB row structs + domain types
│   │   ├── booking.go
│   │   ├── error.go
│   │   ├── headhunt.go
│   │   ├── live_stream.go
│   │   ├── notification.go
│   │   ├── payment.go
│   │   ├── post.go
│   │   ├── review.go
│   │   ├── staff.go
│   │   └── user.go
│   ├── middleware/                       # JWT, rate limiter, CORS
│   └── dto/                             # Request/Response DTOs
│       ├── auth_dto.go
│       ├── media_dto.go
│       ├── post_dto.go
│       ├── staff_dto.go
│       └── user_dto.go
├── pkg/
│   ├── database/postgres.go             # GORM connection setup
│   ├── cache/redis.go                   # Redis connection
│   ├── jwt/jwt.go                       # Token generation/validation
│   ├── ulid/ulid.go                     # ULID generator
│   ├── response/response.go             # Standard error envelope helpers
│   ├── email/                           # Email sending (password reset)
│   └── storage/                         # Object storage abstraction (S3/R2/local)
├── migrations/                          # golang-migrate SQL (up + down), 14 migrations
├── seeds/demo_data.sql                  # Dev/staging seed accounts
├── uploads/                             # Local file uploads (dev only)
├── docker-compose.yml                   # PostgreSQL + Redis
└── Dockerfile
```

## Live API Routes

### Public routes
| Method | Route | Handler |
|---|---|---|
| GET | `/health` | healthHandler.Check |
| POST | `/api/v1/auth/login` | authHandler.Login |
| POST | `/api/v1/auth/register` | authHandler.Register |
| POST | `/api/v1/auth/refresh` | authHandler.Refresh |
| POST | `/api/v1/auth/google` | authHandler.GoogleSignIn |
| POST | `/api/v1/auth/password-reset/request` | authHandler.RequestPasswordReset |
| POST | `/api/v1/auth/password-reset/confirm` | authHandler.ConfirmPasswordReset |
| GET | `/api/v1/staff/job-categories` | staffHandler.GetJobCategories |
| GET | `/api/v1/staff` | staffHandler.ListStaff |

### Protected routes (JWT required)
| Method | Route | Handler |
|---|---|---|
| POST | `/api/v1/auth/logout` | authHandler.Logout |
| GET | `/api/v1/auth/me` | authHandler.Me |
| POST | `/api/v1/auth/privacy-policy/accept` | authHandler.AcceptPrivacyPolicy |
| POST | `/api/v1/auth/change-password` | authHandler.ChangePassword |
| PATCH | `/api/v1/users/me` | userHandler.UpdateProfile |
| POST | `/api/v1/staff/profile` | staffHandler.CreateProfile |
| PATCH | `/api/v1/staff/profile` | staffHandler.UpdateProfile |
| GET | `/api/v1/staff/me` | staffHandler.GetMyProfile |
| GET | `/api/v1/staff/:userID` | staffHandler.GetProfile |
| POST | `/api/v1/staff/portfolio/photos` | staffHandler.AddPortfolioPhoto |
| DELETE | `/api/v1/staff/portfolio/photos/:photoID` | staffHandler.DeletePortfolioPhoto |
| PATCH | `/api/v1/staff/portfolio/photos/reorder` | staffHandler.ReorderPortfolioPhotos |
| POST | `/api/v1/posts` | postHandler.CreatePost |
| GET | `/api/v1/posts/feed` | postHandler.GetFeed |
| GET | `/api/v1/posts/mine` | postHandler.GetMyPosts |
| GET | `/api/v1/posts/:postID` | postHandler.GetPostByID |
| POST | `/api/v1/media/upload-url` | mediaHandler.GenerateUploadURL |
| POST | `/api/v1/media/upload` | mediaHandler.UploadFile |
| DELETE | `/api/v1/media` | mediaHandler.DeleteFile |

## Error Response Envelope

Every non-2xx response uses this shape:

```json
{ "error": "<machine_code>", "message": "<human_readable>" }
```

Error codes: `bad_request`, `invalid_token`, `unauthorized`, `forbidden`, `account_disabled`, `not_found`, `conflict`, `validation_error`, `server_error`, `rate_limited`.

Use helpers from `pkg/response/` — never construct error JSON manually in handlers.

## JWT Middleware

Applied to all `/api/v1/*` routes EXCEPT the public routes listed above. Sets `c.Locals("userID")` and `c.Locals("role")` on success.

## Primary Keys

All tables use ULID (`VARCHAR(26)`), generated in Go via `pkg/ulid.New()`. Never use PostgreSQL sequences or UUID functions.

## Database Migrations

14 migrations covering: users, refresh_tokens, password_reset_tokens, oauth_fields, staff_profiles, staff_portfolio_photos, posts, likes/comments/follows, bookings, tips/points, reviews, notifications, constraint renaming, lat/lng for staff.

Always create new migration files with the next sequence number. Never modify existing migrations.

## Environment Variables

See `.env.example` for all required variables. All env access goes through `internal/config/config.go`.

## Not Yet Implemented

Payments, live streaming, chat/messaging, real-time notifications (WebSocket/FCM), search (Elasticsearch), rankings, subscriptions, headhunting system, admin endpoints.

## Project Context

This is the backend API for **staffsearch** — a platform connecting customers with service staff (beauticians, nail artists, massage therapists). The Flutter frontend is at `../staff-search-app/`.
