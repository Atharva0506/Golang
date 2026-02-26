package main

import (
	"context"
	"errors"
)

// ==========================================
// 🚀 03_Advanced: Master Integration Test
// ==========================================
// This test simulates a production "Event Processing Service".
// You must implement the missing logic using Advanced Go Concepts!

// Task 1: Interfaces (04_interfaces_advanced)
// Create an interface named 'MessageStore' that has two methods:
// 1. FetchMessages(limit int) ([]string, error)
// 2. MarkAsProcessed(id int) error
type MessageStore interface {
	// TODO: Define the methods here
}

// Task 2: Architecture & Design Patterns (06_design_patterns & 10_testing_mocks)
// Create the 'EventProcessor' struct.
// It MUST have a field named 'store' of type MessageStore.
// It MUST have a field named 'maxWorkers' of type int.
type EventProcessor struct {
	// TODO: Define the fields here
}

// Write a Functional Option named 'WithWorkers(count int)' that returns a function to set 'maxWorkers'.
type ProcessorOption func(*EventProcessor)

func WithWorkers(count int) ProcessorOption {
	// TODO: Implement the functional option
	return nil
}

// Write the Constructor 'NewEventProcessor(store MessageStore, opts ...ProcessorOption) *EventProcessor'
// Default maxWorkers to 1.
func NewEventProcessor(store MessageStore, opts ...ProcessorOption) *EventProcessor {
	// TODO: Implement the constructor
	return nil
}

// Task 3: Context & Goroutines (01_context & 07_concurrency_patterns)
// Write a method on EventProcessor:
// `func (e *EventProcessor) ProcessAll(ctx context.Context, limit int) error`
// 1. It must call `e.store.FetchMessages(limit)`. If error, return it.
// 2. It must loop through the messages. For each message, it should print "Processing: [message]".
// 3. It MUST use `select { case <-ctx.Done(): ... }` to stop processing immediately and return `context.DeadlineExceeded` if the context expires before the loop finishes!
func (e *EventProcessor) ProcessAll(ctx context.Context, limit int) error {
	// TODO: Implement context-aware processing
	return errors.New("Not implemented")
}

func main() {
	// Run `go test -v` in this directory to see if you pass the senior challenges!
}
