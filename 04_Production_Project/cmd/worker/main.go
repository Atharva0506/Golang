package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Atharva0506/trading_bot/internal/config"
	"github.com/Atharva0506/trading_bot/internal/worker"
	"github.com/Atharva0506/trading_bot/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.NewLogger(&cfg.Logger)
	slog.SetDefault(log)

	pool := worker.NewWorkerPool(3, 100)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool.Start(ctx)
	pool.Submit(&worker.LogJob{Message: "worker process started!"})
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	slog.Info("shutting down worker...")
	cancel()
	pool.Shutdown()
}
