package main

import (
	"context"
	"strings"
	"testing"
)

// ============================================================
// Task 1: Benchmarks — ConcatSlow vs ConcatFast
// ============================================================

func BenchmarkConcatSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConcatSlow(500)
	}
}

func BenchmarkConcatFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConcatFast(500)
	}
}

// TestConcatFastCorrectness checks that ConcatFast produces the correct output.
func TestConcatFastCorrectness(t *testing.T) {
	result := ConcatFast(5)
	expected := strings.Repeat("Go", 5)
	if result != expected {
		t.Errorf("ConcatFast(5) = %q, want %q", result, expected)
	}
}

// TestConcatFastIsEfficient checks that ConcatFast uses at most 1 allocation
// — i.e., it uses strings.Builder, not naive += concatenation.
func TestConcatFastIsEfficient(b *testing.T) {
	result := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ConcatFast(500)
		}
	})

	if result.AllocsPerOp() > 1 {
		b.Errorf("ConcatFast should use ≤1 allocation per op, got %d. Did you use strings.Builder?", result.AllocsPerOp())
	}
}

// ============================================================
// Task 2: Tracing tests
// ============================================================

func TestSpanTracer_RootSpan(t *testing.T) {
	tracer := &SpanTracer{ServiceName: "test-service"}
	ctx := context.Background()

	newCtx, span := tracer.Start(ctx, "root-operation")

	if span == nil {
		t.Fatal("Start returned a nil span")
	}
	if newCtx == ctx {
		t.Error("Start must return a NEW context with the span stored inside it")
	}
	span.End()
}

func TestSpanTracer_ChildInheritsTraceID(t *testing.T) {
	tracer := &SpanTracer{ServiceName: "test-service"}
	ctx := context.Background()

	parentCtx, parent := tracer.Start(ctx, "parent")
	_, child := tracer.Start(parentCtx, "child")

	if parent == nil || child == nil {
		t.Fatal("Start returned a nil span")
	}
	if parent.TraceID == "" {
		t.Error("parent TraceID must not be empty")
	}
	if child.TraceID != parent.TraceID {
		t.Errorf("child span must inherit TraceID from parent: got %q, want %q",
			child.TraceID, parent.TraceID)
	}
	if child.ParentID != parent.SpanID {
		t.Errorf("child.ParentID must equal parent.SpanID: got %q, want %q",
			child.ParentID, parent.SpanID)
	}
}

func TestSpanFromContext_NilWhenEmpty(t *testing.T) {
	span := SpanFromContext(context.Background())
	if span != nil {
		t.Error("SpanFromContext must return nil when no span is in context")
	}
}

func TestSpanFromContext_ReturnsActiveSpan(t *testing.T) {
	tracer := &SpanTracer{ServiceName: "test-service"}
	ctx, _ := tracer.Start(context.Background(), "test-op")

	retrieved := SpanFromContext(ctx)
	if retrieved == nil {
		t.Fatal("SpanFromContext must return the active span stored by Start")
	}
	if !strings.HasSuffix(retrieved.Name, "test-op") {
		t.Errorf("retrieved span name = %q, expected it to contain 'test-op'", retrieved.Name)
	}
}

func TestSpan_SetAttr(t *testing.T) {
	tracer := &SpanTracer{ServiceName: "test-service"}
	_, span := tracer.Start(context.Background(), "op")
	span.SetAttr("symbol", "BTC")
	span.SetAttr("price", "65000")

	if span.Attrs == nil {
		t.Fatal("Attrs map must not be nil after SetAttr")
	}
	if span.Attrs["symbol"] != "BTC" {
		t.Errorf("expected Attrs[symbol] = BTC, got %q", span.Attrs["symbol"])
	}
	if span.Attrs["price"] != "65000" {
		t.Errorf("expected Attrs[price] = 65000, got %q", span.Attrs["price"])
	}
}

func TestSpan_EndRecordsTime(t *testing.T) {
	tracer := &SpanTracer{ServiceName: "test-service"}
	_, span := tracer.Start(context.Background(), "timed-op")

	if span.EndTime.IsZero() != true {
		t.Error("EndTime must be zero before End() is called")
	}

	span.End()

	if span.EndTime.IsZero() {
		t.Error("EndTime must be set after End() is called")
	}
	if span.EndTime.Before(span.StartTime) {
		t.Error("EndTime must not be before StartTime")
	}
}
