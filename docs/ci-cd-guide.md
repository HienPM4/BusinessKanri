# CI/CD Guide (GitHub Actions)

This repo includes two workflows:
- CI: .github/workflows/ci.yml
- CD: .github/workflows/cd.yml

## 1) What CI does
On push/PR to main:
- Build and test backend Go.
- Start PostgreSQL service and apply migration check.
- Build frontend if frontend/package.json exists.

## 2) What CD does
On push to main (or manual run):
- Optional DB migrations if DATABASE_URL secret exists.
- Optional frontend deployment via Vercel deploy hook.
- Optional backend deployment via Render deploy hook.

## 3) Configure GitHub Secrets
Go to GitHub repo:
Settings -> Secrets and variables -> Actions -> New repository secret.

Create these secrets:
- DATABASE_URL
  Example: postgres://user:password@host:5432/dbname?sslmode=require
- VERCEL_DEPLOY_HOOK_URL
  From Vercel project -> Settings -> Git -> Deploy Hooks
- RENDER_DEPLOY_HOOK_URL
  From Render service -> Settings -> Deploy Hook

If a secret is missing, the related CD job is skipped automatically.

## 4) Trigger workflows
- CI runs automatically on push/PR to main.
- CD runs on push to main.
- You can also run CD manually from:
  Actions -> CD -> Run workflow.

## 5) Recommended branch policy
Use branch protection for main:
- Require pull request before merge.
- Require CI workflow to pass.

## 6) Notes
- Frontend CI is conditional and runs only when frontend/package.json exists.
- Migration check in CI applies all files in backend/migrations/*.sql.
- Keep db/schema.sql and migration files aligned.
