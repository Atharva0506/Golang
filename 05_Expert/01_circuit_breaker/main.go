package main

import (
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"
)

// ============================================================
// Circuit Breaker Pattern
//
// Problem: When a downstream service is degraded, every request
// to it wastes a goroutine, a connection, and latency budget.
//
// Solution: Wrap calls in a circuit breaker.
// - CLOSED  → normal operation; failures are counted.
// - OPEN    → calls are rejected immediately without touching the service.
// - HALF-OPEN → one probe request is allowed; if it succeeds, reset to CLOSED.
// ============================================================

// State represents the three possible circuit breaker states.
type State int

const (
	StateClosed   State = iota // normal operation
	StateOpen                  // fast-failing
	StateHalfOpen              // testing recovery
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF-OPEN"
	default:
		return "UNKNOWN"
	}
}

// ErrCircuitOpen is returned when the breaker is OPEN and rejects a call.
var ErrCircuitOpen = errors.New("circuit breaker is OPEN — call rejected")

// CircuitBreaker wraps a function call with a failure counter and state machine.
type CircuitBreaker struct {
	mu sync.Mutex

	state            State
	failureCount     int
	failureThreshold int       // how many failures before opening
	resetTimeout     time.Duration
	lastFailure      time.Time
}

// NewCircuitBreaker creates a breaker that opens after `threshold` consecutive
// failures and attempts recovery after `resetTimeout`.
func NewCircuitBreaker(threshold int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            StateClosed,
		failureThreshold: threshold,
		resetTimeout:     resetTimeout,
	}
}

// Call executes fn inside the circuit breaker.
// Returns ErrCircuitOpen immediately when the breaker is OPEN.
func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()

	switch cb.state {
	case StateOpen:
		// Check whether we should transition to HALF-OPEN.
		if time.Since(cb.lastFailure) >= cb.resetTimeout {
			slog.Info("circuit breaker transitioning to HALF-OPEN")
			cb.state = StateHalfOpen
		} else {
			cb.mu.Unlock()
			return ErrCircuitOpen
		}
	case StateHalfOpen:
		// Allow exactly one probe through — hold the lock through the call.
	}

	cb.mu.Unlock()

	// Execute the actual call.
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failureCount++
		cb.lastFailure = time.Now()

		if cb.state == StateHalfOpen || cb.failureCount >= cb.failureThreshold {
			slog.Warn("circuit breaker OPENING", "failures", cb.failureCount)
			cb.state = StateOpen
		}
		return err
	}

	// Success → reset.
	if cb.state == StateHalfOpen {
		slog.Info("circuit breaker CLOSING after successful probe")
	}
	cb.state = StateClosed
	cb.failureCount = 0
	return nil
}

func (cb *CircuitBreaker) State() State {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}

// ============================================================
// Demo
// ============================================================

var callCount int

// unstableService fails for the first 4 calls, then succeeds.
func unstableService() error {
	callCount++
	if callCount <= 4 {
		return fmt.Errorf("service error (call %d)", callCount)
	}
	return nil
}

func main() {
	cb := NewCircuitBreaker(3, 200*time.Millisecond)

	for i := 1; i <= 10; i++ {
		err := cb.Call(unstableService)
		slog.Info("call result",
			"attempt", i,
			"error", err,
			"breaker_state", cb.State(),
		)
		time.Sleep(50 * time.Millisecond)

		// After attempt 5, wait long enough for the reset timeout to expire.
		if i == 5 {
			slog.Info("waiting for reset timeout...")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
