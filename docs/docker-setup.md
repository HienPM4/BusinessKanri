# Docker Setup

This project can run without local Go or Node by using Docker.

## Prerequisite
- Install Docker Desktop and ensure Docker Engine is running.

## Run backend + PostgreSQL
From project root:

docker compose up -d --build

Check services:

docker compose ps

Health check:
- API: http://localhost:8080/health

## Apply migration
Migration SQL is available at backend/migrations/001_init.sql.
You can apply it inside PostgreSQL container:

docker compose exec -T postgres psql -U postgres -d sales_db -f /dev/stdin < backend/migrations/001_init.sql

## Stop services

docker compose down

## Remove data volume (optional)

docker compose down -v

## Frontend initialization without local Node
Initialize Next.js in frontend folder using Node container:

docker run --rm -it -v "${PWD}/frontend:/app" -w /app node:20-alpine sh -c "npx create-next-app@latest . --typescript --tailwind --eslint --app --src-dir --import-alias '@/*'"

Install required frontend libraries:

docker run --rm -it -v "${PWD}/frontend:/app" -w /app node:20-alpine sh -c "npm install @tanstack/react-query @tanstack/react-table react-hook-form zod @hookform/resolvers"
