package main

import (
	"context"
	"log/slog"
	"time"
)

func downloadFile(ctx context.Context, filename string) {
	slog.Info("Starting file Download: ", slog.String("file", filename))
	for i := 0; i < 5; i++ {

		select {
		case <-ctx.Done():
			slog.Info("Stoping Download")
			return
		default:
			slog.Info("Processing...")
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	ctx := context.Background()

	kill, cancel := context.WithCancel(ctx)

	go downloadFile(kill, "luffy.txt")

	time.Sleep(2 * time.Second)

	slog.Info("Stop Worker")
	cancel()

	time.Sleep(1 * time.Second)
}
