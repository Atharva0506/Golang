package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Atharva0506/trading_bot/pkg/apperrors"
	"github.com/Atharva0506/trading_bot/pkg/auth"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

func JWTAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if strings.TrimSpace(authHeader) == "" {
				apperrors.HandleHTTPError(w, apperrors.NewUnauthorized("missing authorization header"))
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				apperrors.HandleHTTPError(w, apperrors.NewUnauthorized("invalid authorization format"))
				return
			}

			claims, err := auth.ValidateToken(tokenString, secret)
			if err != nil {
				apperrors.HandleHTTPError(w, apperrors.NewUnauthorized("invalid or expired token"))
				return
			}
			if claims.TokenType != "access" {
				apperrors.HandleHTTPError(w, apperrors.NewUnauthorized("invalid token type"))
				return
			}
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, RoleKey, claims.Role)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
