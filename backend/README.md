# Backend Setup

## 1) Docker-first run (recommended)
From project root:

docker compose up -d --build

Apply migration:

docker compose exec -T postgres psql -U postgres -d sales_db -f /dev/stdin < backend/migrations/001_init.sql

Health check:
- GET http://localhost:8080/health

Auth endpoints:
- POST http://localhost:8080/v1/auth/login
- POST http://localhost:8080/v1/auth/refresh
- GET http://localhost:8080/v1/auth/me (Bearer access token)

Default seeded admin from migrations/002_seed_admin.sql:
- email: admin@example.com
- password: admin123

## 2) Local Go run (optional)
If Go is installed locally:
- go mod tidy
- go run ./cmd/api
