package main

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

// ============================================================
// CQRS — Command Query Responsibility Segregation
//
// Problem: In a high-read system (like a trading dashboard),
// the write model (complex domain logic) and the read model
// (optimised projections) have conflicting requirements.
//
// Solution: Separate the application into:
//   - Command side: validates, applies business rules, stores events.
//   - Query side:   reads from a de-normalised "read model" (projection)
//                   that is optimised purely for read performance.
//
// The two sides are synchronised asynchronously via an event store.
// ============================================================

// ============================================================
// Shared domain types
// ============================================================

type Signal struct {
	ID        string
	Symbol    string
	Action    string
	Price     int
	Timestamp time.Time
}

// ============================================================
// Command side — write model
// ============================================================

// CreateSignalCommand is the intent to create a trade signal.
type CreateSignalCommand struct {
	ID     string
	Symbol string
	Action string
	Price  int
}

// SignalCommandHandler handles write operations and emits events.
type SignalCommandHandler struct {
	mu     sync.Mutex
	store  []*Signal        // simple in-memory event store
	notify chan *Signal      // notifies the query side of new signals
}

func NewSignalCommandHandler() *SignalCommandHandler {
	return &SignalCommandHandler{
		notify: make(chan *Signal, 100),
	}
}

// Handle validates the command, persists the signal, and publishes it.
func (h *SignalCommandHandler) Handle(cmd CreateSignalCommand) error {
	if cmd.Symbol == "" {
		return fmt.Errorf("symbol is required")
	}
	if cmd.Action != "buy" && cmd.Action != "sell" {
		return fmt.Errorf("action must be buy or sell")
	}
	if cmd.Price <= 0 {
		return fmt.Errorf("price must be positive")
	}

	signal := &Signal{
		ID:        cmd.ID,
		Symbol:    cmd.Symbol,
		Action:    cmd.Action,
		Price:     cmd.Price,
		Timestamp: time.Now(),
	}

	h.mu.Lock()
	h.store = append(h.store, signal)
	h.mu.Unlock()

	// Non-blocking publish — the query side is eventually consistent.
	select {
	case h.notify <- signal:
	default:
	}

	slog.Info("command: signal stored", "id", signal.ID, "symbol", signal.Symbol)
	return nil
}

// ============================================================
// Query side — read model (de-normalised projection)
// ============================================================

// SignalSummary is the pre-computed read model — optimised for fast reads.
type SignalSummary struct {
	Symbol    string
	LastPrice int
	BuyCount  int
	SellCount int
}

// SignalQueryHandler maintains the read model by projecting events.
type SignalQueryHandler struct {
	mu       sync.RWMutex
	bySymbol map[string]*SignalSummary
}

func NewSignalQueryHandler() *SignalQueryHandler {
	return &SignalQueryHandler{
		bySymbol: make(map[string]*SignalSummary),
	}
}

// Project updates the read model whenever a new signal arrives.
func (q *SignalQueryHandler) Project(s *Signal) {
	q.mu.Lock()
	defer q.mu.Unlock()

	summary, ok := q.bySymbol[s.Symbol]
	if !ok {
		summary = &SignalSummary{Symbol: s.Symbol}
		q.bySymbol[s.Symbol] = summary
	}

	summary.LastPrice = s.Price
	if s.Action == "buy" {
		summary.BuyCount++
	} else {
		summary.SellCount++
	}
}

// GetSummary returns the pre-computed summary for a symbol — O(1) read.
func (q *SignalQueryHandler) GetSummary(symbol string) (SignalSummary, bool) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	s, ok := q.bySymbol[symbol]
	if !ok {
		return SignalSummary{}, false
	}
	return *s, true
}

// ============================================================
// Wiring — the event bus between command and query sides
// ============================================================

func main() {
	cmdHandler := NewSignalCommandHandler()
	qryHandler := NewSignalQueryHandler()

	// Start the projection goroutine — listens to the event channel
	// and updates the read model asynchronously.
	go func() {
		for signal := range cmdHandler.notify {
			qryHandler.Project(signal)
		}
	}()

	// Issue commands
	commands := []CreateSignalCommand{
		{ID: "1", Symbol: "BTC", Action: "buy", Price: 65000},
		{ID: "2", Symbol: "BTC", Action: "buy", Price: 65500},
		{ID: "3", Symbol: "ETH", Action: "sell", Price: 3200},
		{ID: "4", Symbol: "BTC", Action: "sell", Price: 64000},
		{ID: "5", Symbol: "ETH", Action: "buy", Price: 3300},
	}

	for _, cmd := range commands {
		if err := cmdHandler.Handle(cmd); err != nil {
			slog.Error("command rejected", "error", err)
		}
	}

	// Allow projection to catch up (in production you'd use a proper sync mechanism).
	time.Sleep(10 * time.Millisecond)

	// Query the read model — no joins, no aggregations at query time!
	fmt.Println("\n=== Query Side — Pre-computed Summaries ===")
	for _, sym := range []string{"BTC", "ETH"} {
		if summary, ok := qryHandler.GetSummary(sym); ok {
			slog.Info("summary",
				"symbol", summary.Symbol,
				"last_price", summary.LastPrice,
				"buys", summary.BuyCount,
				"sells", summary.SellCount,
			)
		}
	}
}
