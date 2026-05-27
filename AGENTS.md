# Sales Data Web System - Agent Working File

This file is the single source of truth for architecture and delivery decisions.
Any agent working in this repo must update this file when structure changes.

## 1) Product Goal
Build a web system to replace Google Sheet workflows for sales data entry, storage, tracking, and reporting.

Core outcomes:
- Fast sales order entry.
- Reliable data storage with auditability.
- Daily and monthly revenue reporting.
- Expandable structure for inventory and debt tracking.

## 2) Locked Technology Choices
- Frontend: Next.js, TypeScript, Tailwind CSS, shadcn/ui, TanStack Table.
- Backend: Golang REST API.
- Database: PostgreSQL.

Preferred additions:
- Frontend data + forms: TanStack Query, React Hook Form, Zod.
- Backend auth: JWT access + refresh tokens.
- Migrations: Goose or Atlas.
- Local development environment: Docker + Docker Compose.
- Remote development option: GitHub Codespaces.

## 3) High-Level Architecture
- Client (Next.js) talks to Go API over HTTPS.
- Go API owns business rules and database writes.
- PostgreSQL is the source of truth.

Logical modules:
- Auth and users.
- Customers.
- Products.
- Orders and order items.
- Payments.
- Reporting.

## 4) Data Model (Current v1)
Main tables:
- users
- customers
- products
- orders
- order_items
- payments
- inventory_transactions

See full SQL in db/schema.sql.

## 5) API Surface (Current v1)
Groups:
- /v1/auth
- /v1/customers
- /v1/products
- /v1/orders
- /v1/reports

See request/response contract in docs/api-contract.md.

## 6) Frontend App Sections (Current v1)
- Login page.
- Dashboard page.
- Orders list + filters + export.
- Create/edit order form with keyboard-first flow.
- Customers management page.
- Products management page.

## 7) Non-Functional Baseline
- Role-based access (admin, staff, viewer).
- Input validation on both FE and BE.
- UTC timestamps in database.
- Soft delete only where needed.
- Audit fields: created_at, updated_at, created_by where applicable.

## 8) Update Policy (Must Follow)
When structure changes, update this file in the same task:
1. If tables or fields change: update section 4 and db/schema.sql.
2. If endpoints or payloads change: update section 5 and docs/api-contract.md.
3. If FE routes or modules change: update section 6.
4. If stack decisions change: update section 2.
5. Add a short changelog line in section 9.

Definition of done for architecture changes:
- Code changed.
- AGENTS.md updated.
- Related docs/schema updated.

## 9) Changelog
- 2026-05-27: Initialized architecture baseline, DB model v1, and API contract v1.
- 2026-05-27: Added backend skeleton (health endpoint), migration 001, and project structure docs.
- 2026-05-27: Added Docker Compose setup for PostgreSQL + backend and Docker-based frontend initialization guide.
- 2026-05-27: Added Docker start automation scripts and a single Docker start guide.
- 2026-05-27: Added GitHub Codespaces setup (.devcontainer) and run guide for no-local-install workflow.
- 2026-05-27: Initialized Next.js frontend base in frontend with runnable dev/build/start scripts.

## 10) Next Implementation Step
- Implement Auth module (login, refresh, middleware).
- Implement Product module CRUD endpoints.
- Implement Order create/list/detail/confirm endpoints.
- Initialize Next.js app in frontend folder and build login + orders pages (via Docker if local Node is unavailable).
- Add frontend Docker dev profile and keep start workflow in scripts/docker-start.ps1.
