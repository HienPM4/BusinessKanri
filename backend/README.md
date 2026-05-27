# Backend Setup

## 1) Docker-first run (recommended)
From project root:

docker compose up -d --build

Apply migration:

docker compose exec -T postgres psql -U postgres -d sales_db -f /dev/stdin < backend/migrations/001_init.sql

Health check:
- GET http://localhost:8080/health

## 2) Local Go run (optional)
If Go is installed locally:
- go mod tidy
- go run ./cmd/api
