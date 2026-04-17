package main

import (
	"context"
	"strings"
	"time"
)

// =========================================================================
// 🚀 04_ADVANCED_TEST: PROFILING & TRACING ASSIGNMENT
// =========================================================================
// Two tasks that cover the two new Advanced modules:
//   - 16_profiling  (performance optimization guided by benchmarks)
//   - 17_tracing    (distributed tracing / context propagation)
//
// Run the tests: go test -v ./...
// Run benchmarks: go test -bench=. -benchmem
//
// =========================================================================

// -------------------------------------------------------------------------
// Task 1: Performance Optimization (16_profiling)
// -------------------------------------------------------------------------
// The function below is intentionally slow.
// Your job is to write a FASTER version called `ConcatFast`.
//
// Rule: `ConcatSlow` must remain unchanged — it is the baseline.
//
// Hints:
//   - Benchmark both with `go test -bench=. -benchmem` to measure the difference.
//   - `strings.Builder` is the idiomatic Go tool for this job.
//   - Pre-allocating with `b.Grow(n)` eliminates reallocations entirely.

// ConcatSlow joins n copies of the word "Go" using naive string concatenation.
// Each `+=` creates a brand-new string — O(n²) allocations!
func ConcatSlow(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += "Go"
	}
	return result
}

// ConcatFast joins n copies of the word "Go" efficiently.
// TODO: Implement using strings.Builder with Grow pre-allocation.
// Expected: zero allocations per benchmark iteration (allocs/op = 0).
func ConcatFast(n int) string {
	// TODO: Implement
	return strings.Repeat("Go", n) // Replace this with a strings.Builder implementation
}

// -------------------------------------------------------------------------
// Task 2: Distributed Tracing (17_tracing)
// -------------------------------------------------------------------------
// Implement a minimal tracing system following the patterns from 17_tracing.
//
// You must provide:
//   1. A `Span` struct with at minimum: Name, StartTime, EndTime, and a map of Attrs.
//   2. A `SpanTracer` struct with a `Start(ctx, name)` method.
//      - Start must inject the span into the returned context.
//   3. A `SpanFromContext(ctx)` function that retrieves the active span from context.
//   4. A `Span.SetAttr(key, value string)` method.
//   5. A `Span.End()` method that records the EndTime.
//
// Design rules:
//   - A child span created from a context that already holds a parent span
//     must carry the parent's TraceID.
//   - SpanFromContext must return nil (not panic) when no span is in context.

// contextKey is an unexported type to avoid context key collisions.
type contextKey string

const activeSpanKey contextKey = "active_span"

// Span represents a single traced operation.
// All fields are required by the tests — fill them in inside Start() and End().
type Span struct {
	Name      string
	TraceID   string
	SpanID    string
	ParentID  string
	StartTime time.Time
	EndTime   time.Time
	Attrs     map[string]string
}

// End records the end time of the span.
func (s *Span) End() {
	// TODO: set s.EndTime = time.Now()
}

// SetAttr attaches a key-value pair to the span.
func (s *Span) SetAttr(key, value string) {
	// TODO: initialise s.Attrs if nil, then store key → value
}

// SpanTracer creates new spans.
type SpanTracer struct {
	ServiceName string
}

// Start creates a new span, stores it in the context, and returns both.
// If a parent span already exists in ctx, the new span must be its child
// (same TraceID, parent's SpanID as ParentID).
func (t *SpanTracer) Start(ctx context.Context, name string) (context.Context, *Span) {
	// TODO: Implement — create a Span, inject it into a new context, return both.
	return ctx, &Span{} // replace with real implementation
}

// SpanFromContext retrieves the active span from ctx.
// Returns nil if no span has been stored.
func SpanFromContext(ctx context.Context) *Span {
	// TODO: read the span from ctx using activeSpanKey
	return nil
}

func main() {
	// The ultimate evaluation! Run `go test -v` to check your code.
	_ = time.Now()
}
