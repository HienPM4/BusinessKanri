package auth

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type Service struct {
    jwtSecret  []byte
    accessTTL  time.Duration
    refreshTTL time.Duration
}

type UserClaims struct {
    UserID    string `json:"uid"`
    Email     string `json:"email"`
    FullName  string `json:"fullName"`
    Role      string `json:"role"`
    TokenType string `json:"tokenType"`
    jwt.RegisteredClaims
}

type TokenPair struct {
    AccessToken  string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
}

type UserTokenInput struct {
    ID       string
    Email    string
    FullName string
    Role     string
}

func NewService(secret string, accessTTL time.Duration, refreshTTL time.Duration) *Service {
    return &Service{
        jwtSecret:  []byte(secret),
        accessTTL:  accessTTL,
        refreshTTL: refreshTTL,
    }
}

func (s *Service) GenerateTokenPair(user UserTokenInput) (TokenPair, error) {
    accessToken, err := s.generateToken(user, "access", s.accessTTL)
    if err != nil {
        return TokenPair{}, err
    }

    refreshToken, err := s.generateToken(user, "refresh", s.refreshTTL)
    if err != nil {
        return TokenPair{}, err
    }

    return TokenPair{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }, nil
}

func (s *Service) ValidateToken(tokenString string, expectedTokenType string) (*UserClaims, error) {
    claims := &UserClaims{}

    parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }

        return s.jwtSecret, nil
    })
    if err != nil {
        return nil, err
    }

    if !parsedToken.Valid {
        return nil, errors.New("invalid token")
    }

    if claims.TokenType != expectedTokenType {
        return nil, errors.New("invalid token type")
    }

    return claims, nil
}

func (s *Service) generateToken(user UserTokenInput, tokenType string, ttl time.Duration) (string, error) {
    now := time.Now().UTC()

    claims := UserClaims{
        UserID:    user.ID,
        Email:     user.Email,
        FullName:  user.FullName,
        Role:      user.Role,
        TokenType: tokenType,
        RegisteredClaims: jwt.RegisteredClaims{
            Subject:   user.ID,
            IssuedAt:  jwt.NewNumericDate(now),
            ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString(s.jwtSecret)
}
