package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Atharva0506/trading_bot/internal/config"
	grpcdelivery "github.com/Atharva0506/trading_bot/internal/delivery/grpc"
	delivery "github.com/Atharva0506/trading_bot/internal/delivery/http"
	"github.com/Atharva0506/trading_bot/internal/delivery/websocket"
	"github.com/Atharva0506/trading_bot/internal/di"
	"github.com/Atharva0506/trading_bot/pkg/database"
	"github.com/Atharva0506/trading_bot/pkg/logger"
	signalpb "github.com/Atharva0506/trading_bot/proto/signal/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// WebSocket Hub
	hub := websocket.NewHub()
	go hub.Run()

	// HTTP Server
	router := delivery.NewRouter(c.UserHandler, c.SignalHandler, hub, cfg)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		Handler:      router,
	}

	go func() {
		slog.Info("HTTP server starting", "port", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("HTTP server failed", "error", err)
			os.Exit(1)
		}
	}()

	// gRPC Server
	grpcPort := 9090
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		slog.Error("failed to listen for gRPC", "error", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	signalGRPCHandler := grpcdelivery.NewSignalGRPCHandler(c.SignalService)
	signalpb.RegisterSignalServiceServer(grpcServer, signalGRPCHandler)
	reflection.Register(grpcServer) // enables grpcurl/grpcui for testing

	go func() {
		slog.Info("gRPC server starting", "port", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("gRPC server failed", "error", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	slog.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	grpcServer.GracefulStop()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("HTTP server shutdown failed", "error", err)
	}

	slog.Info("servers exited properly")
}
