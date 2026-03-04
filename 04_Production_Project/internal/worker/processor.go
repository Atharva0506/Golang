package worker

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

// Job defines any background task that can be processed by the worker pool.
type Job interface {
	Name() string
	Execute(ctx context.Context) error
}

// WorkerPool manages a pool of goroutines that process jobs from a buffered channel.
type WorkerPool struct {
	jobs        chan Job
	workerCount int
	wg          sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool with the given number of workers and queue capacity.
func NewWorkerPool(workerCount, queueSize int) *WorkerPool {
	return &WorkerPool{
		jobs:        make(chan Job, queueSize),
		workerCount: workerCount,
	}
}

// Start launches the worker goroutines. Must be called with a cancellable context.
func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(ctx, i)
	}
	slog.Info("worker pool started", "workers", wp.workerCount)
}

// worker processes jobs from the jobs channel until the context is cancelled or the channel is closed.
func (wp *WorkerPool) worker(ctx context.Context, id int) {
	defer wp.wg.Done()
	slog.Info("worker started", "id", id)
	for {
		select {
		case <-ctx.Done():
			slog.Info("worker stopping", "id", id)
			return
		case job, ok := <-wp.jobs:
			if !ok {
				return
			}
			slog.Info("processing job", "worker", id, "job", job.Name())
			start := time.Now()
			if err := job.Execute(ctx); err != nil {
				slog.Error("job failed", "worker", id, "job", job.Name(), "error", err)
			} else {
				slog.Info("job completed", "worker", id, "job", job.Name(), "duration", time.Since(start))
			}
		}
	}
}

// Submit pushes a job onto the queue. If the queue is full, the job is dropped.
func (wp *WorkerPool) Submit(job Job) {
	select {
	case wp.jobs <- job:
		slog.Info("job submitted", "job", job.Name())
	default:
		slog.Warn("job queue full, dropping job", "job", job.Name())
	}
}

// Shutdown closes the job channel and waits for all in-progress jobs to complete.
func (wp *WorkerPool) Shutdown() {
	close(wp.jobs)
	wp.wg.Wait()
	slog.Info("worker pool shut down")
}

// LogJob is a sample job that logs a message. Useful for testing the worker pool.
type LogJob struct {
	Message string
}

// Name returns the job name for logging.
func (j *LogJob) Name() string { return "LogJob" }

// Execute logs the message and simulates work.
func (j *LogJob) Execute(ctx context.Context) error {
	slog.Info("executing log job", "message", j.Message)
	time.Sleep(1 * time.Second)
	return nil
}
