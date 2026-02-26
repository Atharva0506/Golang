package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func slowHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Slow Handler started")
	time.Sleep(3 * time.Second)
	slog.Info("Finshed Slow request")
	fmt.Fprint(w, "Finshed Slow request")
}
func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(slowHandler),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Listen", slog.Any("Error", err))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Shutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Shutdown", slog.Any("Error", err))
	}
	slog.Info("Server exiting perfectly cleanly.")
}
