package http

import (
	"net/http"

	"github.com/Atharva0506/trading_bot/internal/config"
	"github.com/Atharva0506/trading_bot/internal/delivery/websocket"
	"github.com/Atharva0506/trading_bot/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// NewRouter creates and configures the application router with all routes.
func NewRouter(userHandler *UserHandler, signalHandler *SignalHandler, hub *websocket.Hub, cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	limiter := middleware.NewIPRateLimiter(10, 20)
	r.Use(middleware.RequestLogger)
	r.Use(middleware.RateLimiter(limiter))

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

	r.Route("/api/v1/signals", func(r chi.Router) {
		r.Use(middleware.JWTAuth(cfg.JWT.Secret))
		r.Post("/", signalHandler.CreateSignal)
		r.Get("/", signalHandler.GetAllSignals)
		r.Get("/{symbol}", signalHandler.GetSignalsBySymbol)
	})

	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	return r
}
