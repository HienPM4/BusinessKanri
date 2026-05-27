# GG Sheet Project

Sales data web system replacing Google Sheet workflow.

## Current Status
- Architecture baseline created.
- DB schema v1 created.
- API contract v1 created.
- Backend skeleton initialized.
- Auth module initialized (login, refresh, JWT middleware, /v1/auth/me).
- Frontend setup guide added.
- Docker Compose setup added for backend + PostgreSQL.

## Important Docs
- AGENTS.md
- db/schema.sql
- docs/api-contract.md
- docs/project-structure.md
- docs/docker-setup.md
- docs/docker-start-guide.md
- docs/github-codespaces-guide.md
- docs/ci-cd-guide.md

## Quick Start (GitHub Codespaces)
If local Docker Desktop/Go/Node cannot be installed, use Codespaces:
1. Push repo to GitHub.
2. Create Codespace from GitHub UI.
3. Follow docs/github-codespaces-guide.md.

## Quick Start (Docker)
From project root:

powershell -ExecutionPolicy Bypass -File ./scripts/docker-start.ps1

For full stack (includes frontend):

powershell -ExecutionPolicy Bypass -File ./scripts/docker-start.ps1 -WithFrontend

Health check:
- http://localhost:8080/health

## CI/CD (GitHub Actions)
- CI workflow: .github/workflows/ci.yml
- CD workflow: .github/workflows/cd.yml
- Setup guide: docs/ci-cd-guide.md
