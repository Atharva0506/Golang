# 03_Advanced: The Go Senior Engineer Roadmap

This module contains the final step of the Go mastery journey. Here you will find 15 distinct production-grade concepts that separate Junior Go developers from Senior Go developers.

## Modules

* **`01_context`**: Managing request timeouts and cancellations across API boundaries.
* **`02_pointers`**: Memory efficiency and passing references instead of bulky copies.
* **`03_buffers`**: Pre-allocating slice memory (`make([]int, 0, size)`) and buffered channels to prevent deadlocks.
* **`04_interfaces_advanced`**: Type assertions (`v, ok := i.(string)`), the empty interface (`any`), and Interface Composition.
* **`05_reflection`**: Inspecting struct tags and fields at runtime (how `json` and `gorm` work under the hood).
* **`06_design_patterns`**: The Go-famous "Functional Options" pattern for highly configurable, future-proof struct initialization.
* **`07_concurrency_patterns`**: Fan-in, Fan-out, and fixed-size Worker Pools for massive data processing.
* **`08_middleware`**: The HTTP Decorator pattern (wrapping routes with Auth, Logging, and Panic Recovery).
* **`09_database_sql`**: Connection Pooling, safe error variable reuse (`:=` vs `=`), and Transactions.
* **`10_testing_mocks`**: Dependency Injection and Interface Mocking (faking databases or APIs to save money during tests).
* **`11_benchmarking`**: Scientifically proving memory efficiency using `func BenchmarkX(b *testing.B)` and `-benchmem`.
* **`12_graceful_shutdown`**: Using OS Signals and `server.Shutdown(ctx)` to safely drain traffic before Kubernetes kills a pod.
* **`13_websockets`**: Persistent, bidirectional TCP communication for live-chat and dashboards using `gorilla/websocket`.
* **`14_grpc`**: High-performance, binary-compressed Server-to-Server RPC communication with Protobufs.
* **`15_cgo`**: Calling native C code (like SQLite or FFmpeg) directly from inside Go files.
