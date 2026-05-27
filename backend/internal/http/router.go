package http

import (
    "net/http"

    "github.com/your-org/gg-sheet-project/backend/internal/http/handlers"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /health", handlers.Health)

    return mux
}
