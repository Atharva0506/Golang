package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

// ============================================================
// Lightweight manual tracing — the same concepts as OpenTelemetry
// without the external dependency.
//
// In production you would replace this package with:
//   go.opentelemetry.io/otel
// and export spans to Jaeger / Honeycomb / Datadog / etc.
// ============================================================

// contextKey is an unexported type to avoid key collisions in context.
type contextKey string

const activeSpanKey contextKey = "active_span"

// Span represents a single unit of work being traced.
type Span struct {
	TraceID   string
	SpanID    string
	ParentID  string
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Attrs     map[string]string
}

// End records the finish time and logs the span.
func (s *Span) End() {
	s.EndTime = time.Now()
	duration := s.EndTime.Sub(s.StartTime)
	slog.Info("span finished",
		"trace_id", s.TraceID,
		"span_id", s.SpanID,
		"parent_id", s.ParentID,
		"name", s.Name,
		"duration_ms", duration.Milliseconds(),
		"attrs", s.Attrs,
	)
}

// SetAttr adds a key-value attribute to the span (equivalent to otel.Span.SetAttributes).
func (s *Span) SetAttr(key, value string) {
	if s.Attrs == nil {
		s.Attrs = make(map[string]string)
	}
	s.Attrs[key] = value
}

// Tracer creates spans. In a real OTEL setup this would hold the exporter.
type Tracer struct {
	ServiceName string
}

// Start creates a new child span, injecting it into the returned context.
// If there is already a span in the context, the new span becomes its child.
func (t *Tracer) Start(ctx context.Context, name string) (context.Context, *Span) {
	parentID := ""
	traceID := uuid.New().String()

	if parent, ok := ctx.Value(activeSpanKey).(*Span); ok {
		parentID = parent.SpanID
		traceID = parent.TraceID // child shares the parent's trace ID
	}

	span := &Span{
		TraceID:   traceID,
		SpanID:    uuid.New().String(),
		ParentID:  parentID,
		Name:      fmt.Sprintf("%s/%s", t.ServiceName, name),
		StartTime: time.Now(),
		Attrs:     make(map[string]string),
	}

	newCtx := context.WithValue(ctx, activeSpanKey, span)
	return newCtx, span
}

// SpanFromContext returns the active span stored in the context, or nil.
func SpanFromContext(ctx context.Context) *Span {
	span, _ := ctx.Value(activeSpanKey).(*Span)
	return span
}

// ============================================================
// Example application code instrumented with tracing
// ============================================================

var tracer = &Tracer{ServiceName: "trading-bot"}

// handleRequest represents an incoming HTTP/gRPC request.
func handleRequest(ctx context.Context, symbol string) error {
	// 1. Start a root span for the entire request.
	ctx, span := tracer.Start(ctx, "HandleRequest")
	defer span.End()
	span.SetAttr("symbol", symbol)

	// 2. Call a child function — context propagates the trace automatically.
	price, err := fetchPrice(ctx, symbol)
	if err != nil {
		span.SetAttr("error", err.Error())
		return err
	}

	err = saveSignal(ctx, symbol, price)
	if err != nil {
		span.SetAttr("error", err.Error())
		return err
	}

	span.SetAttr("price", fmt.Sprintf("%d", price))
	return nil
}

// fetchPrice simulates an outbound API call to a price feed.
func fetchPrice(ctx context.Context, symbol string) (int, error) {
	// Child span: its parent_id will point to the HandleRequest span.
	ctx, span := tracer.Start(ctx, "FetchPrice")
	defer span.End()
	span.SetAttr("symbol", symbol)

	// Simulate network latency
	time.Sleep(20 * time.Millisecond)
	span.SetAttr("source", "binance-api")
	return 65000, nil
}

// saveSignal simulates a database write.
func saveSignal(ctx context.Context, symbol string, price int) error {
	// Another child span — shares the same trace_id.
	_, span := tracer.Start(ctx, "SaveSignal")
	defer span.End()
	span.SetAttr("symbol", symbol)
	span.SetAttr("table", "signals")

	time.Sleep(5 * time.Millisecond)
	return nil
}

func main() {
	slog.Info("--- Distributed Tracing Demo ---")
	slog.Info("All spans with the same trace_id belong to one request.")
	slog.Info("")

	ctx := context.Background()

	if err := handleRequest(ctx, "BTC"); err != nil {
		slog.Error("request failed", "error", err)
	}
}
