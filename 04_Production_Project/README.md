# Production-Grade Golang Backend Architecture

## Overview
This repository contains a skeleton for a highly scalable, production-ready backend system written in Go. It enforces **Clean Architecture** principles, prioritizing modularity, extreme testability, and strict separation of concerns.

## Features
- **Protocols Supported**: REST APIs, gRPC Services, WebSockets
- **Infrastructure**: Dockerized (Multi-stage), `docker-compose` ready
- **Architecture**: Modular layout (Clean Architecture / Domain-Driven Design)
- **Security & Reliability**: JWT Authentication, Rate Limiting, Graceful Shutdowns, Custom App Errors
- **Observability**: Production-grade structured JSON logging
- **Concurrency**: Independent background worker processing
- **Testing**: Unified layer for Unit Tests, DB-Integration Tests, and mock generation.

---

## Folder Structure & Responsibilities

```text
├── cmd/
│   ├── api/       # Entrypoint for the HTTP/gRPC/WS web server
│   └── worker/    # Entrypoint for the background job processing daemon
├── internal/      # Private application and domain code (Not importable by other Git repos)
│   ├── config/    # Environment variable mapping and yaml parsing
│   ├── models/    # Core domain structs (e.g., User, Order) alongside DB/JSON tags
│   ├── repository/# Database interactions. Maps rows to domain models
│   ├── service/   # Pure business logic. Glues requests, repositories, and third parties together
│   ├── delivery/  # Network layer boundaries (HTTP Router, gRPC PB definitions, Websocket Hub)
│   ├── middleware/# Request interceptors (Auth, Rate Limiting, System logging)
│   ├── worker/    # Background task definitions to keep HTTP responders fast
│   └── di/        # Dependency Injection definitions (Injecting Repos -> Services -> Handlers)
├── pkg/           # Public library code (Logging wrappers, Generic DB Connectors, Auth parsers)
├── tests/         # Broad tests that cross internal boundaries (e.g., real DB Integration Tests)
```

**Why this structure?**
It strictly prevents circular dependencies and isolates your core business logic (`internal/service`) away from any specific networking transport (`internal/delivery`) or specific database engine (`internal/repository`). 

---

## Development Workflow

### How to Run the Project
1. Copy the example config: `cp .env.example .env`
2. Spin up local dependencies: `docker-compose up -d db cache`
3. Run the API: `make run` or `go run cmd/api/main.go`

### How to Run Tests
- Unit Tests (Fast): `go test -short ./...`
- Integration Tests (Slower, requires active DB): `go test -run Integration ./tests/integration/...`

### How to Build
The provided `Makefile` exposes specific build scripts, yielding a single, stripped binary.
`make build`

### Adding a New Module (E.g. "Products")
1. Define the `Product` struct in `internal/models/product.go`
2. Create `internal/repository/product_repo.go` (SQL transactions)
3. Create `internal/service/product_service.go` (Business rules)
4. Expose endpoints via `internal/delivery/http/product_handler.go`
5. Wire them together inside `internal/di/container.go`

---

## Architecture Best Practices Included
- **Dependency Injection**: Structs accept interfaces. For example, the `UserService` requires a `UserRepository` interface. This enables isolated unit testing.
- **Context Plumbing**: `context.Context` is the first parameter to every database and network call, allowing for instant cascading timeouts (preventing memory leaks).
- **Graceful Shutdown**: The `main.go` files intercept `os.Interrupt` and block exiting until pending active HTTP requests finish writing their responses.

---

## Recommended Project Ideas for this Architecture

If you want to fill these shell files with real logic, here are 3 complex backend systems uniquely suited for this architecture:

1. **Real-Time Delivery Tracker Backend**
   - **REST API**: For User Registration, Order placement, and Payment confirmation
   - **WebSockets**: Live GPS coordinates broadcasting from drivers to customers
   - **Workers**: Handling offline push notifications through APNs/FCM
   
2. **Microservice Authentication Provider (Like Auth0)**
   - **gRPC**: Blazing fast inner-network communication where external microservices check `ValidateToken`
   - **REST API**: For managing OAuth credentials and Social Logins
   - **Background workers**: Sending magic login link emails safely without slowing the API
   - **Middleware**: Aggressive rate limiting against brute force attempts

3. **High-Frequency Trading & Notification Bot**
   - **WebSockets (Ingest)**: Consuming massive firehoses of live Crypto/Stock prices
   - **Services**: Running pattern-matching algorithms in memory
   - **Database Repo**: Upserting the triggered signals instantly
   - **gRPC/REST**: Triggering webhook alerts out to connected partner services
