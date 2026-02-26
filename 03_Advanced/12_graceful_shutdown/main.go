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
	// 1. We MUST build the actual `http.Server` struct so we have a pointer to it.
	// You cannot use the simple `http.ListenAndServe(":8080", nil)` for graceful shutdowns!
	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(slowHandler),
	}

	// 2. Launch the server in a GOROUTINE!
	// ListenAndServe blocks the thread forever. We need the main thread to stay awake and listen for Kill Signals!
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Listen", slog.Any("Error", err))
		}
	}()
	// 3. Create a channel to intercept OS Signals (like Ctrl+C or Kubernetes SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// 4. BLOCK FOREVER! The main thread freezes right here until the user hits Ctrl+C!
	<-quit
	slog.Info("Shutting down gracefully...")
	// 5. We give the server a 5-second deadline to finish active transactions (like our 3-second slowHandler)!
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 6. Gracefully drain the server. It stops accepting NEW traffic, but lets CURRENT traffic finish building their response.
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Shutdown", slog.Any("Error", err))
	}
	slog.Info("Server exiting perfectly cleanly.")
}
