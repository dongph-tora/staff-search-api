# staff-search-api

Go + Fiber v3 backend for **staffsearch** — REST API serving the Flutter app.

## Stack

- **Go 1.26** · Fiber v3
- **PostgreSQL 17** · GORM (AutoMigrate in dev)
- **Redis 7** · sessions, pub/sub
- **JWT** (access 1h / refresh 30d)
- **Storage** — Cloudflare R2 or local `./uploads/`
- **Email** — SMTP or no-op (dev)

## Requirements

- Go 1.21+
- Docker + Docker Compose (for PostgreSQL & Redis)

## Quick start

```bash
# 1. Clone & enter directory
cd staff-search-api

# 2. Start PostgreSQL + Redis
docker compose up -d

# 3. Copy env and fill in values
cp .env.example .env

# 4. Download dependencies
go mod download

# 5. Run the server (AutoMigrate runs on startup)
go run .
```

Server starts on `http://localhost:8080`.

## Environment variables

| Variable | Default | Description |
|---|---|---|
| `APP_PORT` | `8080` | HTTP port |
| `APP_BASE_URL` | `http://localhost:8080` | Used in password reset links |
| `DATABASE_URL` | see `.env.example` | PostgreSQL connection string |
| `REDIS_URL` | `redis://localhost:6380/0` | Redis connection string |
| `JWT_SECRET` | — | **Required.** Change in production |
| `GOOGLE_CLIENT_ID` | — | Google OAuth (optional) |
| `SMTP_HOST` | — | Leave empty to use no-op mailer |
| `STORAGE_PROVIDER` | `r2` | `r2` / `local` |

See `.env.example` for the full list.

## Seed demo data

```bash
psql $DATABASE_URL -f seeds/demo_data.sql
```

Demo accounts (password: `demo123`):

| Role | Email |
|---|---|
| User | `user@demo.com` |
| Staff | `staff@demo.com` |

## Auth endpoints

| Method | Path | Description |
|---|---|---|
| POST | `/api/v1/auth/register` | Register |
| POST | `/api/v1/auth/login` | Login |
| POST | `/api/v1/auth/refresh` | Refresh access token |
| POST | `/api/v1/auth/logout` | Logout |
| GET | `/api/v1/auth/me` | Current user |
| POST | `/api/v1/auth/google` | Google sign-in |
| POST | `/api/v1/auth/password-reset/request` | Request password reset email |
| POST | `/api/v1/auth/password-reset/confirm` | Confirm reset with token |

## Build Docker image

```bash
docker build -t staff-search-api .
docker run --env-file .env -p 8080:8080 staff-search-api
```

## Project structure

```
staff-search-api/
├── internal/
│   ├── config/       # Env config
│   ├── dto/          # Request / response structs
│   ├── handler/      # HTTP handlers (Fiber)
│   ├── model/        # GORM models
│   ├── repository/   # DB queries
│   └── service/      # Business logic
├── migrations/       # SQL migration files
├── pkg/              # Shared packages (jwt, email, storage, ulid)
├── router/           # Route registration
├── seeds/            # Demo seed SQL
├── main.go
└── docker-compose.yml
```
