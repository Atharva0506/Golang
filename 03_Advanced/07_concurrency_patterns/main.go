package main

import (
	"log/slog"
	"sync"
	"time"
)

func emailWorker(workerId int, emailJobs <-chan string, wg *sync.WaitGroup) {
	// 1. Defer Done() at the top of the function!
	// This ensures that when the channel closes and the `for` loop finishes, the worker announces it is going home!
	defer wg.Done()

	for email := range emailJobs {
		// Used slog.Int instead of string(rune) so it actually prints the number!
		slog.Info("Worker Processing", slog.Int("WorkerID", workerId), slog.String("Email", email))

		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan string, 10)

	for i := 1; i <= 3; i++ {
		// 2. We Add(1) to the WaitGroup BEFORE launching the goroutine!
		// We are tracking the 3 Workers, not the individual emails!
		wg.Add(1)
		go emailWorker(i, jobs, &wg)
	}
	emails := []string{"bob@gmail", "alice@gmail", "john@gmail", "sarah@gmail", "mike@gmail", "tom@gmail"}
	for _, email := range emails {
		jobs <- email
	}
	close(jobs)
	wg.Wait()
}
