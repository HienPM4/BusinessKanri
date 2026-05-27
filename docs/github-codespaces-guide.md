# GitHub Codespaces Guide (No Docker Desktop Required)

Use this guide if you cannot install Go/Node/Docker locally.

## 1) Push project to GitHub
- Create a GitHub repository.
- Push current project code to the repository.

## 2) Create a Codespace
- Open your repository on GitHub.
- Click Code -> Codespaces -> Create codespace on main.
- Wait until VS Code web opens.

This repo already includes:
- .devcontainer/devcontainer.json
- .devcontainer/docker-compose.yml

So Codespaces will automatically start:
- app development container
- PostgreSQL service container

## 3) Start backend in Codespaces
From project root terminal:

cd backend
go mod tidy
go run ./cmd/api

Backend health check:
- http://localhost:8080/health

## 4) Apply migration (inside Codespaces)
From project root terminal:

for file in backend/migrations/*.sql; do psql "postgres://postgres:postgres@postgres:5432/sales_db?sslmode=disable" -f "$file"; done

Default admin account after migration:
- email: admin@example.com
- password: admin123

## 5) Initialize frontend in Codespaces
From project root terminal:

cd frontend
npx create-next-app@latest . --typescript --tailwind --eslint --app --src-dir --import-alias "@/*"
npm install @tanstack/react-query @tanstack/react-table react-hook-form zod @hookform/resolvers date-fns clsx lucide-react

Run frontend:

npm run dev -- --host 0.0.0.0 --port 3000

Open forwarded ports:
- 3000: frontend
- 8080: backend

## 6) Daily workflow
- Start/stop Codespace from GitHub UI.
- No local runtime installation required.
- Commit and push code from Codespaces directly.

## 7) Optional: using local VS Code
- Install GitHub Codespaces extension.
- Connect to your codespace from desktop VS Code.
