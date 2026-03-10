# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

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
│   ├── handler/                         # One file per domain: auth_handler.go, etc.
│   ├── service/                         # One file per domain: auth_service.go, etc.
│   ├── repository/                      # One file per domain: user_repository.go, etc.
│   ├── model/                           # DB row structs + domain types
│   ├── middleware/                       # JWT, rate limiter, CORS
│   └── dto/                             # Request/Response DTOs
├── pkg/
│   ├── database/postgres.go             # GORM connection setup
│   ├── cache/redis.go                   # Redis connection
│   ├── jwt/jwt.go                       # Token generation/validation
│   ├── ulid/ulid.go                     # ULID generator
│   └── response/response.go            # Standard error envelope helpers
├── migrations/                          # golang-migrate SQL (up + down)
└── seeds/demo_data.sql                  # Dev/staging seed accounts
```

## Error Response Envelope

Every non-2xx response uses this shape:

```json
{ "error": "<machine_code>", "message": "<human_readable>" }
```

Error codes: `bad_request`, `invalid_token`, `unauthorized`, `forbidden`, `account_disabled`, `not_found`, `conflict`, `validation_error`, `server_error`, `rate_limited`.

Use helpers from `pkg/response/` — never construct error JSON manually in handlers.

## JWT Middleware

Applied to all `/api/v1/*` routes EXCEPT:
- `POST /api/v1/auth/login`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/google`
- `POST /api/v1/auth/apple`
- `POST /api/v1/auth/refresh`

Sets `c.Locals("userID")` and `c.Locals("role")` on success.

## Primary Keys

All tables use ULID (`VARCHAR(26)`), generated in Go via `pkg/ulid.New()`. Never use PostgreSQL sequences or UUID functions.

## Environment Variables

See `.env.example` for all required variables. All env access goes through `internal/config/config.go`.

## Project Context

This is the backend API for **staffsearch** — a platform connecting customers with service staff (beauticians, nail artists, massage therapists). The Flutter frontend is at `../staff-search-app/`.

Core domain concepts:
- **Users**: role-based (`user` | `staff` | `admin`)
- **Staff**: searchable by category, location, rating
- **Bookings**: `pending` → `confirmed` → `completed` | `cancelled`
- **Services**: staff-defined service menus with price/duration
- **Live streams**: real-time sessions via Agora SDK
- **Tips/Gifts**: coin-based payments via Stripe
