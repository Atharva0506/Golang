# Contributing to the Go Learning Path

Thank you for your interest! This document explains how the repository is organised and how to work through the assignments.

---

## Repository Structure

| Directory | Purpose |
|-----------|---------|
| `01_Beginner/` | Annotated examples covering Go fundamentals |
| `02_Intermediate/` | Examples for structuring real applications |
| `03_Advanced/` | Production-grade patterns and tooling |
| `04_Production_Project/` | A complete, runnable Clean Architecture backend skeleton |
| `05_Expert/` | Distributed-systems patterns for senior engineers |
| `02_Intermediate_Test/` | Interactive assignments for Intermediate concepts |
| `03_Advanced_Test/` | Comprehensive integration-style assignment for Advanced concepts |
| `04_Advanced_Test/` | Profiling & tracing assignments for the new Advanced modules |

---

## How to Complete an Assignment

Each `_Test` directory contains two files:

| File | Role |
|------|------|
| `main.go` | Stub functions with `// TODO` comments — **this is where you write your code** |
| `main_test.go` | Test suite — **do not modify this file** |

1. Open `main.go` and implement every `TODO` item.
2. Run the tests to check your work:

```bash
cd 02_Intermediate_Test
go test -v ./...
```

3. When all tests pass, move on to the next section.

---

## Running Benchmarks (04_Advanced_Test)

The profiling assignment includes benchmarks to let you measure your optimisation:

```bash
cd 04_Advanced_Test
go test -bench=. -benchmem ./...
```

You should see `ConcatFast` allocate significantly fewer bytes per operation than `ConcatSlow`.

---

## Viewing the Solutions

All solutions are permanently saved on the `solution` branch:

```bash
# View any solution
git checkout solution

# Return to your own work
git checkout main
```

---

## Running the Production Project Locally

```bash
cd 04_Production_Project

# 1. Copy the example config
cp .env.example .env

# 2. Spin up PostgreSQL and Redis
docker-compose up -d db cache

# 3. Run the API server
make run

# 4. Run the background worker (separate terminal)
make run-worker
```

Visit `http://localhost:8080/health` to confirm the server is up.

To bring up the full observability stack (Prometheus + Grafana):

```bash
docker-compose up -d
```

- **Prometheus**: http://localhost:9091
- **Grafana** (admin/admin): http://localhost:3000

---

## Adding a New Learning Module

1. Create a new directory in the appropriate section, e.g. `02_Intermediate/16_my_topic/`.
2. Add a `main.go` with a `package main` and an annotated `main()` function.
3. If the module needs external dependencies, initialise a `go.mod`:
   ```bash
   cd 02_Intermediate/16_my_topic
   go mod init my_topic
   go get <dependency>
   ```
4. Update the root `README.md` to list the new module.
5. If the module introduces testable concepts, add a task to the relevant `_Test` directory.

---

## Code Style

- Follow standard Go formatting (`gofmt` / `goimports`).
- Comments should explain *why*, not *what*.
- Prefer table-driven tests.
- Run `go vet ./...` before committing.
