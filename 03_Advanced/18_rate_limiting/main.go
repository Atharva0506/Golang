package main

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

// ============================================================
// PART 1: Token Bucket — implemented from scratch
// ============================================================
//
// How it works:
//   - A "bucket" holds up to `capacity` tokens.
//   - Tokens refill at `rate` tokens-per-second.
//   - Each request consumes one token.
//   - If the bucket is empty, the request is rejected (or waits).

type TokenBucket struct {
	mu       sync.Mutex
	tokens   float64
	capacity float64
	rate     float64    // tokens per second
	lastFill time.Time
}

// NewTokenBucket creates a bucket that starts full.
func NewTokenBucket(capacity, ratePerSecond float64) *TokenBucket {
	return &TokenBucket{
		tokens:   capacity,
		capacity: capacity,
		rate:     ratePerSecond,
		lastFill: time.Now(),
	}
}

// Allow returns true if the request is permitted (consumes one token).
// Returns false immediately if the bucket is empty (non-blocking / "drop" policy).
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastFill).Seconds()
	tb.lastFill = now

	// Refill the bucket proportionally to the time elapsed.
	tb.tokens += elapsed * tb.rate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	if tb.tokens < 1 {
		return false
	}

	tb.tokens--
	return true
}

// ============================================================
// PART 2: Fixed Window Counter — implemented from scratch
// ============================================================
//
// How it works:
//   - Divide time into fixed windows (e.g., 1-second buckets).
//   - Allow up to `limit` requests per window.
//   - Simple, but can allow 2× the limit across a window boundary.

type FixedWindowLimiter struct {
	mu          sync.Mutex
	limit       int
	count       int
	windowStart time.Time
	windowSize  time.Duration
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:       limit,
		windowStart: time.Now(),
		windowSize:  window,
	}
}

func (fw *FixedWindowLimiter) Allow() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	if time.Since(fw.windowStart) >= fw.windowSize {
		// New window: reset the counter.
		fw.count = 0
		fw.windowStart = time.Now()
	}

	if fw.count >= fw.limit {
		return false
	}

	fw.count++
	return true
}

func main() {
	// --- Token Bucket demo ---
	fmt.Println("=== Token Bucket (capacity=5, rate=2/s) ===")
	bucket := NewTokenBucket(5, 2)

	// Burst of 7 requests — first 5 pass, last 2 are dropped.
	for i := 1; i <= 7; i++ {
		if bucket.Allow() {
			slog.Info("request allowed", "request", i)
		} else {
			slog.Warn("request DENIED (bucket empty)", "request", i)
		}
	}

	// Wait 1 second — bucket refills with 2 tokens.
	fmt.Println("\n  [waiting 1 second for refill...]")
	time.Sleep(1 * time.Second)

	for i := 8; i <= 10; i++ {
		if bucket.Allow() {
			slog.Info("request allowed", "request", i)
		} else {
			slog.Warn("request DENIED", "request", i)
		}
	}

	// --- Fixed Window demo ---
	fmt.Println("\n=== Fixed Window Counter (limit=3 per 500ms) ===")
	fw := NewFixedWindowLimiter(3, 500*time.Millisecond)

	for i := 1; i <= 5; i++ {
		if fw.Allow() {
			slog.Info("request allowed", "request", i)
		} else {
			slog.Warn("request DENIED (window full)", "request", i)
		}
	}

	fmt.Println("\n  [waiting 500ms for new window...]")
	time.Sleep(500 * time.Millisecond)

	for i := 6; i <= 8; i++ {
		if fw.Allow() {
			slog.Info("request allowed", "request", i)
		} else {
			slog.Warn("request DENIED", "request", i)
		}
	}

	// ============================================================
	// PRODUCTION TIP: golang.org/x/time/rate
	// ============================================================
	// In production, use the battle-tested standard library extension:
	//
	//   import "golang.org/x/time/rate"
	//
	//   limiter := rate.NewLimiter(rate.Limit(10), 20) // 10 req/s, burst of 20
	//
	//   if !limiter.Allow() {
	//       http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
	//       return
	//   }
	//
	// This is exactly what the production project's middleware/ratelimit.go uses!
	fmt.Println("\nSee 04_Production_Project/internal/middleware/ratelimit.go for the production version using golang.org/x/time/rate.")
}
