package http

import (
	"net/http"

	"github.com/Atharva0506/trading_bot/internal/config"
	"github.com/Atharva0506/trading_bot/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func NewRouter(userHandler *UserHandler, cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger)

	r.Get("/health", healthHandler)

	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
	})

	r.Route("/api/v1/protected", func(r chi.Router) {
		r.Use(middleware.JWTAuth(cfg.JWT.Secret))
		r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
			userID := r.Context().Value(middleware.UserIDKey)
			w.Write([]byte("Hello user " + userID.(uuid.UUID).String()))
		})
	})

	return r
}
