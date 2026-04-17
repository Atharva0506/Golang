# Centralized build and development commands.
# Usage:
#   make run        — run the API server locally
#   make run-worker — run the background worker locally
#   make build      — compile a stripped binary into ./bin/api
#   make test       — run all unit tests with race detector
#   make lint       — run go vet on all packages
#   make check      — run every _Test suite in the repository
#   make docker-up  — start all services with docker-compose
#   make docker-down— stop and remove all containers + volumes
#   make migrate    — apply database migrations
#   make proto      — regenerate gRPC protobuf bindings

BINARY     := bin/api
WORKER_BIN := bin/worker
CMD_API    := ./cmd/api
CMD_WORKER := ./cmd/worker
MIGRATIONS := ./migrations

.PHONY: run run-worker build build-worker test lint check docker-up docker-down migrate proto clean

run:
	go run $(CMD_API)/main.go

run-worker:
	go run $(CMD_WORKER)/main.go

build:
	@mkdir -p bin
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BINARY) $(CMD_API)/main.go

build-worker:
	@mkdir -p bin
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(WORKER_BIN) $(CMD_WORKER)/main.go

test:
	go test -race -count=1 ./...

lint:
	go vet ./...

# Run every _Test suite found in the repo root.
check:
	@echo "==> Running 02_Intermediate_Test"
	cd ../02_Intermediate_Test && go test -v ./...
	@echo "==> Running 03_Advanced_Test"
	cd ../03_Advanced_Test && go test -v ./...
	@echo "==> Running 04_Advanced_Test"
	cd ../04_Advanced_Test && go test -v ./...
	@echo "==> Running 04_Production_Project unit tests"
	go test -race -count=1 ./...

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down -v

migrate:
	go run -tags migrate ./cmd/api/main.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/**/*.proto

clean:
	rm -rf bin/

