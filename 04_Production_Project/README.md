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

---

## 🚀 Execution Roadmap



### Phase 1: Domain Modeling & Contracts (The "Core")
1. **Define Domain Entities (`internal/models/`)**
   - *Task:* Create the core structs (`User`, `Signal`, `Notification`). 
   - *Why:* These are pure Go structs. No business logic, no SQL. Just data representation with JSON and DB tags.
2. **Define Repository Interfaces (`internal/repository/`)**
   - *Task:* Define what actions the system can perform (e.g., `UserRepository` interface with `Create`, `GetByID`, `UpdateStatus`).
   - *Why:* We define the *contract* of database interactions before writing actual SQL. This allows us to mock the database for testing.

### Phase 2: Business Logic & Unit Testing (The "Brain")
3. **Implement Services (`internal/service/`)**
   - *Task:* Build the business rules (e.g., `AuthService` handling password hashing and JWT generation, `TradingService` matching signals).
   - *Why:* Services take the Repository interfaces as dependencies. They orchestrate the logic but don't know *how* data is stored or *where* it's requested from.
4. **Write Strict Unit Tests (`tests/` or alongside service)**
   - *Task:* Use mocks to test your services. 
   - *Why:* Test edge cases (invalid email, duplicate user, insufficient balance) without ever spinning up a real Postgres database.

### Phase 3: Infrastructure & Integration (The "Storage")
5. **Implement Repositories (`internal/repository/postgres/` or `mysql/`)**
   - *Task:* Write the actual SQL queries (using `database/sql`, `sqlx`, or an ORM like `gorm`) implementing the interfaces defined in Phase 1.
   - *Why:* Now that the business logic works, we plug in the real storage.
6. **Write Database Integration Tests (`tests/integration/`)**
   - *Task:* Spin up a Dockerized test database (e.g., using `testcontainers`), run real queries, and truncate tables between tests.

### Phase 4: Delivery Layer (The "Skin")
7. **Implement Transport Handlers (`internal/delivery/http/` & `internal/delivery/grpc/`)**
   - *Task:* Write HTTP handlers (Gin, Echo, Fiber, or stdlib `net/http`) and gRPC servers.
   - *Why:* Handlers parse JSON/Protobuf requests, call the Service layer, and return formatted responses. They handle *no* business rules, only HTTP/gRPC routing.
8. **Add Middleware (`internal/middleware/`)**
   - *Task:* Inject Auth JWT validation, Rate Limiting, and Request/Response Logging.

### Phase 5: Concurrency & Observability
9. **Setup Background Workers (`cmd/worker/` and `internal/worker/`)**
   - *Task:* Implement Kafka/RabbitMQ or simple Go channels for sending email notifications or processing trading signals asynchronously.
10. **Structured Logging & Metrics**
    - *Task:* Wire up `slog` for structured JSON logs. Add Prometheus metrics for API latencies.

### Phase 6: Assembly & Deployment
11. **Dependency Injection (`internal/di/` & `cmd/api/main.go`)**
    - *Task:* Wire everything together. Initialize the DB, pass it to the Repo -> pass Repo to Service -> pass Service to Handler. Register Handlers to the router.
12. **Graceful Shutdown & Dockerization**
    - *Task:* Implement `os.Signal` catching to shut down the server gracefully. Update `Dockerfile` and `docker-compose.yml` for multi-stage production builds.
