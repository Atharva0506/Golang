package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Atharva0506/trading_bot/internal/config"
	delivery "github.com/Atharva0506/trading_bot/internal/delivery/http"
	"github.com/Atharva0506/trading_bot/internal/delivery/websocket"
	"github.com/Atharva0506/trading_bot/internal/di"
	"github.com/Atharva0506/trading_bot/pkg/database"
	"github.com/Atharva0506/trading_bot/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.NewLogger(&cfg.Logger)
	slog.SetDefault(log)

	db, err := database.NewConnection(&cfg.Database)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := database.RunMigrations(db, "migrations"); err != nil {
		slog.Error("failed to run database migrations", "error", err)
		os.Exit(1)
	}

	c := di.NewContainer(db, cfg)

	hub := websocket.NewHub()
	go hub.Run()

	router := delivery.NewRouter(c.UserHandler, c.SignalHandler, hub, cfg)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		Handler:      router,
	}

	go func() {
		slog.Info("server starting", "port", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	slog.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}

	slog.Info("server exited properly")

}
