package http

import (
	"github.com/Atharva0506/trading_bot/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger)

	r.Get("/health", healthHandler)
	return r
}
