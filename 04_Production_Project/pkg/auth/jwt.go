package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role"`
	TokenType string    `json:"token_type"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uuid.UUID, role, secret, tokenType string, expiration time.Duration) (string, error) {
	claims := &Claims{
		UserID:    userID,
		Role:      role,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := payload.SignedString([]byte(secret))
	if err != nil {
		return "", err // TDOD ADD stranderd error throuout application
	}
	return token, nil
}

func ValidateToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims") // or use your apperrors package
	}
	return claims, nil
}
