package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
)

// ============================================================
// Saga Pattern (Orchestration style)
//
// Problem: A distributed transaction spans multiple services
// (e.g., Payment + Inventory + Shipping). If one step fails,
// the completed steps must be reversed to maintain consistency.
// A traditional 2PC lock is impractical across service boundaries.
//
// Solution: Break the transaction into ordered Steps.
// Each Step has an Execute function AND a Compensate (undo) function.
// If any step fails, the Saga executes compensations in reverse order.
// ============================================================

// Step defines one unit of work within a saga.
type Step struct {
	Name       string
	Execute    func(ctx context.Context) error
	Compensate func(ctx context.Context) error
}

// Saga orchestrates a sequence of Steps with automatic compensation.
type Saga struct {
	steps []Step
}

// AddStep appends a step to the saga.
func (s *Saga) AddStep(step Step) {
	s.steps = append(s.steps, step)
}

// Run executes steps in order.
// On first failure it runs all completed compensations in reverse order.
// Returns the original execution error (not compensation errors).
func (s *Saga) Run(ctx context.Context) error {
	completed := make([]Step, 0, len(s.steps))

	for _, step := range s.steps {
		slog.Info("saga: executing step", "step", step.Name)
		if err := step.Execute(ctx); err != nil {
			slog.Error("saga: step failed, starting compensation", "step", step.Name, "error", err)

			for i := len(completed) - 1; i >= 0; i-- {
				c := completed[i]
				slog.Info("saga: compensating", "step", c.Name)
				if cErr := c.Compensate(ctx); cErr != nil {
					// Log but continue — best-effort rollback.
					slog.Error("saga: compensation failed", "step", c.Name, "error", cErr)
				}
			}
			return fmt.Errorf("saga failed at step %q: %w", step.Name, err)
		}
		completed = append(completed, step)
	}

	slog.Info("saga: all steps completed successfully")
	return nil
}

// ============================================================
// Demo: Order fulfilment saga
// ============================================================

func main() {
	// Simulate a successful saga
	slog.Info("=== Happy path ===")
	if err := runOrderSaga(false); err != nil {
		slog.Error("saga failed", "error", err)
	}

	fmt.Println()

	// Simulate a failure at the shipping step
	slog.Info("=== Failure at Shipping step ===")
	if err := runOrderSaga(true); err != nil {
		slog.Error("saga failed", "error", err)
	}
}

func runOrderSaga(failShipping bool) error {
	ctx := context.Background()
	saga := &Saga{}

	// Step 1: Reserve inventory
	reserved := false
	saga.AddStep(Step{
		Name: "ReserveInventory",
		Execute: func(ctx context.Context) error {
			reserved = true
			slog.Info("  inventory reserved")
			return nil
		},
		Compensate: func(ctx context.Context) error {
			reserved = false
			slog.Info("  inventory reservation released")
			return nil
		},
	})

	// Step 2: Charge payment
	charged := false
	saga.AddStep(Step{
		Name: "ChargePayment",
		Execute: func(ctx context.Context) error {
			charged = true
			slog.Info("  payment charged")
			return nil
		},
		Compensate: func(ctx context.Context) error {
			charged = false
			slog.Info("  payment refunded")
			return nil
		},
	})

	// Step 3: Create shipment (may fail)
	saga.AddStep(Step{
		Name: "CreateShipment",
		Execute: func(ctx context.Context) error {
			if failShipping {
				return errors.New("shipping provider unavailable")
			}
			slog.Info("  shipment created")
			return nil
		},
		Compensate: func(ctx context.Context) error {
			slog.Info("  shipment cancelled")
			return nil
		},
	})

	err := saga.Run(ctx)
	slog.Info("final state", "inventory_reserved", reserved, "payment_charged", charged)
	return err
}
