package main

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
)

// ============================================================
// Generic In-Process Event Bus
//
// The EventBus decouples publishers from subscribers.
// Publishers fire events without knowing who is listening.
// Subscribers register handlers for specific event types.
//
// Pattern: Observer / Pub-Sub
// ============================================================

// Handler is a function that processes an event of type T.
type Handler[T any] func(ctx context.Context, event T)

// subscription wraps a handler with an ID for safe removal.
type subscription[T any] struct {
	id      int
	handler Handler[T]
}

// EventBus is a generic, goroutine-safe publish/subscribe event bus.
type EventBus[T any] struct {
	mu      sync.RWMutex
	subs    []subscription[T]
	nextID  int
	channel chan eventEnvelope[T]
}

type eventEnvelope[T any] struct {
	ctx   context.Context
	event T
}

// NewEventBus creates a new bus and starts the internal dispatch goroutine.
func NewEventBus[T any](bufferSize int) *EventBus[T] {
	bus := &EventBus[T]{
		channel: make(chan eventEnvelope[T], bufferSize),
	}
	go bus.dispatch()
	return bus
}

// dispatch is the single goroutine that fans out events to all subscribers.
// Using a single dispatcher goroutine means handlers see events in order.
func (b *EventBus[T]) dispatch() {
	for envelope := range b.channel {
		b.mu.RLock()
		handlers := make([]Handler[T], len(b.subs))
		for i, sub := range b.subs {
			handlers[i] = sub.handler
		}
		b.mu.RUnlock()

		for _, h := range handlers {
			h(envelope.ctx, envelope.event)
		}
	}
}

// Subscribe registers a handler and returns an unsubscribe function.
func (b *EventBus[T]) Subscribe(h Handler[T]) (unsubscribe func()) {
	b.mu.Lock()
	id := b.nextID
	b.nextID++
	b.subs = append(b.subs, subscription[T]{id: id, handler: h})
	b.mu.Unlock()

	return func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		for i, sub := range b.subs {
			if sub.id == id {
				b.subs = append(b.subs[:i], b.subs[i+1:]...)
				return
			}
		}
	}
}

// Publish sends an event to all subscribers (non-blocking if buffer not full).
func (b *EventBus[T]) Publish(ctx context.Context, event T) {
	b.channel <- eventEnvelope[T]{ctx: ctx, event: event}
}

// Close shuts down the dispatch loop.
func (b *EventBus[T]) Close() {
	close(b.channel)
}

// ============================================================
// Domain events for a trading system
// ============================================================

type TradeSignal struct {
	Symbol string
	Action string // "buy" or "sell"
	Price  int
}

type UserRegistered struct {
	Email string
	Role  string
}

func main() {
	ctx := context.Background()

	// --- Example 1: Trade signal bus (typed — only TradeSignal events) ---
	signalBus := NewEventBus[TradeSignal](100)
	defer signalBus.Close()

	// Subscribe two independent handlers
	unsub1 := signalBus.Subscribe(func(ctx context.Context, s TradeSignal) {
		slog.Info("Logger: trade signal received",
			"symbol", s.Symbol, "action", s.Action, "price", s.Price)
	})

	unsub2 := signalBus.Subscribe(func(ctx context.Context, s TradeSignal) {
		if s.Action == "buy" && s.Price > 60000 {
			slog.Warn("RiskEngine: HIGH-PRICE BUY detected!", "symbol", s.Symbol, "price", s.Price)
		}
	})

	signalBus.Publish(ctx, TradeSignal{Symbol: "BTC", Action: "buy", Price: 65000})
	signalBus.Publish(ctx, TradeSignal{Symbol: "ETH", Action: "sell", Price: 3200})
	signalBus.Publish(ctx, TradeSignal{Symbol: "BTC", Action: "buy", Price: 50000})

	// Unsubscribe the risk engine — it will no longer receive events
	unsub2()
	fmt.Println("\n[Risk engine unsubscribed]")
	signalBus.Publish(ctx, TradeSignal{Symbol: "BTC", Action: "buy", Price: 70000})

	// Drain the buffer before moving on
	for len(signalBus.channel) > 0 {
	}

	// --- Example 2: User registration bus ---
	userBus := NewEventBus[UserRegistered](50)
	defer userBus.Close()

	userBus.Subscribe(func(ctx context.Context, u UserRegistered) {
		slog.Info("WelcomeEmailWorker: sending email", "email", u.Email)
	})

	userBus.Subscribe(func(ctx context.Context, u UserRegistered) {
		slog.Info("AuditLog: new user registered", "email", u.Email, "role", u.Role)
	})

	userBus.Publish(ctx, UserRegistered{Email: "luffy@onepiece.com", Role: "trader"})
	userBus.Publish(ctx, UserRegistered{Email: "zoro@onepiece.com", Role: "admin"})

	// Give dispatch goroutines time to finish
	for len(userBus.channel) > 0 {
	}
	_ = unsub1
	fmt.Println("\nDone! Both buses processed all events.")
}
