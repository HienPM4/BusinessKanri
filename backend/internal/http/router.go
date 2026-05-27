package http

import (
    "database/sql"
    "net/http"
    "strconv"
    "time"

    "github.com/your-org/gg-sheet-project/backend/internal/auth"
    "github.com/your-org/gg-sheet-project/backend/internal/config"
    "github.com/your-org/gg-sheet-project/backend/internal/http/handlers"
    "github.com/your-org/gg-sheet-project/backend/internal/http/middleware"
)

func NewRouter(db *sql.DB, cfg config.Config) http.Handler {
    mux := http.NewServeMux()

    accessTTLMinutes, _ := strconv.Atoi(cfg.AccessTokenTTLMinutes)
    if accessTTLMinutes <= 0 {
        accessTTLMinutes = 30
    }

    refreshTTLHours, _ := strconv.Atoi(cfg.RefreshTokenTTLHours)
    if refreshTTLHours <= 0 {
        refreshTTLHours = 168
    }

    authService := auth.NewService(
        cfg.JWTSecret,
        time.Duration(accessTTLMinutes)*time.Minute,
        time.Duration(refreshTTLHours)*time.Hour,
    )
    authHandler := handlers.NewAuthHandler(db, authService)
    authMiddleware := middleware.NewAuthMiddleware(authService)

    mux.HandleFunc("GET /health", handlers.Health)
    mux.HandleFunc("POST /v1/auth/login", authHandler.Login)
    mux.HandleFunc("POST /v1/auth/refresh", authHandler.Refresh)
    mux.HandleFunc("GET /v1/auth/me", authMiddleware.RequireAuth(authHandler.Me))

    return mux
}
