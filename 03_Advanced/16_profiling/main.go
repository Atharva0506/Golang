package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // Side-effect import: registers /debug/pprof/* HTTP handlers
	"time"
)

// --- CPU-intensive work to profile ---

// slowFibonacci deliberately uses recursion to waste CPU cycles.
// This is the kind of function pprof will highlight in a flame graph.
func slowFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return slowFibonacci(n-1) + slowFibonacci(n-2)
}

// generateData fills a large slice with random numbers, causing heap allocations.
func generateData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	return data
}

// workHandler runs some CPU work on each request so pprof can capture a profile.
func workHandler(w http.ResponseWriter, r *http.Request) {
	result := slowFibonacci(35)
	_ = generateData(100_000)
	fmt.Fprintf(w, "fib(35) = %d\n", result)
}

func main() {
	// net/http/pprof automatically registers these routes on the default mux when imported:
	//
	//  /debug/pprof/          — index page with links to all profiles
	//  /debug/pprof/cmdline   — current process command line
	//  /debug/pprof/profile   — 30-second CPU profile (download with `go tool pprof`)
	//  /debug/pprof/symbol    — symbol lookup
	//  /debug/pprof/trace     — execution trace
	//  /debug/pprof/heap      — a heap allocation snapshot
	//  /debug/pprof/goroutine — stack traces of all current goroutines

	http.HandleFunc("/work", workHandler)

	slog.Info("pprof server starting on :6060")
	slog.Info("visit http://localhost:6060/debug/pprof/ in your browser")
	slog.Info("")
	slog.Info("HOW TO CAPTURE A CPU PROFILE:")
	slog.Info("  1. Keep this server running and hit /work a few times:")
	slog.Info("     curl http://localhost:6060/work")
	slog.Info("  2. Capture a 10-second CPU profile:")
	slog.Info("     go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10")
	slog.Info("  3. Inside pprof, run: top10  or  web")
	slog.Info("")
	slog.Info("HOW TO CAPTURE A HEAP PROFILE:")
	slog.Info("  go tool pprof http://localhost:6060/debug/pprof/heap")

	// Simulate background work so there's always something to profile
	go func() {
		for {
			_ = slowFibonacci(30)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	if err := http.ListenAndServe(":6060", nil); err != nil {
		slog.Error("server failed", "error", err)
	}
}
