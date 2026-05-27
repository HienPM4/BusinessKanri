package config

import "os"

type Config struct {
    Port                  string
    DatabaseURL           string
    JWTSecret             string
    AccessTokenTTLMinutes string
    RefreshTokenTTLHours  string
}

func Load() Config {
    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }

    return Config{
        Port:                  port,
        DatabaseURL:           os.Getenv("DATABASE_URL"),
        JWTSecret:             getEnv("JWT_SECRET", "dev-secret-change-me"),
        AccessTokenTTLMinutes: getEnv("ACCESS_TOKEN_TTL_MINUTES", "30"),
        RefreshTokenTTLHours:  getEnv("REFRESH_TOKEN_TTL_HOURS", "168"),
    }
}

func getEnv(key string, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }

    return value
}
