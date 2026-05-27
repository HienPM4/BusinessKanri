# Docker Start Guide

This guide helps you start the whole project with Docker, even when local Go/Node are unavailable.

## 1) Prerequisite
- Docker Desktop installed and running.
- Open terminal at project root.

## 2) First time setup
Run backend + database and apply migration:

powershell -ExecutionPolicy Bypass -File ./scripts/docker-start.ps1

Verify:
- Backend health: http://localhost:8080/health

## 3) Start full stack (backend + db + frontend)

powershell -ExecutionPolicy Bypass -File ./scripts/docker-start.ps1 -WithFrontend

This command will:
- Start PostgreSQL and backend.
- Apply migration 001.
- Auto-initialize frontend if frontend/package.json is missing.
- Start frontend on port 3000.

Verify:
- Frontend: http://localhost:3000
- Backend health: http://localhost:8080/health

## 4) Useful commands
View containers:

docker compose ps

View logs:

docker compose logs -f backend
docker compose --profile frontend logs -f frontend
docker compose logs -f postgres

Stop all:

docker compose --profile frontend down

Stop and remove volume data:

docker compose --profile frontend down -v

## 5) Re-run migration manually

Get-Content "backend/migrations/001_init.sql" | docker compose exec -T postgres psql -U postgres -d sales_db -f -

## 6) Troubleshooting
- If frontend service says package.json missing, run:

powershell -ExecutionPolicy Bypass -File ./scripts/docker-init-frontend.ps1

- If port 5432/8080/3000 is busy, change port mapping in docker-compose.yml.
