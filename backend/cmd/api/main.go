package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/your-org/gg-sheet-project/backend/internal/config"
    appdb "github.com/your-org/gg-sheet-project/backend/internal/db"
    apphttp "github.com/your-org/gg-sheet-project/backend/internal/http"
)

func main() {
    cfg := config.Load()

    db, err := appdb.Open(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("open db: %v", err)
    }
    defer db.Close()

    router := apphttp.NewRouter(db, cfg)

    srv := &http.Server{
        Addr:         ":" + cfg.Port,
        Handler:      router,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    go func() {
        log.Printf("api listening on :%s", cfg.Port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %v", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Printf("server shutdown error: %v", err)
    }
}
