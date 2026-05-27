package middleware

import (
    "context"
    "net/http"
    "strings"

    "github.com/your-org/gg-sheet-project/backend/internal/auth"
)

type contextKey string

const userContextKey contextKey = "authUser"

type AuthMiddleware struct {
    authService *auth.Service
}

func NewAuthMiddleware(authService *auth.Service) *AuthMiddleware {
    return &AuthMiddleware{authService: authService}
}

func (m *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

        if tokenString == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "missing bearer token", http.StatusUnauthorized)
            return
        }

        claims, err := m.authService.ValidateToken(tokenString, "access")
        if err != nil {
            http.Error(w, "invalid access token", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), userContextKey, claims)
        next(w, r.WithContext(ctx))
    }
}

func CurrentUser(r *http.Request) (*auth.UserClaims, bool) {
    claims, ok := r.Context().Value(userContextKey).(*auth.UserClaims)
    return claims, ok
}
