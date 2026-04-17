package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// Always configure a custom http.Client with a timeout.
// The default http.DefaultClient has NO timeout — a slow server will hang your program forever!
var client = &http.Client{
	Timeout: 10 * time.Second,
}

// 1. Simple GET request
func getPost(id int) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)

	resp, err := client.Get(url)
	if err != nil {
		slog.Error("GET failed", "error", err)
		return
	}
	defer resp.Body.Close() // ALWAYS close the body to free the underlying TCP connection!

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("read body failed", "error", err)
		return
	}

	slog.Info("GET response", "status", resp.StatusCode, "body", string(body))
}

// 2. POST request with a JSON body
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func createPost(post Post) {
	data, err := json.Marshal(post)
	if err != nil {
		slog.Error("marshal failed", "error", err)
		return
	}

	resp, err := client.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		slog.Error("POST failed", "error", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	slog.Info("POST response", "status", resp.StatusCode, "body", string(body))
}

// 3. Request with custom headers and context
// Use http.NewRequestWithContext to attach a context for cancellation/timeout control.
func getWithHeaders(ctx context.Context, url, apiKey string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// 4. Retry logic — a real-world pattern for transient failures.
// Retry up to maxAttempts times with exponential backoff.
func getWithRetry(url string, maxAttempts int) ([]byte, error) {
	var lastErr error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, err := client.Do(req)

		if err == nil && resp.StatusCode < 500 {
			defer resp.Body.Close()
			return io.ReadAll(resp.Body)
		}

		if resp != nil {
			resp.Body.Close()
		}

		lastErr = fmt.Errorf("attempt %d failed: %w", attempt, err)
		backoff := time.Duration(attempt*attempt) * 100 * time.Millisecond // exponential backoff
		slog.Warn("retrying request", "attempt", attempt, "backoff", backoff)
		time.Sleep(backoff)
	}

	return nil, fmt.Errorf("all %d attempts failed: %w", maxAttempts, lastErr)
}

func main() {
	fmt.Println("--- 1. Simple GET ---")
	getPost(1)

	fmt.Println("\n--- 2. POST with JSON Body ---")
	createPost(Post{Title: "Go HTTP Client", Body: "Making requests the right way", UserID: 1})

	fmt.Println("\n--- 3. GET with Context Deadline ---")
	// A 3-second deadline is passed down through the context.
	// If the server takes longer, the request is cancelled automatically.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	body, err := getWithHeaders(ctx, "https://jsonplaceholder.typicode.com/posts/2", "my-api-key")
	if err != nil {
		slog.Error("request with headers failed", "error", err)
	} else {
		slog.Info("response with headers", "body", string(body))
	}

	fmt.Println("\n--- 4. Retry Logic ---")
	_, err = getWithRetry("https://jsonplaceholder.typicode.com/posts/3", 3)
	if err != nil {
		slog.Error("retry failed", "error", err)
	} else {
		slog.Info("retry succeeded")
	}
}
