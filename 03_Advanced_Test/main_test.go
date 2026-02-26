package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

// Mock implementation of the Notifier interface for testing!
type MockNotifier struct {
	Count int32
}

func (m *MockNotifier) Notify(ctx context.Context, u *User) error {
	// Simulate work
	select {
	case <-time.After(10 * time.Millisecond):
		atomic.AddInt32(&m.Count, 1)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Test Pipeline
func TestAdvancedPipeline(t *testing.T) {
	// --- TASK 1: Pointers & Buffers ---
	t.Run("Task 1: AllocateUserBatch", func(t *testing.T) {
		batch := AllocateUserBatch(100)
		if batch == nil {
			t.Fatalf("Task 1 Failed: Returned nil")
		}
		if cap(*batch) != 100 {
			t.Errorf("Task 1 Failed: Expected capacity 100, got %d. Did you use make() correctly?", cap(*batch))
		}
	})

	// --- TASK 2 & 3: Interfaces, Mocks, and Functional Options ---
	t.Run("Task 2 & 3: Architecture", func(t *testing.T) {
		mock := &MockNotifier{}
		svc := NewNotificationService(mock, WithWorkers(5))

		if svc == nil {
			t.Fatalf("Task 3 Failed: Returned nil service")
		}
		if svc.workerCount != 5 {
			t.Errorf("Task 3 Failed: Expected workerCount 5, got %d", svc.workerCount)
		}
	})

	// --- TASK 4: Context & Concurrency ---
	t.Run("Task 4: Concurrency and Timeout", func(t *testing.T) {
		mock := &MockNotifier{}
		svc := NewNotificationService(mock, WithWorkers(5))

		users := []*User{{ID: 1}, {ID: 2}, {ID: 3}}

		// Force a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
		defer cancel()

		err := svc.ProcessUsers(ctx, users)
		if err == nil {
			t.Errorf("Task 4 Failed: Expected context timeout error, but got nil")
		}
	})

	// --- TASK 5: Reflection ---
	t.Run("Task 5: Reflection", func(t *testing.T) {
		tags := ExtractDBTags(User{})
		if len(tags) != 2 || tags[0] != "primary_key" || tags[1] != "email_address" {
			t.Errorf("Task 5 Failed: Failed to extract struct tags correctly using reflection")
		}
	})

	// --- TASK 6: Middleware ---
	t.Run("Task 6: Middleware", func(t *testing.T) {
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
		mw := RequestLogger(nextHandler)

		if mw == nil {
			t.Fatalf("Task 6 Failed: Middleware returned nil")
		}

		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()

		mw.ServeHTTP(rec, req)

		if rec.Header().Get("X-Middleware-Applied") != "true" {
			t.Errorf("Task 6 Failed: Header not applied by middleware")
		}
	})

	// --- TASK 8: Graceful Shutdown ---
	t.Run("Task 8: OS Signals", func(t *testing.T) {
		ch := SetupShutdownChannel()
		if ch == nil {
			t.Fatalf("Task 8 Failed: Channel is nil")
		}
		if cap(ch) == 0 {
			t.Errorf("Task 8 Failed: OS Signal channel should be buffered!")
		}
	})
}

// --- TASK 11: Benchmarking ---
// We write the benchmark here so the user can run `go test -bench=.`
func BenchmarkAllocateUserBatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AllocateUserBatch(100)
	}
}
