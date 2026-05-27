param(
    [switch]$WithFrontend
)

$ErrorActionPreference = "Stop"
Set-Location (Join-Path $PSScriptRoot "..")

Write-Host "Starting PostgreSQL + backend..."
docker compose up -d --build postgres backend

Write-Host "Applying migration 001..."
Get-Content "backend/migrations/001_init.sql" | docker compose exec -T postgres psql -U postgres -d sales_db -f -

if ($WithFrontend) {
    if (-not (Test-Path "frontend/package.json")) {
        Write-Host "frontend/package.json not found. Initializing frontend first..."
        & "$PSScriptRoot/docker-init-frontend.ps1"
    }

    Write-Host "Starting frontend profile..."
    docker compose --profile frontend up -d frontend
}

Write-Host "Done."
Write-Host "Backend health: http://localhost:8080/health"
if ($WithFrontend) {
    Write-Host "Frontend app: http://localhost:3000"
}
