package main

import (
	"fmt"
	"log/slog"
	"time"
)

// ============================================================
// Event Sourcing
//
// Problem: Traditional CRUD systems overwrite state. Once you
// update a record, the previous state is gone — no audit trail,
// no ability to replay what happened, no temporal queries.
//
// Solution: Never update state directly.
// Store an immutable, append-only log of domain EVENTS.
// The current state is always derived by replaying the event log.
//
// Benefits:
//   - Complete audit trail (every state change is recorded)
//   - Temporal queries (what was the balance on Tuesday?)
//   - Event replay (rebuild read models from scratch)
//   - Debug production issues by replaying exact event sequences
// ============================================================

// ============================================================
// Events — immutable facts that happened in the past
// (named in past tense: SignalCreated, NOT CreateSignal)
// ============================================================

type EventType string

const (
	EventSignalCreated   EventType = "SignalCreated"
	EventSignalCancelled EventType = "SignalCancelled"
	EventPriceUpdated    EventType = "PriceUpdated"
)

// Event is the envelope that wraps every domain event.
type Event struct {
	ID        int
	Type      EventType
	Timestamp time.Time
	Data      interface{}
}

// SignalCreatedData carries the payload for a SignalCreated event.
type SignalCreatedData struct {
	SignalID string
	Symbol   string
	Action   string
	Price    int
}

// SignalCancelledData carries the payload for a SignalCancelled event.
type SignalCancelledData struct {
	SignalID string
	Reason   string
}

// PriceUpdatedData carries the payload for a PriceUpdated event.
type PriceUpdatedData struct {
	SignalID string
	NewPrice int
}

// ============================================================
// Event Store — append-only log
// ============================================================

type EventStore struct {
	events []Event
	nextID int
}

// Append adds a new event to the store.
func (es *EventStore) Append(eventType EventType, data interface{}) Event {
	es.nextID++
	e := Event{
		ID:        es.nextID,
		Type:      eventType,
		Timestamp: time.Now(),
		Data:      data,
	}
	es.events = append(es.events, e)
	slog.Info("event stored", "id", e.ID, "type", e.Type)
	return e
}

// GetAll returns a copy of the full event log.
func (es *EventStore) GetAll() []Event {
	out := make([]Event, len(es.events))
	copy(out, es.events)
	return out
}

// ============================================================
// Aggregate — the current state rebuilt by replaying events
// ============================================================

// SignalState is the current in-memory state of a signal aggregate.
type SignalState struct {
	ID        string
	Symbol    string
	Action    string
	Price     int
	Cancelled bool
	Version   int // how many events have been applied
}

// Apply updates SignalState by applying a single event.
// This function is pure — no side effects, no I/O.
func (s *SignalState) Apply(e Event) {
	s.Version++
	switch e.Type {
	case EventSignalCreated:
		d := e.Data.(SignalCreatedData)
		s.ID = d.SignalID
		s.Symbol = d.Symbol
		s.Action = d.Action
		s.Price = d.Price

	case EventSignalCancelled:
		s.Cancelled = true

	case EventPriceUpdated:
		d := e.Data.(PriceUpdatedData)
		s.Price = d.NewPrice
	}
}

// Rebuild reconstructs the current state by replaying ALL events for a signal.
func Rebuild(signalID string, events []Event) *SignalState {
	state := &SignalState{}
	for _, e := range events {
		// Filter to events that belong to this aggregate.
		switch d := e.Data.(type) {
		case SignalCreatedData:
			if d.SignalID != signalID {
				continue
			}
		case SignalCancelledData:
			if d.SignalID != signalID {
				continue
			}
		case PriceUpdatedData:
			if d.SignalID != signalID {
				continue
			}
		}
		state.Apply(e)
	}
	return state
}

// ============================================================
// Demo
// ============================================================

func main() {
	store := &EventStore{}

	// 1. A signal is created.
	store.Append(EventSignalCreated, SignalCreatedData{
		SignalID: "sig-001", Symbol: "BTC", Action: "buy", Price: 65000,
	})

	// 2. The price changes.
	store.Append(EventPriceUpdated, PriceUpdatedData{
		SignalID: "sig-001", NewPrice: 66000,
	})

	// 3. A second signal is created.
	store.Append(EventSignalCreated, SignalCreatedData{
		SignalID: "sig-002", Symbol: "ETH", Action: "sell", Price: 3200,
	})

	// 4. The first signal is cancelled.
	store.Append(EventSignalCancelled, SignalCancelledData{
		SignalID: "sig-001", Reason: "market moved",
	})

	// ---- Rebuild current state from events ---
	fmt.Println("\n=== Rebuilding state from event log ===")

	sig001 := Rebuild("sig-001", store.GetAll())
	slog.Info("sig-001 current state",
		"symbol", sig001.Symbol,
		"action", sig001.Action,
		"price", sig001.Price,
		"cancelled", sig001.Cancelled,
		"version", sig001.Version,
	)

	sig002 := Rebuild("sig-002", store.GetAll())
	slog.Info("sig-002 current state",
		"symbol", sig002.Symbol,
		"action", sig002.Action,
		"price", sig002.Price,
		"cancelled", sig002.Cancelled,
		"version", sig002.Version,
	)

	// ---- Temporal query: what was sig-001's price after event 2? ---
	fmt.Println("\n=== Temporal query: sig-001 after event #2 ===")
	firstTwoEvents := store.GetAll()[:2]
	historical := Rebuild("sig-001", firstTwoEvents)
	slog.Info("sig-001 at version 2",
		"price", historical.Price,
		"cancelled", historical.Cancelled,
	)

	// ---- Full audit log ---
	fmt.Println("\n=== Full Event Log ===")
	for _, e := range store.GetAll() {
		slog.Info("event", "id", e.ID, "type", e.Type, "data", fmt.Sprintf("%+v", e.Data))
	}
}
