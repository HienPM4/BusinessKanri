package handlers

import (
    "context"
    "database/sql"
    "encoding/json"
    "errors"
    "net/http"

    "golang.org/x/crypto/bcrypt"

    "github.com/your-org/gg-sheet-project/backend/internal/auth"
    "github.com/your-org/gg-sheet-project/backend/internal/http/middleware"
)

type AuthHandler struct {
    db          *sql.DB
    authService *auth.Service
}

type authUserRecord struct {
    ID           string
    Email        string
    PasswordHash string
    FullName     string
    Role         string
}

type loginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type refreshRequest struct {
    RefreshToken string `json:"refreshToken"`
}

type authUserResponse struct {
    ID       string `json:"id"`
    Email    string `json:"email"`
    FullName string `json:"fullName"`
    Role     string `json:"role"`
}

func NewAuthHandler(db *sql.DB, authService *auth.Service) *AuthHandler {
    return &AuthHandler{db: db, authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req loginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeJSONError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid request body")
        return
    }

    if req.Email == "" || req.Password == "" {
        writeJSONError(w, http.StatusBadRequest, "VALIDATION_ERROR", "email and password are required")
        return
    }

    user, err := h.findUserByEmail(r.Context(), req.Email)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            writeJSONError(w, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Invalid email or password")
            return
        }

        writeJSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to process login")
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        writeJSONError(w, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Invalid email or password")
        return
    }

    tokenPair, err := h.authService.GenerateTokenPair(auth.UserTokenInput{
        ID:       user.ID,
        Email:    user.Email,
        FullName: user.FullName,
        Role:     user.Role,
    })
    if err != nil {
        writeJSONError(w, http.StatusInternalServerError, "TOKEN_ERROR", "Failed to generate token")
        return
    }

    writeJSON(w, http.StatusOK, map[string]interface{}{
        "accessToken":  tokenPair.AccessToken,
        "refreshToken": tokenPair.RefreshToken,
        "user": authUserResponse{
            ID:       user.ID,
            Email:    user.Email,
            FullName: user.FullName,
            Role:     user.Role,
        },
    })
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
    var req refreshRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeJSONError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid request body")
        return
    }

    if req.RefreshToken == "" {
        writeJSONError(w, http.StatusBadRequest, "VALIDATION_ERROR", "refreshToken is required")
        return
    }

    claims, err := h.authService.ValidateToken(req.RefreshToken, "refresh")
    if err != nil {
        writeJSONError(w, http.StatusUnauthorized, "INVALID_TOKEN", "Invalid refresh token")
        return
    }

    user, err := h.findUserByID(r.Context(), claims.UserID)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            writeJSONError(w, http.StatusUnauthorized, "USER_NOT_FOUND", "User not found")
            return
        }

        writeJSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to refresh token")
        return
    }

    tokenPair, err := h.authService.GenerateTokenPair(auth.UserTokenInput{
        ID:       user.ID,
        Email:    user.Email,
        FullName: user.FullName,
        Role:     user.Role,
    })
    if err != nil {
        writeJSONError(w, http.StatusInternalServerError, "TOKEN_ERROR", "Failed to generate access token")
        return
    }

    writeJSON(w, http.StatusOK, map[string]string{
        "accessToken": tokenPair.AccessToken,
    })
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
    claims, ok := middleware.CurrentUser(r)
    if !ok {
        writeJSONError(w, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized")
        return
    }

    writeJSON(w, http.StatusOK, map[string]interface{}{
        "user": authUserResponse{
            ID:       claims.UserID,
            Email:    claims.Email,
            FullName: claims.FullName,
            Role:     claims.Role,
        },
    })
}

func (h *AuthHandler) findUserByEmail(ctx context.Context, email string) (authUserRecord, error) {
    query := `
SELECT id, email, password_hash, full_name, role
FROM users
WHERE email = $1 AND is_active = TRUE
LIMIT 1
`

    row := h.db.QueryRowContext(ctx, query, email)

    var user authUserRecord
    err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.FullName, &user.Role)
    return user, err
}

func (h *AuthHandler) findUserByID(ctx context.Context, id string) (authUserRecord, error) {
    query := `
SELECT id, email, password_hash, full_name, role
FROM users
WHERE id = $1 AND is_active = TRUE
LIMIT 1
`

    row := h.db.QueryRowContext(ctx, query, id)

    var user authUserRecord
    err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.FullName, &user.Role)
    return user, err
}

func writeJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    _ = json.NewEncoder(w).Encode(payload)
}

func writeJSONError(w http.ResponseWriter, statusCode int, code string, message string) {
    writeJSON(w, statusCode, map[string]interface{}{
        "error": map[string]interface{}{
            "code":    code,
            "message": message,
            "details": []string{},
        },
    })
}
