# 05 Expert — Distributed Systems Patterns

This section covers production-grade distributed systems patterns used by senior Go engineers. Each module is a self-contained, runnable example with detailed comments explaining the *why* behind every design decision.

## Modules

| Module | Pattern | Key Problem Solved |
|--------|---------|-------------------|
| `01_circuit_breaker` | Circuit Breaker | Prevent cascade failures when a downstream service is degraded |
| `02_saga` | Saga (Orchestration) | Roll back a multi-step distributed transaction when one step fails |
| `03_cqrs` | CQRS | Decouple the write model (business rules) from the read model (query projections) |
| `04_event_sourcing` | Event Sourcing | Replace mutable state with an immutable event log — full audit trail + temporal queries |

## How to Run

Each module is a standalone Go program with no external dependencies:

```bash
go run 01_circuit_breaker/main.go
go run 02_saga/main.go
go run 03_cqrs/main.go
go run 04_event_sourcing/main.go
```

## Reading Order

These patterns often compose together:

```
Event Sourcing  →  CQRS  →  Saga  →  Circuit Breaker
(store events)    (split R/W)  (multi-service txn)  (fault tolerance)
```

A production system like a trading platform might use **all four**:
1. **Event Sourcing** stores every trade signal as an immutable event.
2. **CQRS** projects those events into fast-read dashboards without blocking writes.
3. **Saga** ensures that a payment + inventory + shipment transaction either fully succeeds or fully rolls back.
4. **Circuit Breaker** wraps calls to external payment providers so a slow provider can't take down the whole API.
