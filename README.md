# Go Learning Path

This repository contains my practice code and notes as I journey from a beginner to an advanced Go programmer.

## Structure

The learning path is divided into five main sections:

### 🟢 01 Beginner (Foundations)
* 01_hello_world
* 02_variables_and_types
* 03_user_input
* 04_math_and_crypto
* 05_pointers
* 06_arrays_and_slices
* 07_maps
* 08_structs
* 09_if_else
* 10_switch
* 11_loops
* 12_functions
* 13_methods
* 14_defer
* 15_files

### 📝 01_Beginner_Test
Interactive programming assignments to practice Beginner Go concepts.

### 🟡 02 Intermediate (Structuring Applications)
* 01_interfaces
* 02_error_handling
* 03_custom_errors
* 04_generics
* 05_goroutines
* 06_channels
* 07_waitgroups
* 08_mutex
* 09_json
* 10_modules
* 11_building_api
* 12_regex — Pattern matching with the `regexp` package
* 13_http_client — Outbound HTTP requests, retries, and context-scoped timeouts
* 14_environment_variables — `os.Getenv`, `os.LookupEnv`, and typed config loading
* 15_embedded_files — `//go:embed` for shipping static assets inside your binary

### 📝 02_Intermediate_Test
Interactive programming assignments to practice Intermediate Go concepts (WaitGroups, Mutexes, Interfaces, Regex, HTTP Client, Environment Variables).

### 🔴 03 Advanced (Production Ready)
* 01_context
* 02_pointers
* 03_buffers
* 04_interfaces_advanced
* 05_reflection
* 06_design_patterns
* 07_concurrency_patterns
* 08_middleware
* 09_database_sql
* 10_testing_mocks
* 11_benchmarking
* 12_graceful_shutdown
* 13_websockets
* 14_grpc
* 15_cgo
* 16_profiling — `net/http/pprof` endpoint, CPU + heap profiling with `go tool pprof`
* 17_tracing — Distributed tracing concepts; manual span propagation via `context`
* 18_rate_limiting — Token bucket & fixed-window algorithms from scratch
* 19_event_driven — Generic in-process event bus with channels
* 20_cli_app — Production CLI with `cobra` + `viper`

### 📝 03_Advanced_Test
A massive, production-grade integration test assignment covering Pointers, Interfaces, Reflection, Design Patterns, Worker Pools, Middleware, SQL, and OS Signals.

### 📝 04_Advanced_Test
Profiling & tracing assignments: optimise a slow function using benchmarks, implement a full distributed tracing span system.

---

### 🚀 04 Production Project
A complete, production-ready "Clean Architecture" Golang backend. It includes:
- Decoupled layers (Delivery, Service, Repository)
- Full Postgres repository implementations for Users, Signals, and Notifications
- Docker Compose stack with PostgreSQL, Redis, Prometheus, and Grafana
- JWT authentication, rate limiting, and request-metrics middleware
- WebSocket hub for live signal broadcasting
- Background worker pool for async notification delivery
- gRPC service alongside the REST API
- Graceful shutdown with OS signal handling
- GitHub Actions CI pipeline

---

### ⭐ 05 Expert (Distributed Systems Patterns)
Production patterns used by senior engineers at scale:
* 01_circuit_breaker — Prevent cascade failures from a degraded downstream service
* 02_saga — Orchestrate multi-service distributed transactions with automatic compensation
* 03_cqrs — Separate the write model from optimised read projections
* 04_event_sourcing — Store state as an immutable event log; rebuild via replay

---

### smart-contract
go Lang gin REST apis for english auction smart contract

---

## 💡 How to Check the Solutions!

Throughout this learning path, you will find interactive `.go` assignments inside the `_Test` directories where you must write the implementation to make the `go test -v` script pass.

If you get stuck or want to see the correct Go Senior Engineer implementation for **any** of the tests, the completed code is permanently saved to the `solution` branch!

To view the solutions, open your terminal and run:
```bash
# Switch to the solutions branch
git checkout solution
```

To go back to your own code (or the blank assignments) run:
```bash
# Switch back to the main branch
git checkout main
```

## 🤝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on adding new modules, working through assignments, and running the production project locally.
