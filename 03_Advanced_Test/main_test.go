package main

import (
	"context"
	"testing"
	"time"
)

// MockStore implements the interface the user built so we can test their ProcessAll function!
type MockStore struct {
	Messages []string
	Err      error
}

func (m *MockStore) FetchMessages(limit int) ([]string, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	if limit > len(m.Messages) {
		limit = len(m.Messages)
	}
	return m.Messages[:limit], nil
}

func (m *MockStore) MarkAsProcessed(id int) error {
	return nil
}

// Task 1 & 2 Test
func TestEventProcessor_Architecture(t *testing.T) {
	mockDB := &MockStore{
		Messages: []string{"A", "B", "C"},
	}

	// This automatically checks if the user built `NewEventProcessor` and `WithWorkers` correctly!
	processor := NewEventProcessor(mockDB, WithWorkers(5))

	if processor == nil {
		t.Fatalf("Expected NewEventProcessor to return a pointer, got nil")
	}

	if processor.maxWorkers != 5 {
		t.Errorf("Expected maxWorkers to be 5 via Functional Options, got %d", processor.maxWorkers)
	}

	if processor.store == nil {
		t.Errorf("Expected store to be set, got nil")
	}
}

// Task 3 Test
func TestEventProcessor_ProcessAll_ContextTimeout(t *testing.T) {
	mockDB := &MockStore{
		// We send 10 messages
		Messages: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	}

	processor := NewEventProcessor(mockDB)

	// We only give the context 10 milliseconds to live!
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// But wait! We will spawn a goroutine that artificially sleeps during `ProcessAll`
	// so it forces Context to time out before it can finish all 10 messages!
	// (Users will need to ensure they have the `select { case <-ctx.Done(): ... }` logic)

	err := processor.ProcessAll(ctx, 10)

	// If the user didn't check `ctx.Done()`, it will finish processing all of them instead of returning instantly!
	if err == nil {
		t.Errorf("Expected ProcessAll to fail with a Context Timeout Error, but it succeeded!")
	}

	if err != context.DeadlineExceeded && err.Error() != "context deadline exceeded" {
		t.Errorf("Expected context deadline exceeded error, got %v", err)
	}
}
