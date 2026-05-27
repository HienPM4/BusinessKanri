$ErrorActionPreference = "Stop"

Set-Location (Join-Path $PSScriptRoot "..")

if (Test-Path "frontend/package.json") {
    Write-Host "frontend/package.json already exists. Skip init."
    exit 0
}

Write-Host "Initializing Next.js app in ./frontend via Docker..."

docker run --rm -v "${PWD}/frontend:/app" -w /app node:20-alpine sh -c "npx create-next-app@latest . --typescript --tailwind --eslint --app --src-dir --import-alias '@/*' --yes"

docker run --rm -v "${PWD}/frontend:/app" -w /app node:20-alpine sh -c "npm install @tanstack/react-query @tanstack/react-table react-hook-form zod @hookform/resolvers date-fns clsx lucide-react"

Write-Host "Frontend initialization completed."
