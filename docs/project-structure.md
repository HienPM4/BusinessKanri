# Project Structure (v1)

## Root
- AGENTS.md
- db/schema.sql
- docs/api-contract.md
- docs/project-structure.md
- backend/
- frontend/

## Backend (Golang)
- cmd/api/main.go
- internal/config/config.go
- internal/db/postgres.go
- internal/http/router.go
- internal/http/handlers/health.go
- migrations/001_init.sql
- .env.example
- go.mod

Planned next backend modules:
- internal/auth
- internal/customers
- internal/products
- internal/orders
- internal/payments
- internal/reports

## Frontend (Next.js)
Recommended initialization command:
- npx create-next-app@latest frontend --typescript --tailwind --eslint --app --src-dir --import-alias "@/*"

Then install UI and data libraries:
- npm install @tanstack/react-query @tanstack/react-table react-hook-form zod @hookform/resolvers
- npx shadcn@latest init

Planned frontend pages:
- /login
- /dashboard
- /orders
- /orders/new
- /products
- /customers

## Conventions
- Keep API versioned under /v1.
- Keep SQL schema and migration in sync.
- Update AGENTS.md on every architecture or structure change.
